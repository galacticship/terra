package prism

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Governance struct {
	*terra.Contract
}

func NewGovernance(querier *terra.Querier) (*Governance, error) {
	c, err := terra.NewContract(querier, "terra1h4al753uvwmhxwhn2dlvm9gfk0jkf52xqasmq2")
	if err != nil {
		return nil, errors.Wrap(err, "init base contract")
	}
	return &Governance{
		c,
	}, nil
}

type XPrismState struct {
	ExchangeRate               decimal.Decimal `json:"exchange_rate"`
	EffectiveXprismSupply      decimal.Decimal `json:"effective_xprism_supply"`
	EffectiveUnderlyingPrism   decimal.Decimal `json:"effective_underlying_prism"`
	TotalPendingWithdrawXprism decimal.Decimal `json:"total_pending_withdraw_xprism"`
	TotalPendingWithdrawPrism  decimal.Decimal `json:"total_pending_withdraw_prism"`
}

func (g *Governance) XPrismState(ctx context.Context) (XPrismState, error) {
	var q struct {
		XprismState struct{} `json:"xprism_state"`
	}
	type response struct {
		ExchangeRate               decimal.Decimal `json:"exchange_rate"`
		EffectiveXprismSupply      decimal.Decimal `json:"effective_xprism_supply"`
		EffectiveUnderlyingPrism   decimal.Decimal `json:"effective_underlying_prism"`
		TotalPendingWithdrawXprism decimal.Decimal `json:"total_pending_withdraw_xprism"`
		TotalPendingWithdrawPrism  decimal.Decimal `json:"total_pending_withdraw_prism"`
	}
	var r response
	err := g.QueryStore(ctx, q, &r)
	if err != nil {
		return XPrismState{}, errors.Wrap(err, "querying contract store")
	}
	return XPrismState{
		ExchangeRate:               r.ExchangeRate,
		EffectiveXprismSupply:      terra.XPRISM.ValueFromTerra(r.EffectiveXprismSupply),
		EffectiveUnderlyingPrism:   terra.PRISM.ValueFromTerra(r.EffectiveUnderlyingPrism),
		TotalPendingWithdrawXprism: terra.XPRISM.ValueFromTerra(r.TotalPendingWithdrawXprism),
		TotalPendingWithdrawPrism:  terra.PRISM.ValueFromTerra(r.TotalPendingWithdrawPrism),
	}, nil

}
