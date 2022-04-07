package anchor

import (
	"context"

	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Market struct {
	*terra.Contract
}

func NewMarket(querier *terra.Querier) (*Market, error) {
	contract, err := terra.NewContract(querier, "terra1sepfj7s0aeg5967uxnfk4thzlerrsktkpelm5s")
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	return &Market{
		contract,
	}, nil
}

type BorrowerInfo struct {
	InterestIndex  decimal.Decimal
	RewardIndex    decimal.Decimal
	LoanAmount     decimal.Decimal
	PendingRewards decimal.Decimal
}

func (m Market) BorrowerInfo(ctx context.Context, borrower cosmos.AccAddress) (BorrowerInfo, error) {
	type query struct {
		BorrowerInfo struct {
			Borrower string `json:"borrower"`
		} `json:"borrower_info"`
	}
	var q query
	q.BorrowerInfo.Borrower = borrower.String()

	type response struct {
		Borrower       string          `json:"borrower"`
		InterestIndex  decimal.Decimal `json:"interest_index"`
		RewardIndex    decimal.Decimal `json:"reward_index"`
		LoanAmount     decimal.Decimal `json:"loan_amount"`
		PendingRewards decimal.Decimal `json:"pending_rewards"`
	}
	var r response
	err := m.QueryStore(ctx, q, &r)
	if err != nil {
		return BorrowerInfo{}, errors.Wrap(err, "querying store")
	}
	return BorrowerInfo{
		InterestIndex:  r.InterestIndex,
		RewardIndex:    r.RewardIndex,
		LoanAmount:     terra.UST.ValueFromTerra(r.LoanAmount),
		PendingRewards: terra.ANC.ValueFromTerra(r.PendingRewards),
	}, nil
}

func (m Market) NewDepositUSTMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		DepositStable struct{} `json:"deposit_stable"`
	}
	return terra.UST.NewMsgSendExecute(sender, m.Contract, amount, q)
}

func (m Market) NewRedeemAUSTMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		RedeemStable struct{} `json:"redeem_stable"`
	}
	return terra.AUST.NewMsgSendExecute(sender, m.Contract, amount, q)
}

func (m Market) NewClaimRewardsMessage(sender cosmos.AccAddress) (cosmos.Msg, error) {
	var q struct {
		ClaimRewards struct{} `json:"claim_rewards"`
	}
	return m.NewMsgExecuteContract(sender, q)
}

func (m Market) NewBorrowStableMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		BorrowStable struct {
			BorrowAmount decimal.Decimal `json:"borrow_amount"`
		} `json:"borrow_stable"`
	}
	q.BorrowStable.BorrowAmount = terra.UST.ValueToTerra(amount)
	return m.NewMsgExecuteContract(sender, q)
}

func (m Market) NewRepayStableMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	var q struct {
		RepayStable struct{} `json:"repay_stable"`
	}
	return terra.UST.NewMsgSendExecute(sender, m.Contract, amount, q)
}
