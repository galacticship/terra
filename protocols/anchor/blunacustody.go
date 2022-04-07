package anchor

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type BLUNACustody struct {
	*terra.Contract
}

func NewBLUNACustody(querier *terra.Querier) (*BLUNACustody, error) {
	contract, err := terra.NewContract(querier, "terra1ptjp2vfjrwh0j0faj9r6katm640kgjxnwwq9kn")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &BLUNACustody{
		Contract: contract,
	}, nil
}

func (c *BLUNACustody) NewDepositCollateralMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		DepositCollateral struct {
		} `json:"deposit_collateral"`
	}
	return terra.BLUNA.NewMsgSendExecute(sender, c.Contract, amount, q)
}

func (c *BLUNACustody) NewWithdrawCollateralMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		WithdrawCollateral struct {
			Amount decimal.Decimal
		} `json:"withdraw_collateral"`
	}
	q.WithdrawCollateral.Amount = terra.BLUNA.ValueToTerra(amount)
	return c.NewMsgExecuteContract(sender, q)
}
