package astroport

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
	contract, err := terra.NewContract(querier, "terra1f68wt2ch3cx2g62dxtc8v68mkdh5wchdgdjwz7")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}

	return &Staking{
		Contract: contract,
	}, nil
}

func (s *Staking) NewEnterMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Enter struct{} `json:"enter"`
	}
	return terra.ASTRO.NewMsgSendExecute(sender, s.Contract, amount, q)
}

func (s *Staking) NewLeaveMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Leave struct{} `json:"leave"`
	}
	return terra.XASTRO.NewMsgSendExecute(sender, s.Contract, amount, q)
}
