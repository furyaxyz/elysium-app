package keeper_test

import (
	"bytes"
	"testing"

	testutil "github.com/furyaxyz/elysium-app/test/util"
	"github.com/furyaxyz/elysium-app/x/qgb/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test that valset creation produces the expected normalized power values.
func TestCurrentValsetNormalization(t *testing.T) {
	// Setup the overflow test
	maxPower64 := make([]uint64, 64)             // users with max power (approx 2^63)
	expPower64 := make([]uint64, 64)             // expected scaled powers
	evmAddrs64 := make([]gethcommon.Address, 64) // need 64 eth addresses for this test
	for i := 0; i < 64; i++ {
		maxPower64[i] = uint64(9223372036854775807)
		expPower64[i] = 67108864 // 2^32 split amongst 64 validators
		evmAddrs64[i] = gethcommon.BytesToAddress(bytes.Repeat([]byte{byte(i + 1)}, 20))
	}

	// any lower than this and a validator won't be created
	const minStake = 1000000

	specs := map[string]struct {
		srcPowers []uint64
		expPowers []uint64
	}{
		"one": {
			srcPowers: []uint64{minStake},
			expPowers: []uint64{4294967296},
		},
		"two": {
			srcPowers: []uint64{minStake * 99, minStake * 1},
			expPowers: []uint64{4252017623, 42949672},
		},
		"four equal": {
			srcPowers: []uint64{minStake, minStake, minStake, minStake},
			expPowers: []uint64{1073741824, 1073741824, 1073741824, 1073741824},
		},
		"four equal max power": {
			srcPowers: []uint64{4294967296, 4294967296, 4294967296, 4294967296},
			expPowers: []uint64{1073741824, 1073741824, 1073741824, 1073741824},
		},
		"overflow": {
			srcPowers: maxPower64,
			expPowers: expPower64,
		},
	}
	for msg, spec := range specs {
		spec := spec
		t.Run(msg, func(t *testing.T) {
			input, ctx := testutil.SetupTestChain(t, spec.srcPowers)
			r, err := input.QgbKeeper.GetCurrentValset(ctx)
			require.NoError(t, err)
			rMembers, err := types.BridgeValidators(r.Members).ToInternal()
			require.NoError(t, err)
			assert.Equal(t, spec.expPowers, rMembers.GetPowers())
		})
	}
}

func TestCheckingLatestAttestationNonceInValsets(t *testing.T) {
	input := testutil.CreateTestEnvWithoutAttestationNonceInit(t)
	k := input.QgbKeeper
	// create a validator to have a  realistic scenario
	testutil.CreateValidator(
		t,
		input,
		testutil.AccAddrs[0],
		testutil.AccPubKeys[0],
		0,
		testutil.ValAddrs[0],
		testutil.ConsPubKeys[0],
		testutil.StakingAmount,
		testutil.EVMAddrs[0],
	)
	// Run the staking endblocker to ensure valset is correct in state
	staking.EndBlocker(input.Context, input.StakingKeeper)
	tests := []struct {
		name          string
		requestFunc   func() error
		expectedError error
	}{
		{
			name: "check latest nonce before getting the latest valset",
			requestFunc: func() error {
				_, err := k.GetLatestValset(input.Context)
				return err
			},
			expectedError: types.ErrLatestAttestationNonceStillNotInitialized,
		},
		{
			name: "check latest nonce before getting the current valset",
			requestFunc: func() error {
				_, err := k.GetCurrentValset(input.Context)
				return err
			},
			expectedError: types.ErrLatestAttestationNonceStillNotInitialized,
		},
		{
			name: "check latest nonce before getting last valset before nonce",
			requestFunc: func() error {
				_, err := k.GetLastValsetBeforeNonce(input.Context, 1)
				return err
			},
			expectedError: types.ErrLatestAttestationNonceStillNotInitialized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.requestFunc()
			assert.ErrorIs(t, err, tt.expectedError)
		})
	}
}
