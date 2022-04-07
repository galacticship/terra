package terra

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type PriceService struct {
	routers        []Router
	cache          map[string]cachedPrice
	cacheMutex     *sync.Mutex
	cacheTimeout   time.Duration
	maxRouteLength int
}

type cachedPrice struct {
	price       decimal.Decimal
	updatedTime time.Time
}

type PriceServiceOption func(s *PriceService) *PriceService

func WithCacheTimeout(timeout time.Duration) PriceServiceOption {
	return func(s *PriceService) *PriceService {
		s.cacheTimeout = timeout
		return s
	}
}

func WithMaxRouteLength(maxRouteLenght int) PriceServiceOption {
	return func(s *PriceService) *PriceService {
		s.maxRouteLength = maxRouteLenght
		return s
	}
}

func NewPriceService(options ...PriceServiceOption) *PriceService {
	p := &PriceService{
		routers:        nil,
		cache:          make(map[string]cachedPrice),
		cacheMutex:     &sync.Mutex{},
		cacheTimeout:   30 * time.Second,
		maxRouteLength: 2,
	}
	for _, option := range options {
		p = option(p)
	}
	return p
}

func (s *PriceService) AddRouter(router Router) {
	s.routers = append(s.routers, router)
}

func (s *PriceService) GetPriceCached(ctx context.Context, token Token) (decimal.Decimal, error) {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()
	if p, ok := s.cache[token.Id()]; ok && p.updatedTime.Add(s.cacheTimeout).After(time.Now()) {
		return p.price, nil
	}
	p, err := s.getCurrentPrice(ctx, token)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "getting current price")
	}
	s.cache[token.Id()] = cachedPrice{
		price:       p,
		updatedTime: time.Now(),
	}
	return p, nil
}

func (s *PriceService) getCurrentPrice(ctx context.Context, token Token) (decimal.Decimal, error) {
	var bestprice decimal.Decimal
	var bestroute Route
	for _, router := range s.routers {
		newprice, newroute, err := router.SimulateSwap(ctx, token, UST, decimal.NewFromInt(1), s.maxRouteLength)
		if err != nil && err != ErrNoRouteFund {
			return decimal.Zero, errors.Wrapf(err, "simulating swap with router %s", router)
		}
		if newprice.GreaterThan(bestprice) || (newprice.Equals(bestprice) && len(newroute) < len(bestroute)) {
			bestprice = newprice
			bestroute = newroute
		}
	}
	return bestprice, nil
}
