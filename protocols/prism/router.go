package prism

import (
	"github.com/galacticship/terra"
	"github.com/pkg/errors"
)

type Router struct {
	*terra.BaseRouter
}

func NewRouter(querier *terra.Querier) (terra.Router, error) {
	r, err := terra.NewBaseRouter(querier, "terra1yrc0zpwhuqezfnhdgvvh7vs5svqtgyl7pu3n6c", NewAssetInfoFactory(), newOperation)
	if err != nil {
		return nil, errors.Wrap(err, "creating base router")
	}

	PRISMUST, err := NewPair(querier, "terra19d2alknajcngdezrdhq40h6362k92kz23sz62u", terra.PRISM, terra.UST, terra.PRISM_PRISMUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMUST pair")
	}
	PRISMLUNA, err := NewPair(querier, "terra1r38qlqt69lez4nja5h56qwf4drzjpnu8gz04jd", terra.PRISM, terra.LUNA, terra.PRISM_PRISMLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMLUNA pair")
	}
	PRISMPLUNA, err := NewPair(querier, "terra1persuahr6f8fm6nyup0xjc7aveaur89nwgs5vs", terra.PRISM, terra.PLUNA, terra.PRISM_PRISMPLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMPLUNA pair")
	}
	PRISMXPRISM, err := NewPair(querier, "terra1czynvm64nslq2xxavzyrrhau09smvana003nrf", terra.PRISM, terra.XPRISM, terra.PRISM_PRISMXPRISMLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMXPRISM pair")
	}
	PRISMCLUNA, err := NewPair(querier, "terra1yxgq5y6mw30xy9mmvz9mllneddy9jaxndrphvk", terra.PRISM, terra.CLUNA, terra.PRISM_PRISMCLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMCLUNA pair")
	}
	PRISMYLUNA, err := NewPair(querier, "terra1kqc65n5060rtvcgcktsxycdt2a4r67q2zlvhce", terra.PRISM, terra.YLUNA, terra.PRISM_PRISMYLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMYLUNA pair")
	}

	r.SetPairs(
		PRISMUST,
		PRISMLUNA,
		PRISMPLUNA,
		PRISMXPRISM,
		PRISMCLUNA,
		PRISMYLUNA,
	)

	return &Router{
		BaseRouter: r,
	}, nil
}

type operation struct {
	Swap struct {
		OfferAssetInfo terra.AssetInfo `json:"offer_asset_info"`
		AskAssetInfo   terra.AssetInfo `json:"ask_asset_info"`
	} `json:"prism_swap"`
}

func newOperation(aiFactory terra.AssetInfoFactory, offer terra.Token, ask terra.Token) interface{} {
	var res operation
	res.Swap.OfferAssetInfo = aiFactory.NewFromToken(offer)
	res.Swap.AskAssetInfo = aiFactory.NewFromToken(ask)
	return res
}

func (r Router) String() string {
	return "prism"
}
