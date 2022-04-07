package anchor

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Overseer struct {
	*terra.Contract
}

func NewOverseer(querier *terra.Querier) (*Overseer, error) {
	contract, err := terra.NewContract(querier, "terra1tmnqgvg567ypvsvk6rwsga3srp7e3lg6u0elp8")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &Overseer{
		Contract: contract,
	}, nil
}

func (o *Overseer) BorrowLimit(ctx context.Context, borrower cosmos.AccAddress) (decimal.Decimal, error) {
	type query struct {
		BorrowLimit struct {
			Borrower string `json:"borrower"`
		} `json:"borrow_limit"`
	}
	var q query
	q.BorrowLimit.Borrower = borrower.String()

	type response struct {
		Borrower    string          `json:"borrower"`
		BorrowLimit decimal.Decimal `json:"borrow_limit"`
	}
	var r response
	err := o.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying store")
	}
	return terra.UST.ValueFromTerra(r.BorrowLimit), nil
}

func (o *Overseer) Collaterals(ctx context.Context, borrower cosmos.AccAddress) (map[string]decimal.Decimal, error) {
	type query struct {
		Collaterals struct {
			Borrower string `json:"borrower"`
		} `json:"collaterals"`
	}
	var q query
	q.Collaterals.Borrower = borrower.String()
	type response struct {
		Borrower    string      `json:"borrower"`
		Collaterals [][2]string `json:"collaterals"`
	}
	var r response
	err := o.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, errors.Wrap(err, "querying store")
	}
	res := make(map[string]decimal.Decimal)
	for _, collateral := range r.Collaterals {
		token, err := terra.Cw20TokenFromAddress(ctx, o.Contract.Querier(), collateral[0])
		if err != nil {
			continue
		}
		value, err := decimal.NewFromString(collateral[1])
		if err != nil {
			continue
		}
		res[collateral[0]] = token.ValueFromTerra(value)
	}
	return res, nil
}

func (o *Overseer) Collateral(ctx context.Context, borrower cosmos.AccAddress, token terra.Token) (decimal.Decimal, error) {
	collaterals, err := o.Collaterals(ctx, borrower)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "getting collaterals")
	}
	res := decimal.Zero
	found := false
	for s, d := range collaterals {
		if s == token.Address().String() {
			found = true
			res = d
			break
		}
	}
	if !found {
		return decimal.Zero, errors.Errorf("No collateral found for %s", token.Symbol())
	}
	return res, nil
}

func (o *Overseer) NewUnlockCollateralMessage(sender cosmos.AccAddress, token terra.Cw20Token, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		UnlockCollateral struct {
			Collaterals [][2]string `json:"collaterals"`
		} `json:"unlock_collateral"`
	}
	var q query
	q.UnlockCollateral.Collaterals = append(q.UnlockCollateral.Collaterals, [2]string{
		token.Address().String(),
		token.ValueToTerra(amount).String(),
	})
	return o.NewMsgExecuteContract(sender, q)
}

func (o *Overseer) NewLockCollateralMessage(sender cosmos.AccAddress, token terra.Cw20Token, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		LockCollateral struct {
			Collaterals [][2]string `json:"collaterals"`
		} `json:"lock_collateral"`
	}
	var q query
	q.LockCollateral.Collaterals = append(q.LockCollateral.Collaterals, [2]string{
		token.Address().String(),
		token.ValueToTerra(amount).String(),
	})
	return o.NewMsgExecuteContract(sender, q)
}
