package mars

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/galacticship/terra/protocols/astroport"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Bootstrap struct {
	*terra.Contract

	marsustpair terra.Pair
}

func (b *Bootstrap) RewardPair() terra.Pair {
	return b.marsustpair
}

func NewBootstrap(ctx context.Context, querier *terra.Querier) (*Bootstrap, error) {
	contract, err := terra.NewContract(querier, "terra1hgyamk2kcy3stqx82wrnsklw9aq7rask5dxfds")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	b := &Bootstrap{
		Contract: contract,
	}
	cfg, err := b.Config(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting config")
	}
	p, err := astroport.NewXykPair(querier, cfg.MarsUstPoolAddress, terra.MARS, terra.UST, terra.ASTRO_MARSUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "getting pool pair object")
	}
	b.marsustpair = p
	return b, nil
}

type BootstrapConfig struct {
	MarsUstPoolAddress string
}

func (b *Bootstrap) Config(ctx context.Context) (BootstrapConfig, error) {
	var q struct {
		Config struct {
		} `json:"config"`
	}
	type response struct {
		Owner                   string `json:"owner"`
		MarsTokenAddress        string `json:"mars_token_address"`
		AstroTokenAddress       string `json:"astro_token_address"`
		AirdropContractAddress  string `json:"airdrop_contract_address"`
		LockdropContractAddress string `json:"lockdrop_contract_address"`
		AstroportLpPool         string `json:"astroport_lp_pool"`
		LpTokenAddress          string `json:"lp_token_address"`
		MarsLpStakingContract   string `json:"mars_lp_staking_contract"`
		GeneratorContract       string `json:"generator_contract"`
		MarsRewards             string `json:"mars_rewards"`
		MarsVestingDuration     int    `json:"mars_vesting_duration"`
		LpTokensVestingDuration int    `json:"lp_tokens_vesting_duration"`
		InitTimestamp           int    `json:"init_timestamp"`
		MarsDepositWindow       int    `json:"mars_deposit_window"`
		UstDepositWindow        int    `json:"ust_deposit_window"`
		WithdrawalWindow        int    `json:"withdrawal_window"`
	}

	var r response
	err := b.QueryStore(ctx, q, &r)
	if err != nil {
		return BootstrapConfig{}, errors.Wrap(err, "querying contract store")
	}
	return BootstrapConfig{
		MarsUstPoolAddress: r.AstroportLpPool,
	}, nil
}

type BootstrapUserInfo struct {
	WithdrawableLpShares decimal.Decimal `json:"withdrawable_lp_shares"`
}

func (b *Bootstrap) UserInfo(ctx context.Context, address cosmos.AccAddress) (BootstrapUserInfo, error) {
	type query struct {
		UserInfo struct {
			Address string `json:"address"`
		} `json:"user_info"`
	}
	var q query
	q.UserInfo.Address = address.String()
	type response struct {
		WithdrawableLpShares decimal.Decimal `json:"withdrawable_lp_shares"`
	}

	var r response
	err := b.QueryStore(ctx, q, &r)
	if err != nil {
		return BootstrapUserInfo{}, errors.Wrap(err, "querying contract store")
	}
	return BootstrapUserInfo{
		WithdrawableLpShares: b.marsustpair.LpToken().ValueFromTerra(r.WithdrawableLpShares),
	}, nil
}

func (b *Bootstrap) NewClaimRewardsMessage(sender cosmos.AccAddress, withDrawUnlockedShares bool) (cosmos.Msg, error) {
	type query struct {
		ClaimRewards struct {
			WithdrawUnlockedShares bool `json:"withdraw_unlocked_shares"`
		} `json:"claim_rewards"`
	}
	var q query
	q.ClaimRewards.WithdrawUnlockedShares = withDrawUnlockedShares
	return b.NewMsgExecuteContract(sender, q)
}
