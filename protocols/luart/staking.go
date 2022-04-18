package luart

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Staking struct {
	*terra.Contract
}

func NewStaking(querier *terra.Querier) (*Staking, error) {
	c, err := terra.NewContract(querier, "terra1dlcwvsy6t7skge7s2dtdvr75lakltwr3xk9j2d")
	if err != nil {
		return nil, errors.Wrap(err, "init base contract")
	}
	return &Staking{
		c,
	}, nil
}

func (s *Staking) NewBondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Bond struct{} `json:"bond"`
	}
	return terra.LUART.NewMsgSendExecute(sender, s.Contract, amount, q)
}

func (s *Staking) NewSubmitToUnbondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		SubmitToUnbond struct {
			Amount decimal.Decimal `json:"amount"`
		} `json:"submit_to_unbond"`
	}
	q.SubmitToUnbond.Amount = terra.LUART.ValueToTerra(amount)
	return s.NewMsgExecuteContract(sender, q)
}
