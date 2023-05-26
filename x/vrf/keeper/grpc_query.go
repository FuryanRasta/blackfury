package keeper

import (
	"github.com/blackfurystation/blackfury/v3/x/vrf/types"
)

var _ types.QueryServer = Keeper{}
