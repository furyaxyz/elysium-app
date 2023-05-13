package proof

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/furyaxyz/elysium-app/pkg/appconsts"
	"github.com/furyaxyz/elysium-app/pkg/da"
	appns "github.com/furyaxyz/elysium-app/pkg/namespace"
	"github.com/furyaxyz/elysium-app/pkg/shares"
	"github.com/furyaxyz/elysium-app/pkg/square"
	"github.com/furyaxyz/elysium-app/pkg/wrapper"
	"github.com/tendermint/tendermint/crypto/merkle"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
)

// NewTxInclusionProof returns a new share inclusion proof for the given
// transaction index.
func NewTxInclusionProof(txs [][]byte, txIndex uint64) (types.ShareProof, error) {
	if txIndex >= uint64(len(txs)) {
		return types.ShareProof{}, fmt.Errorf("txIndex %d out of bounds", txIndex)
	}

	builder, err := square.NewBuilder(appconsts.DefaultMaxSquareSize, txs...)
	if err != nil {
		return types.ShareProof{}, err
	}

	dataSquare, err := builder.Export()
	if err != nil {
		return types.ShareProof{}, err
	}

	shareRange, err := builder.FindTxShareRange(int(txIndex))
	if err != nil {
		return types.ShareProof{}, err
	}

	namespace := getTxNamespace(txs[txIndex])
	return NewShareInclusionProof(dataSquare, dataSquare.Size(), namespace, uint64(shareRange.Start), uint64(shareRange.End))
}

func getTxNamespace(tx []byte) (ns appns.Namespace) {
	_, isBlobTx := types.UnmarshalBlobTx(tx)
	if isBlobTx {
		return appns.PayForBlobNamespace
	}
	return appns.TxNamespace
}

// NewShareInclusionProof returns an NMT inclusion proof for a set of shares to the data root.
// Expects the share range to be pre-validated.
// Note: only supports inclusion proofs for shares belonging to the same namespace.
func NewShareInclusionProof(
	allRawShares []shares.Share,
	squareSize uint64,
	namespace appns.Namespace,
	startShare uint64,
	endShare uint64,
) (types.ShareProof, error) {
	startRow := startShare / squareSize
	endRow := endShare / squareSize
	startLeaf := startShare % squareSize
	endLeaf := endShare % squareSize

	eds, err := da.ExtendShares(shares.ToBytes(allRawShares))
	if err != nil {
		return types.ShareProof{}, err
	}

	edsRowRoots := eds.RowRoots()

	// create the binary merkle inclusion proof for all the square rows to the data root
	_, allProofs := merkle.ProofsFromByteSlices(append(edsRowRoots, eds.ColRoots()...))
	rowProofs := make([]*merkle.Proof, endRow-startRow+1)
	rowRoots := make([]tmbytes.HexBytes, endRow-startRow+1)
	for i := startRow; i <= endRow; i++ {
		rowProofs[i-startRow] = allProofs[i]
		rowRoots[i-startRow] = edsRowRoots[i]
	}

	// get the extended rows containing the shares.
	rows := make([][]shares.Share, endRow-startRow+1)
	for i := startRow; i <= endRow; i++ {
		shares, err := shares.FromBytes(eds.Row(uint(i)))
		if err != nil {
			return types.ShareProof{}, err
		}
		rows[i-startRow] = shares
	}

	var shareProofs []*tmproto.NMTProof //nolint:prealloc
	var rawShares [][]byte
	for i, row := range rows {
		// create an nmt to generate a proof.
		// we have to re-create the tree as the eds one is not accessible.
		tree := wrapper.NewErasuredNamespacedMerkleTree(squareSize, uint(i))
		for _, share := range row {
			tree.Push(
				share.ToBytes(),
			)
		}

		// make sure that the generated root is the same as the eds row root.
		if !bytes.Equal(rowRoots[i].Bytes(), tree.Root()) {
			return types.ShareProof{}, errors.New("eds row root is different than tree root")
		}

		startLeafPos := startLeaf
		endLeafPos := endLeaf

		// if this is not the first row, then start with the first leaf
		if i > 0 {
			startLeafPos = 0
		}
		// if this is not the last row, then select for the rest of the row
		if i != (len(rows) - 1) {
			endLeafPos = squareSize - 1
		}

		rawShares = append(rawShares, shares.ToBytes(row[startLeafPos:endLeafPos+1])...)
		proof, err := tree.ProveRange(int(startLeafPos), int(endLeafPos+1))
		if err != nil {
			return types.ShareProof{}, err
		}

		shareProofs = append(shareProofs, &tmproto.NMTProof{
			Start:    int32(proof.Start()),
			End:      int32(proof.End()),
			Nodes:    proof.Nodes(),
			LeafHash: proof.LeafHash(),
		})
	}

	return types.ShareProof{
		RowProof: types.RowProof{
			RowRoots: rowRoots,
			Proofs:   rowProofs,
			StartRow: uint32(startRow),
			EndRow:   uint32(endRow),
		},
		Data:             rawShares,
		ShareProofs:      shareProofs,
		NamespaceID:      namespace.ID,
		NamespaceVersion: uint32(namespace.Version),
	}, nil
}
