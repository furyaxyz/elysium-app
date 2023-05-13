package keeper_test

import (
	"testing"

	testkeeper "github.com/furyaxyz/elysium-app/test/util/keeper"
	"github.com/furyaxyz/elysium-app/x/blob/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlobKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.GasPerBlobByte, k.GasPerBlobByte(ctx))
}
