package astroport

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type BootstrapAuction struct {
	*terra.Contract
	astroustpair terra.Pair
}

func (b *BootstrapAuction) RewardPair() terra.Pair {
	return b.astroustpair
}

func NewBootstrapAuction(ctx context.Context, querier *terra.Querier) (*BootstrapAuction, error) {
	contract, err := terra.NewContract(querier, "terra1tvld5k6pus2yh7pcu7xuwyjedn7mjxfkkkjjap")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	b := &BootstrapAuction{
		Contract: contract,
	}
	cfg, err := b.Config(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting config")
	}
	p, err := NewXykPair(querier, cfg.AstroUstPoolAddress, terra.ASTRO, terra.UST, terra.ASTRO_ASTROUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "getting pool pair object")
	}
	b.astroustpair = p
	return b, nil
}

type BootstrapAuctionUserInfo struct {
	WithdrawableLpShares decimal.Decimal `json:"withdrawable_lp_shares"`
}

func (b *BootstrapAuction) UserInfo(ctx context.Context, address cosmos.AccAddress) (BootstrapAuctionUserInfo, error) {
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
		return BootstrapAuctionUserInfo{}, errors.Wrap(err, "querying contract store")
	}
	return BootstrapAuctionUserInfo{
		WithdrawableLpShares: b.astroustpair.LpToken().ValueFromTerra(r.WithdrawableLpShares),
	}, nil
}

type BootstrapAuctionConfig struct {
	AstroUstPoolAddress string
}

func (b *BootstrapAuction) Config(ctx context.Context) (BootstrapAuctionConfig, error) {
	var q struct {
		Config struct {
		} `json:"config"`
	}
	type response struct {
		Owner                   string `json:"owner"`
		AstroTokenAddress       string `json:"astro_token_address"`
		AirdropContractAddress  string `json:"airdrop_contract_address"`
		LockdropContractAddress string `json:"lockdrop_contract_address"`
		PoolInfo                struct {
			AstroUstPoolAddress    string `json:"astro_ust_pool_address"`
			AstroUstLpTokenAddress string `json:"astro_ust_lp_token_address"`
		} `json:"pool_info"`
		GeneratorContract       string `json:"generator_contract"`
		AstroIncentiveAmount    string `json:"astro_incentive_amount"`
		LpTokensVestingDuration int    `json:"lp_tokens_vesting_duration"`
		InitTimestamp           int    `json:"init_timestamp"`
		DepositWindow           int    `json:"deposit_window"`
		WithdrawalWindow        int    `json:"withdrawal_window"`
	}

	var r response
	err := b.QueryStore(ctx, q, &r)
	if err != nil {
		return BootstrapAuctionConfig{}, errors.Wrap(err, "querying contract store")
	}
	return BootstrapAuctionConfig{
		AstroUstPoolAddress: r.PoolInfo.AstroUstPoolAddress,
	}, nil
}

func (b *BootstrapAuction) NewClaimAllRewardsMessage(ctx context.Context, sender cosmos.AccAddress) (cosmos.Msg, decimal.Decimal, error) {
	r, err := b.UserInfo(ctx, sender)
	if err != nil {
		return nil, decimal.Zero, errors.Wrap(err, "getting user info")
	}
	msg, err := b.NewClaimRewardsMessage(sender, r.WithdrawableLpShares)
	if err != nil {
		return nil, decimal.Zero, errors.Wrap(err, "creating message")
	}
	return msg, r.WithdrawableLpShares, nil
}

func (b *BootstrapAuction) NewClaimRewardsMessage(sender cosmos.AccAddress, withdrawLpShares decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		ClaimRewards struct {
			WithdrawLpShares decimal.Decimal `json:"withdraw_lp_shares"`
		} `json:"claim_rewards"`
	}
	q.ClaimRewards.WithdrawLpShares = b.astroustpair.LpToken().ValueToTerra(withdrawLpShares)
	return b.NewMsgExecuteContract(sender, q)
}
