package astroport

import (
	"github.com/galacticship/terra"
	"github.com/pkg/errors"
)

type router struct {
	*terra.BaseRouter
}

func NewRouter(querier *terra.Querier) (terra.Router, error) {
	r, err := terra.NewBaseRouter(querier, "terra16t7dpwwgx9n3lq6l6te3753lsjqwhxwpday9zx", terra.NewAssetInfoFactory(), newOperation)
	if err != nil {
		return nil, errors.Wrap(err, "creating base router")
	}

	LUNAUST, err := NewXykPair(querier, "terra1m6ywlgn6wrjuagcmmezzz2a029gtldhey5k552", terra.LUNA, terra.UST, terra.ASTRO_LUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNAUST pair")
	}
	BLUNAUST, err := NewXykPair(querier, "terra1wdwg06ksy3dfvkys32yt4yqh9gm6a9f7qmsh37", terra.BLUNA, terra.UST, terra.ASTRO_BLUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init BLUNAUST pair")
	}
	ANCUST, err := NewXykPair(querier, "terra1qr2k6yjjd5p2kaewqvg93ag74k6gyjr7re37fs", terra.ANC, terra.UST, terra.ASTRO_ANCUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ANCUST pair")
	}
	MIRUST, err := NewXykPair(querier, "terra143xxfw5xf62d5m32k3t4eu9s82ccw80lcprzl9", terra.MIR, terra.UST, terra.ASTRO_MIRUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MIRUST pair")
	}
	MINEUST, err := NewXykPair(querier, "terra134m8n2epp0n40qr08qsvvrzycn2zq4zcpmue48", terra.MINE, terra.UST, terra.ASTRO_MINEUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MINEUST pair")
	}
	SKUJIKUJI, err := NewXykPair(querier, "terra1hlq6ye6km5sq2pcnmrvlf784gs9zygt0akwvsu", terra.SKUJI, terra.KUJI, terra.ASTRO_SKUJIKUJILP)
	if err != nil {
		return nil, errors.Wrap(err, "init SKUJIKUJI pair")
	}
	MARSUST, err := NewXykPair(querier, "terra19wauh79y42u5vt62c5adt2g5h4exgh26t3rpds", terra.MARS, terra.UST, terra.ASTRO_MARSUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MARSUST pair")
	}
	ASTROUST, err := NewXykPair(querier, "terra1l7xu2rl3c7qmtx3r5sd2tz25glf6jh8ul7aag7", terra.ASTRO, terra.UST, terra.ASTRO_ASTROUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ASTROUST pair")
	}
	ASTROLUNA, err := NewXykPair(querier, "terra1nujm9zqa4hpaz9s8wrhrp86h3m9xwprjt9kmf9", terra.ASTRO, terra.LUNA, terra.ASTRO_ASTROLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init ASTROLUNA pair")
	}

	LUNABLUNA, err := NewStablePair(querier, "terra1j66jatn3k50hjtg2xemnjm8s7y8dws9xqa5y8w", terra.LUNA, terra.BLUNA, terra.ASTRO_LUNABLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNABLUNA pair")
	}
	MARSXMARS, err := NewStablePair(querier, "terra1dawj5mr2qt2nlurge30lfgjg6ly4ls99yeyd25", terra.MARS, terra.XMARS, terra.ASTRO_MARSXMARSLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MARSXMARS pair")
	}

	r.SetPairs(
		LUNAUST,
		BLUNAUST,
		ANCUST,
		MIRUST,
		MINEUST,
		SKUJIKUJI,
		MARSUST,
		ASTROUST,
		ASTROLUNA,
		LUNABLUNA,
		MARSXMARS,
		//{terra.VKR, terra.UST},
		//{terra.APOLLO, terra.UST},
		//{terra.ORION, terra.UST},
		//{terra.BLUNA, terra.LUNA},
		//{terra.STLUNA, terra.LUNA},
		//{terra.STT, terra.UST},
		//{terra.PSI, terra.UST},
		//{terra.PSI, terra.NLUNA},
		//{terra.WEWSTETH, terra.UST},
		//{terra.PSI, terra.NETH},
		//{terra.XDEFI, terra.UST},
		//{terra.LUART, terra.UST},
		//{terra.ORNE, terra.UST},
		//{terra.HALO, terra.UST},
	)

	return &router{r}, nil
}

type operation struct {
	Swap struct {
		OfferAssetInfo terra.AssetInfo `json:"offer_asset_info"`
		AskAssetInfo   terra.AssetInfo `json:"ask_asset_info"`
	} `json:"astro_swap"`
}

func newOperation(aiFactory terra.AssetInfoFactory, offer terra.Token, ask terra.Token) interface{} {
	var res operation
	res.Swap.OfferAssetInfo = aiFactory.NewFromToken(offer)
	res.Swap.AskAssetInfo = aiFactory.NewFromToken(ask)
	return res
}

func (r router) String() string {
	return "astroport"
}
