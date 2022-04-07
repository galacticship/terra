package cosmos

import sdktypes "github.com/cosmos/cosmos-sdk/types"

type (
	AccAddress  = sdktypes.AccAddress
	ValAddress  = sdktypes.ValAddress
	ConsAddress = sdktypes.ConsAddress
)

var (
	AccAddressFromBech32  = sdktypes.AccAddressFromBech32
	AccAddressFromHex     = sdktypes.AccAddressFromHex
	ValAddressFromBech32  = sdktypes.ValAddressFromBech32
	ValAddressFromHex     = sdktypes.ValAddressFromHex
	ConsAddressFromBech32 = sdktypes.ConsAddressFromBech32
	ConsAddressFromHex    = sdktypes.ConsAddressFromHex
)
