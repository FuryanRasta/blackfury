package app

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	v2 "github.com/furyanrasta/blackfury/v3/app/upgrades/v2"
)

// BeginBlockForks executes any necessary fork logic based upon the current block height.
func BeginBlockForks(ctx sdk.Context, app *Blackfury) {
	switch ctx.BlockHeight() {
	case v2.UpgradeHeight:
		// NOTE: only run for mainnet
		if !strings.HasPrefix(ctx.ChainID(), "blackfury_3001-") {
			return
		}

		upgradePlan := upgradetypes.Plan{
			Name:   v2.UpgradeName,
			Info:   v2.UpgradeInfo,
			Height: v2.UpgradeHeight,
		}
		err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan)
		if err != nil {
			panic(err)
		}
	default:
		// do nothing
		return
	}
}
