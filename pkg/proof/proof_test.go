package proof_test

import (
	"bytes"
	"testing"

	"github.com/furyaxyz/elysium-app/app"
	"github.com/furyaxyz/elysium-app/app/encoding"
	"github.com/furyaxyz/elysium-app/test/util/blobfactory"
	"github.com/furyaxyz/elysium-app/test/util/testfactory"

	"github.com/furyaxyz/elysium-app/pkg/da"
	"github.com/furyaxyz/elysium-app/pkg/proof"
	"github.com/furyaxyz/elysium-app/pkg/square"

	"github.com/furyaxyz/elysium-app/pkg/appconsts"
	appns "github.com/furyaxyz/elysium-app/pkg/namespace"
	"github.com/furyaxyz/elysium-app/pkg/shares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTxInclusionProof(t *testing.T) {
	blockTxs := testfactory.GenerateRandomTxs(50, 500).ToSliceOfBytes()
	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	blockTxs = append(blockTxs, blobfactory.RandBlobTxs(encCfg.TxConfig.TxEncoder(), 50, 500).ToSliceOfBytes()...)
	require.Len(t, blockTxs, 100)

	type test struct {
		name      string
		txs       [][]byte
		txIndex   uint64
		expectErr bool
	}
	tests := []test{
		{
			name:      "empty txs returns error",
			txs:       nil,
			txIndex:   0,
			expectErr: true,
		},
		{
			name:      "txIndex 0 of block data",
			txs:       blockTxs,
			txIndex:   0,
			expectErr: false,
		},
		{
			name:      "last regular transaction of block data",
			txs:       blockTxs,
			txIndex:   49,
			expectErr: false,
		},
		{
			name:      "first blobTx of block data",
			txs:       blockTxs,
			txIndex:   50,
			expectErr: false,
		},
		{
			name:      "last blobTx of block data",
			txs:       blockTxs,
			txIndex:   99,
			expectErr: false,
		},
		{
			name:      "txIndex 100 of block data returns error because only 100 txs",
			txs:       blockTxs,
			txIndex:   100,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proof, err := proof.NewTxInclusionProof(
				tt.txs,
				tt.txIndex,
			)
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.True(t, proof.VerifyProof())
		})
	}
}

func TestNewShareInclusionProof(t *testing.T) {
	ns1 := appns.MustNewV0(bytes.Repeat([]byte{1}, appns.NamespaceVersionZeroIDSize))
	ns2 := appns.MustNewV0(bytes.Repeat([]byte{2}, appns.NamespaceVersionZeroIDSize))
	ns3 := appns.MustNewV0(bytes.Repeat([]byte{3}, appns.NamespaceVersionZeroIDSize))

	encCfg := encoding.MakeConfig(app.ModuleEncodingRegisters...)
	blobTxs := blobfactory.RandBlobTxsWithNamespaces(encCfg.TxConfig.TxEncoder(), []appns.Namespace{ns1, ns2, ns3}, []int{500, 500, 500})
	txs := testfactory.GenerateRandomTxs(50, 500)
	txs = append(txs, blobTxs...)

	dataSquare, err := square.Construct(txs.ToSliceOfBytes(), appconsts.DefaultMaxSquareSize)
	if err != nil {
		panic(err)
	}

	// erasure the data square which we use to create the data root.
	eds, err := da.ExtendShares(shares.ToBytes(dataSquare))
	require.NoError(t, err)

	// create the new data root by creating the data availability header (merkle
	// roots of each row and col of the erasure data).
	dah := da.NewDataAvailabilityHeader(eds)
	dataRoot := dah.Hash()

	type test struct {
		name          string
		startingShare int64
		endingShare   int64
		namespaceID   appns.Namespace
		expectErr     bool
	}
	tests := []test{
		{
			name:          "negative starting share",
			startingShare: -1,
			endingShare:   99,
			namespaceID:   appns.TxNamespace,
			expectErr:     true,
		},
		{
			name:          "negative ending share",
			startingShare: 0,
			endingShare:   -99,
			namespaceID:   appns.TxNamespace,
			expectErr:     true,
		},
		{
			name:          "ending share lower than starting share",
			startingShare: 1,
			endingShare:   0,
			namespaceID:   appns.TxNamespace,
			expectErr:     true,
		},
		{
			name:          "ending share higher than number of shares available in square size of 32",
			startingShare: 0,
			endingShare:   4097,
			namespaceID:   appns.TxNamespace,
			expectErr:     true,
		},
		{
			name:          "1 transaction share",
			startingShare: 0,
			endingShare:   0,
			namespaceID:   appns.TxNamespace,
			expectErr:     false,
		},
		{
			name:          "10 transaction shares",
			startingShare: 0,
			endingShare:   9,
			namespaceID:   appns.TxNamespace,
			expectErr:     false,
		},
		{
			name:          "53 transaction shares",
			startingShare: 0,
			endingShare:   52,
			namespaceID:   appns.TxNamespace,
			expectErr:     false,
		},
		{
			name:          "shares from different namespaces",
			startingShare: 48,
			endingShare:   54,
			namespaceID:   appns.TxNamespace,
			expectErr:     true,
		},
		{
			name:          "shares from PFB namespace",
			startingShare: 53,
			endingShare:   55,
			namespaceID:   appns.PayForBlobNamespace,
			expectErr:     false,
		},
		{
			name:          "blob shares for first namespace",
			startingShare: 56,
			endingShare:   57,
			namespaceID:   ns1,
			expectErr:     false,
		},
		{
			name:          "blob shares for third namespace",
			startingShare: 60,
			endingShare:   61,
			namespaceID:   ns3,
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualNID, err := proof.ParseNamespace(dataSquare, tt.startingShare, tt.endingShare)
			if tt.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.namespaceID, actualNID)
			proof, err := proof.NewShareInclusionProof(
				dataSquare,
				dataSquare.Size(),
				tt.namespaceID,
				uint64(tt.startingShare),
				uint64(tt.endingShare),
			)
			require.NoError(t, err)
			assert.NoError(t, proof.Validate(dataRoot))
		})
	}
}
