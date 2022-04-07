package prism

import (
	"context"
	"time"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Amps struct {
	*terra.Contract
}

func NewAmps(querier *terra.Querier) (*Amps, error) {
	contract, err := terra.NewContract(querier, "terra1pa4amk66q8punljptzmmftf6ylq3ezyzx6kl9m")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &Amps{
		Contract: contract,
	}, nil
}

type BoostInfo struct {
	AmountBonded          decimal.Decimal
	Boost                 decimal.Decimal
	LastUpdatedTime       time.Time
	BoostAccrualStartTime time.Time
}

func (a *Amps) GetBoost(ctx context.Context, user cosmos.AccAddress) (BoostInfo, error) {
	type query struct {
		GetBoost struct {
			User string `json:"user"`
		} `json:"get_boost"`
	}
	type response struct {
		AmtBonded             decimal.Decimal `json:"amt_bonded"`
		TotalBoost            decimal.Decimal `json:"total_boost"`
		LastUpdated           int64           `json:"last_updated"`
		BoostAccrualStartTime int64           `json:"boost_accrual_start_time"`
	}
	var q query
	q.GetBoost.User = user.String()
	var r response
	err := a.QueryStore(ctx, q, &r)
	if err != nil {
		return BoostInfo{}, errors.Wrap(err, "querying contract store")
	}
	return BoostInfo{
		AmountBonded:          terra.XPRISM.ValueFromTerra(r.AmtBonded),
		Boost:                 r.TotalBoost.Shift(-6),
		LastUpdatedTime:       time.Unix(r.LastUpdated, 0),
		BoostAccrualStartTime: time.Unix(r.BoostAccrualStartTime, 0),
	}, nil
}

func (a *Amps) NewBondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		Bond struct {
		} `json:"bond"`
	}
	var q query
	return terra.XPRISM.NewMsgSendExecute(sender, a.Contract, amount, q)
}

func (a *Amps) NewUnbondMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		Unbond struct {
			Amount decimal.Decimal `json:"amount"`
		} `json:"unbond"`
	}
	var q query
	q.Unbond.Amount = terra.XPRISM.ValueToTerra(amount)
	return a.NewMsgExecuteContract(sender, q)
}
