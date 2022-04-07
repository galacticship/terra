package prism

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Farm struct {
	*terra.Contract
}

func NewFarm(querier *terra.Querier) (*Farm, error) {
	c, err := terra.NewContract(querier, "terra1ns5nsvtdxu53dwdthy3yxs6x3w2hf3fclhzllc")
	if err != nil {
		return nil, errors.Wrap(err, "init base contract")
	}
	return &Farm{
		c,
	}, nil
}

func (f *Farm) NewBondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Bond struct{} `json:"bond"`
	}
	return terra.YLUNA.NewMsgSendExecute(sender, f.Contract, amount, q)
}

func (f *Farm) NewUnBondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Unbond struct {
			Amount decimal.Decimal
		} `json:"unbond"`
	}
	q.Unbond.Amount = terra.YLUNA.ValueToTerra(amount)
	return f.NewMsgExecuteContract(sender, q)
}

func (f *Farm) NewActivateBoostMessage(sender cosmos.AccAddress) (cosmos.Msg, error) {
	var q struct {
		ActivateBoost struct{} `json:"activate_boost"`
	}
	return f.NewMsgExecuteContract(sender, q)
}

func (f *Farm) newClaimWithdrawnRewardsMessage(sender cosmos.AccAddress, claimType string) (cosmos.Msg, error) {
	var q struct {
		ClaimWithdrawnRewards struct {
			ClaimType string `json:"claim_type"`
		} `json:"claim_withdrawn_rewards"`
	}
	q.ClaimWithdrawnRewards.ClaimType = claimType
	return f.NewMsgExecuteContract(sender, q)
}

func (f *Farm) NewClaimWithdrawnRewardsMessage(sender cosmos.AccAddress) (cosmos.Msg, error) {
	return f.newClaimWithdrawnRewardsMessage(sender, "Prism")
}
func (f *Farm) NewClaimAndPledgeWithdrawnRewardsMessage(sender cosmos.AccAddress) (cosmos.Msg, error) {
	return f.newClaimWithdrawnRewardsMessage(sender, "Amps")
}
