package terra

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Pair interface {
	CommissionRate() decimal.Decimal
	LpToken() Cw20Token
	Token1() Token
	Token2() Token
	ContractAddressString() string
	ContractAddress() cosmos.AccAddress

	PoolInfo(ctx context.Context) (PoolInfo, error)
	NewWithdrawLiquidityMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error)
	Share(ctx context.Context, lpAmount decimal.Decimal) (token1Amount decimal.Decimal, token2Amount decimal.Decimal, err error)
	NewSimpleSwapMessage(sender cosmos.AccAddress, offerToken Token, amount decimal.Decimal) (cosmos.Msg, error)
	NewSwapMessage(sender cosmos.AccAddress, offerToken Token, amount decimal.Decimal, spread string, beliefPrice decimal.Decimal) (cosmos.Msg, error)
	SimulateSwap(ctx context.Context, offer Token, amount decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal, error)
	Equals(pair Pair) bool
	String() string
}

type BasePair struct {
	*Contract

	token1         Token
	token2         Token
	lpToken        Cw20Token
	commissionRate decimal.Decimal

	aiFactory AssetInfoFactory
}

func NewBasePair(querier *Querier, contractAddress string, token1 Token, token2 Token, lpToken Cw20Token, commissionRate decimal.Decimal, aiFactory AssetInfoFactory) (*BasePair, error) {
	contract, err := NewContract(querier, contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "init contract object")
	}
	p := &BasePair{
		Contract:       contract,
		commissionRate: commissionRate,
		aiFactory:      aiFactory,
		token1:         token1,
		token2:         token2,
		lpToken:        lpToken,
	}
	return p, nil
}

func (p BasePair) CommissionRate() decimal.Decimal {
	return p.commissionRate
}

func (p BasePair) LpToken() Cw20Token {
	return p.lpToken
}

func (p BasePair) Token1() Token {
	return p.token1
}

func (p BasePair) Token2() Token {
	return p.token2
}

func (p BasePair) ContractAddressString() string {
	return p.contractAddress.String()
}
func (p BasePair) ContractAddress() cosmos.AccAddress {
	return p.contractAddress
}

func (p *BasePair) SetToken1(token Token) {
	p.token1 = token
}
func (p *BasePair) SetToken2(token Token) {
	p.token2 = token
}
func (p *BasePair) SetLpToken(token Cw20Token) {
	p.lpToken = token
}

func (p *BasePair) Config(ctx context.Context) (tokens []Token, lpToken Token, err error) {
	var q struct {
		Pair struct {
		} `json:"pair"`
	}
	type response struct {
		AssetInfos     []json.RawMessage `json:"asset_infos"`
		ContractAddr   string            `json:"contract_addr"`
		LiquidityToken string            `json:"liquidity_token"`
	}
	var r response
	err = p.QueryStore(ctx, q, &r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "querying contract store")
	}

	var assets []AssetInfo
	for _, info := range r.AssetInfos {
		a, err := p.aiFactory.DecodeFromJson(info)
		if err != nil {
			return nil, nil, errors.Wrap(err, "decoding asset json")
		}
		assets = append(assets, a)
	}

	if len(assets) < 2 {
		return nil, nil, errors.Errorf("not enough token in pair config: %d", len(assets))
	}
	t1, err := GetTokenFromAssetInfo(ctx, p.Contract.Querier(), assets[0])
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting token1 from asset info")
	}

	t2, err := GetTokenFromAssetInfo(ctx, p.Contract.Querier(), assets[1])
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting token2 from asset info")
	}

	lptoken, err := Cw20TokenFromAddress(ctx, p.Contract.Querier(), r.LiquidityToken)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "getting lp token: %s", r.LiquidityToken)
	}

	return []Token{t1, t2}, lptoken, nil
}

type PoolInfo map[string]decimal.Decimal

func (p BasePair) PoolInfo(ctx context.Context) (PoolInfo, error) {
	type query struct {
		Pool struct {
		} `json:"pool"`
	}
	type response struct {
		Assets []struct {
			Info   json.RawMessage `json:"info"`
			Amount decimal.Decimal `json:"amount"`
		} `json:"assets"`
		TotalShare decimal.Decimal `json:"total_share"`
	}
	var r response
	err := p.QueryStore(ctx, query{}, &r)
	if err != nil {
		return PoolInfo{}, errors.Wrap(err, "querying contract store")
	}
	res := make(map[string]decimal.Decimal)
	for _, a := range r.Assets {
		ai, err := p.aiFactory.DecodeFromJson(a.Info)
		if err != nil {
			return PoolInfo{}, errors.Wrap(err, "decoding asset info")
		}
		if p.token1.Id() == ai.Id() {
			res[ai.Id()] = p.token1.ValueFromTerra(a.Amount)
		} else if p.token2.Id() == ai.Id() {
			res[ai.Id()] = p.token2.ValueFromTerra(a.Amount)
		} else {
			return PoolInfo{}, errors.New("asset unknown")
		}
	}
	return res, nil
}

func (p BasePair) NewWithdrawLiquidityMessage(sender cosmos.AccAddress, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		WithdrawLiquidity struct {
		} `json:"withdraw_liquidity"`
	}
	var q query
	return p.lpToken.NewMsgSendExecute(sender, p.Contract, amount, q)
}

func (p BasePair) Share(ctx context.Context, lpAmount decimal.Decimal) (token1Amount decimal.Decimal, token2Amount decimal.Decimal, err error) {
	type query struct {
		Share struct {
			Amount decimal.Decimal `json:"amount"`
		} `json:"share"`
	}
	type response []struct {
		Info   json.RawMessage `json:"info"`
		Amount decimal.Decimal `json:"amount"`
	}
	var q query
	q.Share.Amount = p.lpToken.ValueToTerra(lpAmount)
	var r response
	err = p.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, decimal.Zero, errors.Wrap(err, "querying contract store")
	}
	if len(r) < 2 {
		return decimal.Zero, decimal.Zero, errors.Errorf("not enough token in share response: %d", len(r))
	}

	for _, a := range r {
		ai, err := p.aiFactory.DecodeFromJson(a.Info)
		if err != nil {
			return decimal.Zero, decimal.Zero, errors.Wrap(err, "decoding asset info")
		}
		if p.token1.Id() == ai.Id() {
			token1Amount = p.token1.ValueFromTerra(a.Amount)
		} else if p.token2.Id() == ai.Id() {
			token2Amount = p.token2.ValueFromTerra(a.Amount)
		} else {
			return decimal.Zero, decimal.Zero, errors.New("asset unknown")
		}
	}
	return token1Amount, token2Amount, nil

}

func (p BasePair) NewSimpleSwapMessage(sender cosmos.AccAddress, offerToken Token, amount decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		Swap struct {
			OfferAsset struct {
				Info   AssetInfo       `json:"info"`
				Amount decimal.Decimal `json:"amount"`
			} `json:"offer_asset"`
		} `json:"swap"`
	}

	var q query
	q.Swap.OfferAsset.Amount = offerToken.ValueToTerra(amount)
	q.Swap.OfferAsset.Info = p.aiFactory.NewFromToken(offerToken)
	return offerToken.NewMsgSendExecute(sender, p.Contract, amount, q)
}

func (p BasePair) NewSwapMessage(sender cosmos.AccAddress, offerToken Token, amount decimal.Decimal, spread string, beliefPrice decimal.Decimal) (cosmos.Msg, error) {
	type query struct {
		Swap struct {
			MaxSpread  string `json:"max_spread"`
			OfferAsset struct {
				Info   AssetInfo       `json:"info"`
				Amount decimal.Decimal `json:"amount"`
			} `json:"offer_asset"`
			BeliefPrice decimal.Decimal `json:"belief_price"`
		} `json:"swap"`
	}

	var q query
	q.Swap.OfferAsset.Info = p.aiFactory.NewFromToken(offerToken)
	q.Swap.OfferAsset.Amount = offerToken.ValueToTerra(amount)
	q.Swap.MaxSpread = spread
	q.Swap.BeliefPrice = beliefPrice
	return offerToken.NewMsgSendExecute(sender, p.Contract, amount, q)
}

func ComputeConstantProductSwap(offerPool decimal.Decimal, askPool decimal.Decimal, offerAmount decimal.Decimal, commissionRate decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	if offerPool.Equals(decimal.Zero) || askPool.Equals(decimal.Zero) || offerAmount.Equals(decimal.Zero) {
		return decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero
	}
	cp := offerPool.Mul(askPool)
	returnAmount := (askPool.Sub(cp.Div(offerPool.Add(offerAmount)))).Truncate(0)
	spread := offerAmount.Mul(askPool).Div(offerPool).Sub(returnAmount).Truncate(0)
	commissionAmount := returnAmount.Mul(commissionRate).Truncate(0)
	beliefPrice := offerAmount.Div(returnAmount)
	returnAmount = returnAmount.Sub(commissionAmount)
	return returnAmount, spread, commissionAmount, beliefPrice
}

func (p BasePair) SimulateSwap(ctx context.Context, offer Token, amount decimal.Decimal) (decimal.Decimal, decimal.Decimal, decimal.Decimal, error) {
	type query struct {
		Simulation struct {
			OfferAsset struct {
				Info   AssetInfo       `json:"info"`
				Amount decimal.Decimal `json:"amount"`
			} `json:"offer_asset"`
		} `json:"simulation"`
	}
	var q query
	q.Simulation.OfferAsset.Info = p.aiFactory.NewFromToken(offer)
	q.Simulation.OfferAsset.Amount = offer.ValueToTerra(amount)
	type response struct {
		ReturnAmount     decimal.Decimal `json:"return_amount"`
		SpreadAmount     decimal.Decimal `json:"spread_amount"`
		CommissionAmount decimal.Decimal `json:"commission_amount"`
	}
	var r response
	err := p.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, decimal.Zero, decimal.Zero, errors.Wrap(err, "querying contract store")
	}

	return r.ReturnAmount, r.SpreadAmount, r.CommissionAmount, nil

}

func (p *BasePair) Equals(pair Pair) bool {
	return p.contractAddress.String() == pair.ContractAddressString()
}

func (p *BasePair) String() string {
	t1symbol := ""
	if p.token1 != nil {
		t1symbol = p.token1.Symbol()
	}
	t2symbol := ""
	if p.token2 != nil {
		t2symbol = p.token2.Symbol()
	}
	return fmt.Sprintf("%s / %s", t1symbol, t2symbol)
}
