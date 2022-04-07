package mars

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Governance struct {
	*terra.Contract
}

func NewGovernance(querier *terra.Querier) (*Governance, error) {
	contract, err := terra.NewContract(querier, "terra1y8wwr5q24msk55x9smwn0ptyt24fxpwm4l7tjl")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &Governance{
		Contract: contract,
	}, nil
}

func (g *Governance) NewStakeMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		Stake struct{} `json:"stake"`
	}
	return terra.MARS.NewMsgSendExecute(sender, g.Contract, amount, q)
}
