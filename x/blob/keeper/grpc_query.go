package keeper

import (
	"github.com/elysiumorg/elysium-app/x/blob/types"
)

var _ types.QueryServer = Keeper{}
