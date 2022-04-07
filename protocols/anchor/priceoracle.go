package anchor

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type PriceOracle struct {
	*terra.Contract
}

func NewPriceOracle(querier *terra.Querier) (*PriceOracle, error) {
	contract, err := terra.NewContract(querier, "terra1cgg6yef7qcdm070qftghfulaxmllgmvk77nc7t")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &PriceOracle{
		Contract: contract,
	}, nil
}

func (o *PriceOracle) Prices(ctx context.Context) (map[string]decimal.Decimal, error) {
	var q struct {
		Prices struct {
		} `json:"prices"`
	}

	type response struct {
		Prices []struct {
			Asset           string          `json:"asset"`
			Price           decimal.Decimal `json:"price"`
			LastUpdatedTime int             `json:"last_updated_time"`
		} `json:"prices"`
	}
	var r response
	err := o.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}

	res := make(map[string]decimal.Decimal)
	for _, price := range r.Prices {
		res[price.Asset] = price.Price
	}
	return res, nil
}

func (o *PriceOracle) Price(ctx context.Context, token terra.Token) (decimal.Decimal, error) {
	prices, err := o.Prices(ctx)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "getting prices")
	}

	res := decimal.Zero
	found := false
	for s, d := range prices {
		if s == token.Address().String() {
			found = true
			res = d
			break
		}
	}
	if !found {
		return decimal.Zero, errors.Errorf("price for %s not found in Price Oracle prices", token.Symbol())
	}
	return res, nil
}
