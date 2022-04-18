package astroport

import (
	"context"

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

func (s *Staking) TotalShares(ctx context.Context) (decimal.Decimal, error) {
	var q struct {
		TotalShares struct{} `json:"total_shares"`
	}
	var r decimal.Decimal
	err := s.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract")
	}
	return terra.XASTRO.ValueFromTerra(r), nil
}

func (s *Staking) TotalDeposit(ctx context.Context) (decimal.Decimal, error) {
	var q struct {
		TotalDeposit struct{} `json:"total_deposit"`
	}
	var r decimal.Decimal
	err := s.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract")
	}
	return terra.ASTRO.ValueFromTerra(r), nil
}

func (s *Staking) XASTROPerASTRO(ctx context.Context) (decimal.Decimal, error) {
	totalShares, err := s.TotalShares(ctx)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "getting total shares")
	}
	totalDeposit, err := s.TotalDeposit(ctx)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "getting total deposit")
	}
	if totalDeposit.Equals(decimal.Zero) {
		return decimal.NewFromInt(1), nil
	}
	return totalShares.Div(totalDeposit), nil
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
