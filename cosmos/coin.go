package cosmos

import sdktypes "github.com/cosmos/cosmos-sdk/types"

type (
	Coin     = sdktypes.Coin
	Coins    = sdktypes.Coins
	DecCoin  = sdktypes.DecCoin
	DecCoins = sdktypes.DecCoins
)

var (
	NewCoin         = sdktypes.NewCoin
	NewInt64Coin    = sdktypes.NewInt64Coin
	NewCoins        = sdktypes.NewCoins
	NewDecCoin      = sdktypes.NewDecCoin
	NewInt64DecCoin = sdktypes.NewInt64DecCoin
	NewDecCoins     = sdktypes.NewDecCoins
)
