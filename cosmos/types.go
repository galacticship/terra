package cosmos

import sdktypes "github.com/cosmos/cosmos-sdk/types"

type (
	Int = sdktypes.Int
	Dec = sdktypes.Dec
)

var (
	NewInt                   = sdktypes.NewInt
	NewIntFromBigInt         = sdktypes.NewIntFromBigInt
	NewIntFromString         = sdktypes.NewIntFromString
	NewIntFromUint64         = sdktypes.NewIntFromUint64
	NewIntWithDecimal        = sdktypes.NewIntWithDecimal
	NewDec                   = sdktypes.NewDec
	NewDecCoinFromCoin       = sdktypes.NewDecCoinFromCoin
	NewDecCoinFromDec        = sdktypes.NewDecCoinFromDec
	NewDecFromBigInt         = sdktypes.NewDecFromBigInt
	NewDecFromBigIntWithPrec = sdktypes.NewDecFromBigIntWithPrec
	NewDecFromInt            = sdktypes.NewDecFromInt
	NewDecFromIntWithPrec    = sdktypes.NewDecFromIntWithPrec
	NewDecFromStr            = sdktypes.NewDecFromStr
	NewDecWithPrec           = sdktypes.NewDecWithPrec
)
