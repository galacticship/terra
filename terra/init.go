package terra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	core "github.com/terra-money/core/types"
)

func init() {
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetCoinType(core.CoinType)
	sdkConfig.SetPurpose(44)
	sdkConfig.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(core.Bech32PrefixConsAddr, core.Bech32PrefixConsPub)
	sdkConfig.SetAddressVerifier(core.AddressVerifier)
	sdkConfig.Seal()
}
