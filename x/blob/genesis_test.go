package blob_test

import (
	"testing"

	keepertest "github.com/furyaxyz/elysium-app/test/util/keeper"
	"github.com/furyaxyz/elysium-app/x/blob"
	"github.com/furyaxyz/elysium-app/x/blob/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.BlobKeeper(t)
	blob.InitGenesis(ctx, *k, genesisState)
	got := blob.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultParams(), got.Params)
}
