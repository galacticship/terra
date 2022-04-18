package mars

import (
	"context"

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

func (g *Governance) XMARSperMARS(ctx context.Context) (decimal.Decimal, error) {
	var q struct {
		XMarsPerMars struct{} `json:"x_mars_per_mars"`
	}
	var r decimal.Decimal
	err := g.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract store")
	}
	return r, nil
}

func (g *Governance) MARSperXMARS(ctx context.Context) (decimal.Decimal, error) {
	var q struct {
		MarsPerXMars struct{} `json:"mars_per_x_mars"`
	}
	var r decimal.Decimal
	err := g.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract store")
	}
	return r, nil
}
