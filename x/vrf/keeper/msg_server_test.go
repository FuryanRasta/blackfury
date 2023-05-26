package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/blackfurystation/blackfury/v3/testutil/keeper"
	"github.com/blackfurystation/blackfury/v3/x/vrf/keeper"
	"github.com/blackfurystation/blackfury/v3/x/vrf/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VRFKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
