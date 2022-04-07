package mars

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Field struct {
	*terra.Contract
}

func NewANCUSTField(querier *terra.Querier) (*Field, error) {
	return NewField(querier, "terra1vapq79y9cqghqny7zt72g4qukndz282uvqwtz6")
}

func NewField(querier *terra.Querier, address string) (*Field, error) {
	contract, err := terra.NewContract(querier, address)
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &Field{
		contract,
	}, nil
}

type FieldSnapshot struct {
	PositionLpValue decimal.Decimal
	DebtValue       decimal.Decimal
	Ltv             decimal.Decimal
}

func (f *Field) Snapshot(ctx context.Context, address string) (FieldSnapshot, error) {
	type query struct {
		Snapshot struct {
			User string `json:"user"`
		} `json:"snapshot"`
	}
	type response struct {
		Time     int `json:"time"`
		Height   int `json:"height"`
		Position struct {
			BondUnits decimal.Decimal `json:"bond_units"`
			DebtUnits decimal.Decimal `json:"debt_units"`
		} `json:"position"`
		Health struct {
			BondAmount decimal.Decimal `json:"bond_amount"`
			BondValue  decimal.Decimal `json:"bond_value"`
			DebtAmount decimal.Decimal `json:"debt_amount"`
			DebtValue  decimal.Decimal `json:"debt_value"`
			Ltv        decimal.Decimal `json:"ltv"`
		} `json:"health"`
	}
	var q query
	q.Snapshot.User = address
	var r response
	err := f.QueryStore(ctx, q, &r)
	if err != nil {
		return FieldSnapshot{}, errors.Wrap(err, "querying contract store")
	}
	res := FieldSnapshot{}
	return res, nil
}
