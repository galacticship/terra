package terra

import (
	"fmt"

	"github.com/galacticship/terra/cosmos"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type RoutePair struct {
	Pair          Pair
	IsToken1First bool
}

func NewRoutePair(pair Pair, token1first bool) RoutePair {
	return RoutePair{
		Pair:          pair,
		IsToken1First: token1first,
	}
}

func (p RoutePair) FirstToken() Token {
	if p.IsToken1First {
		return p.Pair.Token1()
	} else {
		return p.Pair.Token2()
	}
}

func (p RoutePair) SecondToken() Token {
	if p.IsToken1First {
		return p.Pair.Token2()
	} else {
		return p.Pair.Token1()
	}
}

type Route []RoutePair

func NewRoute(pairs ...RoutePair) Route {
	return Route(pairs)
}

func (r Route) Last() RoutePair {
	return r[len(r)-1]
}
func (r Route) First() RoutePair {
	return r[0]
}

func (r Route) Contains(pair Pair) bool {
	for _, t := range r {
		if t.Pair.Equals(pair) {
			return true
		}
	}
	return false
}

func (r Route) Copy() Route {
	res := make(Route, len(r))
	copy(res, r)
	return res
}

func (r Route) String() string {
	var res string
	for _, pair := range r {
		if pair.IsToken1First {
			res += fmt.Sprintf("[%s-%s] ", pair.Pair.Token1().Symbol(), pair.Pair.Token2().Symbol())
		} else {
			res += fmt.Sprintf("[%s-%s] ", pair.Pair.Token2().Symbol(), pair.Pair.Token1().Symbol())
		}
	}
	return res
}

func (r Route) OfferToken() Token {
	if r.First().IsToken1First {
		return r.First().Pair.Token1()
	} else {
		return r.First().Pair.Token2()
	}
}

func (r Route) AskToken() Token {
	if r.Last().IsToken1First {
		return r.Last().Pair.Token2()
	} else {
		return r.Last().Pair.Token1()
	}
}

func (r Route) CopyAndAdd(pair RoutePair) Route {
	return append(r.Copy(), pair)
}

func (r Route) SimulateSwap(offerAmount decimal.Decimal, pools map[Pair]PoolInfo) decimal.Decimal {
	current := r.OfferToken().ValueToTerra(offerAmount)
	for _, pair := range r {
		pool := pools[pair.Pair]
		ra, _, _, _ := ComputeConstantProductSwap(pair.FirstToken().ValueToTerra(pool[pair.FirstToken().Id()]), pair.SecondToken().ValueToTerra(pool[pair.SecondToken().Id()]), current, pair.Pair.CommissionRate())
		current = ra
	}
	return r.AskToken().ValueFromTerra(current)
}

func (r Route) GenerateArbitrageMessages(sender cosmos.AccAddress, offerAmount decimal.Decimal, pools map[Pair]PoolInfo) ([]cosmos.Msg, error) {
	current := r.OfferToken().ValueToTerra(offerAmount)
	var msgs []cosmos.Msg
	for _, pair := range r {
		pool := pools[pair.Pair]
		ra, _, _, bp := ComputeConstantProductSwap(pair.FirstToken().ValueToTerra(pool[pair.FirstToken().Id()]), pair.SecondToken().ValueToTerra(pool[pair.SecondToken().Id()]), current, pair.Pair.CommissionRate())
		newmsg, err := pair.Pair.NewSwapMessage(sender, pair.FirstToken(), current, "0.0", bp)
		if err != nil {
			return nil, errors.Wrapf(err, "generating message for pair %s", pair.Pair)
		}
		msgs = append(msgs, newmsg)
		current = ra
	}
	return msgs, nil
}

func (r Route) Pairs() []Pair {
	var res []Pair
	for _, pair := range r {
		contains := false
		for _, re := range res {
			if re.Equals(pair.Pair) {
				contains = true
				break
			}
		}
		if contains {
			continue
		}
		res = append(res, pair.Pair)
	}
	return res
}
