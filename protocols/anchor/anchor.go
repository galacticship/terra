package anchor

import (
	"github.com/galacticship/terra"
	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Anchor struct {
	Market       *Market
	Overseer     *Overseer
	PriceOracle  *PriceOracle
	BLUNACustody *BLUNACustody
}

func NewAnchor(querier *terra.Querier) (*Anchor, error) {
	market, err := NewMarket(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating market")
	}
	overseer, err := NewOverseer(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating overseer")
	}
	priceOracle, err := NewPriceOracle(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating price oracle")
	}
	blunaCustody, err := NewBLUNACustody(querier)
	if err != nil {
		return nil, errors.Wrap(err, "creating bluna custody")
	}

	return &Anchor{
		Market:       market,
		Overseer:     overseer,
		PriceOracle:  priceOracle,
		BLUNACustody: blunaCustody,
	}, nil
}

func (a *Anchor) NewProvideBLUNAMessages(sender cosmos.AccAddress, amount decimal.Decimal) ([]cosmos.Msg, error) {
	m1, err := a.BLUNACustody.NewDepositCollateralMessage(sender, amount)
	if err != nil {
		return nil, errors.Wrap(err, "creating deposit collateral to BLUNA custody message")
	}
	m2, err := a.Overseer.NewLockCollateralMessage(sender, terra.BLUNA, amount)
	if err != nil {
		return nil, errors.Wrap(err, "creating lock collateral from overseer message")
	}
	return []cosmos.Msg{m1, m2}, nil
}

func (a *Anchor) NewWithdrawBLUNAMessages(sender cosmos.AccAddress, amount decimal.Decimal) ([]cosmos.Msg, error) {
	m1, err := a.Overseer.NewUnlockCollateralMessage(sender, terra.BLUNA, amount)
	if err != nil {
		return nil, errors.Wrap(err, "creating unlock collateral from overseer message")
	}
	m2, err := a.BLUNACustody.NewWithdrawCollateralMessage(sender, amount)
	if err != nil {
		return nil, errors.Wrap(err, "creating withdraw collateral from BLUNA custody message")
	}
	return []cosmos.Msg{m1, m2}, nil
}
