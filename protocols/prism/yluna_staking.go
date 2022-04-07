package prism

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type YLUNAStaking struct {
	*terra.Contract
}

func NewYLUNAStaking(querier *terra.Querier) (*YLUNAStaking, error) {
	contract, err := terra.NewContract(querier, "terra1p7jp8vlt57cf8qwazjg58qngwvarmszsamzaru")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &YLUNAStaking{
		Contract: contract,
	}, nil
}

type RewardInfo struct {
	StakedAmount decimal.Decimal
	Rewards      map[string]decimal.Decimal
}

func (y *YLUNAStaking) GetRewardInfo(ctx context.Context, stakerAddress cosmos.AccAddress) (RewardInfo, error) {
	type query struct {
		RewardInfo struct {
			StakerAddr string `json:"staker_addr"`
		} `json:"reward_info"`
	}
	type response struct {
		StakerAddr   string          `json:"staker_addr"`
		StakedAmount decimal.Decimal `json:"staked_amount"`
		Rewards      []struct {
			Info struct {
				Cw20 string `json:"cw20"`
			} `json:"info"`
			Amount decimal.Decimal `json:"amount"`
		} `json:"rewards"`
	}
	var q query
	q.RewardInfo.StakerAddr = stakerAddress.String()
	var r response
	err := y.QueryStore(ctx, q, &r)
	if err != nil {
		return RewardInfo{}, errors.Wrap(err, "querying contract store")
	}
	res := RewardInfo{
		StakedAmount: terra.YLUNA.ValueFromTerra(r.StakedAmount),
		Rewards:      make(map[string]decimal.Decimal),
	}
	for _, reward := range r.Rewards {
		token, err := terra.Cw20TokenFromAddress(ctx, y.Contract.Querier(), reward.Info.Cw20)
		if err != nil {
			return RewardInfo{}, errors.Wrapf(err, "getting token %s", reward.Info.Cw20)
		}
		res.Rewards[reward.Info.Cw20] = token.ValueFromTerra(reward.Amount)
	}
	return res, nil
}

func (y *YLUNAStaking) NewClaimAndConvertRewardMessage(sender cosmos.AccAddress, claimToken terra.Cw20Token) (cosmos.Msg, error) {
	var q struct {
		ConvertAndClaimRewards struct {
			ClaimAsset struct {
				Cw20 string `json:"cw20"`
			} `json:"claim_asset"`
		} `json:"convert_and_claim_rewards"`
	}
	q.ConvertAndClaimRewards.ClaimAsset.Cw20 = claimToken.Address().String()
	return y.NewMsgExecuteContract(sender, q)
}
