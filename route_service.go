package terra

type RouteService interface {
	RegisterPairs(pairs ...Pair)
	FindRoutes(offer Token, ask Token, maxDepth int) []Route
	FindArbitrages(token Token, maxDepth int) []Route
	GetAllArbitrages(maxDepth int) []Route
}

type routeService struct {
	pairs []Pair
}

func NewRouteService() RouteService {
	return &routeService{}
}

func (s *routeService) RegisterPairs(pairs ...Pair) {
	s.pairs = append(s.pairs, pairs...)
}

func (s *routeService) walkRoute(route Route, ask Token, depth, maxdepth int) []Route {
	var res []Route
	depth++
	if depth > maxdepth {
		return res
	}
	if route.AskToken().Equals(ask) {
		res = append(res, route)
		return res
	}
	for _, pair := range s.pairs {
		if route.Contains(pair) {
			continue
		}
		var newroute Route
		if route.Last().SecondToken().Equals(pair.Token1()) {
			newroute = route.CopyAndAdd(NewRoutePair(pair, true))
		}
		if route.Last().SecondToken().Equals(pair.Token2()) {
			newroute = route.CopyAndAdd(NewRoutePair(pair, false))
		}
		if newroute == nil {
			continue
		}
		res = append(res, s.walkRoute(newroute, ask, depth, maxdepth)...)
	}
	return res
}

func (s *routeService) FindRoutes(offer Token, ask Token, maxDepth int) []Route {
	var res []Route
	for _, pair := range s.pairs {
		if pair.Token1().Equals(offer) {
			res = append(res, s.walkRoute(NewRoute(NewRoutePair(pair, true)), ask, 0, maxDepth)...)
		}
		if pair.Token2().Equals(offer) {
			res = append(res, s.walkRoute(NewRoute(NewRoutePair(pair, false)), ask, 0, maxDepth)...)
		}
	}
	return res
}

func (s *routeService) FindArbitrages(token Token, maxDepth int) []Route {
	return s.FindRoutes(token, token, maxDepth)
}

func (s *routeService) GetAllArbitrages(maxDepth int) []Route {
	var res []Route
	for _, pair := range s.pairs {
		res = append(res, s.walkRoute(NewRoute(NewRoutePair(pair, true)), pair.Token1(), 0, maxDepth)...)
		res = append(res, s.walkRoute(NewRoute(NewRoutePair(pair, false)), pair.Token2(), 0, maxDepth)...)
	}
	return res
}
