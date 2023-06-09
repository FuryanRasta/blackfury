package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "cosmossdk.io/x/auth/types"
	"cosmossdk.io/x/auth/vesting/exported"
	sdkvesting "cosmossdk.io/x/auth/vesting/types"
)

// ModuleCdc references the global erc20 module codec. Note, the codec should
// ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to modules/erc20 and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"cosmos.vesting.v1beta1.VestingAccount",
		(*exported.VestingAccount)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgClawback{},
		&MsgCreateClawbackVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
