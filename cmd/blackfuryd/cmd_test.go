package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"cosmossdk.io/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/furyanrasta/blackfury/v3/app"
	blackfuryd "github.com/furyanrasta/blackfury/v3/cmd/blackfuryd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := blackfuryd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"blackfury-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "highbury_710-1"),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := blackfuryd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"mykey",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.Error(t, err)
}
