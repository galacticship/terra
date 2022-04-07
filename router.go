package terra

import (
	"context"

	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Router interface {
	String() string
	SimulateSwapWithRoute(ctx context.Context, amount decimal.Decimal, route Route) (decimal.Decimal, error)
	SimulateSwap(ctx context.Context, offer Token, ask Token, amount decimal.Decimal, maxRouteLength int) (decimal.Decimal, Route, error)
	FindAllRoutes(offer Token, ask Token, maxLength int) []Route

	NewSwapMessageWithRoute(sender cosmos.AccAddress, route Route, offerAmount decimal.Decimal, askExpectedAmount decimal.Decimal, maxSpread float64) (cosmos.Msg, error)
	NewSwapMessageWithBestRoute(ctx context.Context, sender cosmos.AccAddress, offer Token, ask Token, offerAmount decimal.Decimal, maxRouteLength int, maxSpread float64) (cosmos.Msg, error)
}

type BaseRouter struct {
	*Contract
	routeService RouteService

	operationFactory func(aiFactory AssetInfoFactory, offer Token, ask Token) interface{}
	aiFactory        AssetInfoFactory
}

func NewBaseRouter(querier *Querier, contractAddress string, aiFactory AssetInfoFactory, operationFactory func(aiFactory AssetInfoFactory, offer Token, ask Token) interface{}) (*BaseRouter, error) {
	contract, err := NewContract(querier, contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "creating base contract")
	}
	return &BaseRouter{
		Contract:         contract,
		routeService:     NewRouteService(),
		operationFactory: operationFactory,
		aiFactory:        aiFactory,
	}, nil
}

func (r BaseRouter) ContractAddress() cosmos.AccAddress {
	return r.contractAddress
}

func (r *BaseRouter) SetPairs(pairs ...Pair) {
	r.routeService.RegisterPairs(pairs...)
}

func (r BaseRouter) FindAllRoutes(offer Token, ask Token, maxLength int) []Route {
	return r.routeService.FindRoutes(offer, ask, maxLength)
}

var ErrNoRouteFund = errors.New("no route found")

func (r BaseRouter) SimulateSwap(ctx context.Context, offer Token, ask Token, amount decimal.Decimal, maxRouteLength int) (decimal.Decimal, Route, error) {
	var resValue decimal.Decimal
	var resRoute Route
	routes := r.FindAllRoutes(offer, ask, maxRouteLength)
	for _, route := range routes {

		tmpValue, err := r.SimulateSwapWithRoute(ctx, amount, route)
		if err != nil {
			continue
		}
		if resValue.LessThan(tmpValue) || (resValue.Equals(tmpValue) && len(resRoute) > len(route)) {
			resValue = tmpValue
			resRoute = route
		}
	}
	if resRoute == nil {
		return decimal.Zero, nil, ErrNoRouteFund
	}
	return resValue, resRoute, nil
}

func (r BaseRouter) SimulateSwapWithRoute(ctx context.Context, amount decimal.Decimal, route Route) (decimal.Decimal, error) {

	if len(route) < 1 {
		return decimal.Zero, errors.Errorf("route length must be greater than 1")
	}
	type query struct {
		SimulateSwapOperations struct {
			OfferAmount decimal.Decimal `json:"offer_amount"`
			Operations  []interface{}   `json:"operations"`
		} `json:"simulate_swap_operations"`
	}
	var q query
	q.SimulateSwapOperations.OfferAmount = route[0].FirstToken().ValueToTerra(amount)
	for _, pair := range route {
		q.SimulateSwapOperations.Operations = append(q.SimulateSwapOperations.Operations,
			r.operationFactory(r.aiFactory, pair.FirstToken(), pair.SecondToken()))
	}

	type response struct {
		Amount decimal.Decimal `json:"amount"`
	}
	var resp response
	err := r.QueryStore(ctx, q, &resp)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract store")
	}

	return route[len(route)-1].SecondToken().ValueFromTerra(resp.Amount), nil
}

func (r BaseRouter) NewSwapMessageWithRoute(sender cosmos.AccAddress, route Route, offerAmount decimal.Decimal, askExpectedAmount decimal.Decimal, maxSpread float64) (cosmos.Msg, error) {
	askExpectedAmount = route.AskToken().ValueToTerra(askExpectedAmount)
	minimumReceived := askExpectedAmount.Sub(askExpectedAmount.Mul(decimal.NewFromFloat(maxSpread).Div(decimal.NewFromInt(100)))).Truncate(0)

	type query struct {
		ExecuteSwapOperations struct {
			OfferAmount    decimal.Decimal `json:"offer_amount"`
			MinimumReceive decimal.Decimal `json:"minimum_receive"`
			Operations     []interface{}   `json:"operations"`
		} `json:"execute_swap_operations"`
	}
	var q query
	q.ExecuteSwapOperations.OfferAmount = route.OfferToken().ValueToTerra(offerAmount)
	q.ExecuteSwapOperations.MinimumReceive = minimumReceived
	for _, pair := range route {
		q.ExecuteSwapOperations.Operations = append(q.ExecuteSwapOperations.Operations,
			r.operationFactory(r.aiFactory, pair.FirstToken(), pair.SecondToken()))
	}
	res, err := route.OfferToken().NewMsgSendExecute(sender, r.Contract, offerAmount, q)
	if err != nil {
		return nil, errors.Wrap(err, "generating MsgSendExecute")
	}
	return res, nil
}

func (r BaseRouter) NewSwapMessageWithBestRoute(ctx context.Context, sender cosmos.AccAddress, offer Token, ask Token, offerAmount decimal.Decimal, maxRouteLength int, maxSpread float64) (cosmos.Msg, error) {
	askExpected, bestRoute, err := r.SimulateSwap(ctx, offer, ask, offerAmount, maxRouteLength)
	if err != nil {
		return nil, errors.Wrap(err, "simulating swaps to find best route")
	}
	return r.NewSwapMessageWithRoute(sender, bestRoute, offerAmount, askExpected, maxSpread)
}
