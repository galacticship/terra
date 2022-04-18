package terraswap

import (
	"github.com/galacticship/terra"
	"github.com/pkg/errors"
)

type Router struct {
	*terra.BaseRouter
}

func NewRouter(querier *terra.Querier) (terra.Router, error) {
	r, err := terra.NewBaseRouter(querier, "terra19qx5xe6q9ll4w0890ux7lv2p4mf3csd4qvt3ex", terra.NewAssetInfoFactory(), newOperation)
	if err != nil {
		return nil, errors.Wrap(err, "creating base router")
	}

	LUNAUST, err := NewPair(querier, "terra1tndcaqxkpc5ce9qee5ggqf430mr2z3pefe5wj6", terra.LUNA, terra.UST, terra.TERRASWAP_LUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNAUST pair")
	}
	BLUNALUNA, err := NewPair(querier, "terra1jxazgm67et0ce260kvrpfv50acuushpjsz2y0p", terra.BLUNA, terra.LUNA, terra.TERRASWAP_BLUNALUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init BLUNALUNA pair")
	}
	LUNALUNAX, err := NewPair(querier, "terra1zrzy688j8g6446jzd88vzjzqtywh6xavww92hy", terra.LUNAX, terra.LUNA, terra.TERRASWAP_LUNAXLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNALUNAX pair")
	}
	LUNAXBLUNA, err := NewPair(querier, "terra1x8h5gan6vey5cz2xfyst74mtqsj7746fqj2hze", terra.LUNAX, terra.BLUNA, terra.TERRASWAP_LUNAXBLUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNAXBLUNA pair")
	}
	LUNAXUST, err := NewPair(querier, "terra1llhpkqd5enjfflt27u3jx0jcp5pdn6s9lfadx3", terra.LUNAX, terra.UST, terra.TERRASWAP_LUNAXUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUNAXUST pair")
	}
	BLUNAUST, err := NewPair(querier, "terra1qpd9n7afwf45rkjlpujrrdfh83pldec8rpujgn", terra.BLUNA, terra.UST, terra.TERRASWAP_BLUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init BLUNAUST pair")
	}
	KUJIUST, err := NewPair(querier, "terra1zkyrfyq7x9v5vqnnrznn3kvj35az4f6jxftrl2", terra.KUJI, terra.UST, terra.TERRASWAP_KUJIUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init KUJIUST pair")
	}
	PLUNAUST, err := NewPair(querier, "terra1hngzkju6egu78eyzzw2fn8el9dnjk3rr704z2f", terra.PLUNA, terra.UST, terra.TERRASWAP_PLUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PLUNAUST pair")
	}
	STLUNAUST, err := NewPair(querier, "terra1de8xa55xm83s3ke0s20fc5pxy7p3cpndmmm7zk", terra.STLUNA, terra.UST, terra.TERRASWAP_STLUNAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init STLUNAUST pair")
	}
	ANCUST, err := NewPair(querier, "terra1gm5p3ner9x9xpwugn9sp6gvhd0lwrtkyrecdn3", terra.ANC, terra.UST, terra.TERRASWAP_ANCUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ANCUST pair")
	}
	MIRUST, err := NewPair(querier, "terra1amv303y8kzxuegvurh0gug2xe9wkgj65enq2ux", terra.MIR, terra.UST, terra.TERRASWAP_MIRUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MIRUST pair")
	}
	LOOPUST, err := NewPair(querier, "terra10k7y9qw63tfwj7e3x4uuzru2u9kvtd4ureajhd", terra.LOOP, terra.UST, terra.TERRASWAP_LOOPUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LOOPUST pair")
	}
	LOOPRUST, err := NewPair(querier, "terra18raj59xx32kuz66sfg82kqta6q0aslfs3m8s4r", terra.LOOPR, terra.UST, terra.TERRASWAP_LOOPRUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LOOPRUST pair")
	}
	MINEUST, err := NewPair(querier, "terra178jydtjvj4gw8earkgnqc80c3hrmqj4kw2welz", terra.MINE, terra.UST, terra.TERRASWAP_MINEUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MINEUST pair")
	}
	SKUJIKUJI, err := NewPair(querier, "terra1g8kjs70d5r68j9507s3gwymzc30yaur5j2ccfr", terra.SKUJI, terra.KUJI, terra.TERRASWAP_SKUJIKUJILP)
	if err != nil {
		return nil, errors.Wrap(err, "init SKUJIKUJI pair")
	}
	MARSUST, err := NewPair(querier, "terra15sut89ms4lts4dd5yrcuwcpctlep3hdgeu729f", terra.MARS, terra.UST, terra.TERRASWAP_MARSUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init MARSUST pair")
	}
	PRISMXPRISM, err := NewPair(querier, "terra1urt608par6rkcancsjzm76472phptfwq397gpm", terra.PRISM, terra.XPRISM, terra.TERRASWAP_PRISMXPRISMLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMXPRISM pair")
	}
	PRISMUST, err := NewPair(querier, "terra1ag6fqvxz33nqg78830k5c27n32mmqlcrcgqejl", terra.PRISM, terra.UST, terra.TERRASWAP_PRISMUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PRISMUST pair")
	}
	CLUNALUNA, err := NewPair(querier, "terra1ejyqwcemr5kda5pxwz27t2ja784j3d0nj0v6lh", terra.CLUNA, terra.LUNA, terra.TERRASWAP_CLUNALUNALP)
	if err != nil {
		return nil, errors.Wrap(err, "init CLUNALUNA pair")
	}
	ASTROUST, err := NewPair(querier, "terra1pufczag48fwqhsmekfullmyu02f93flvfc9a25", terra.ASTRO, terra.UST, terra.TERRASWAP_ASTROUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ASTROUST pair")
	}
	AUSTUST, err := NewPair(querier, "terra1z50zu7j39s2dls8k9xqyxc89305up0w7f7ec3n", terra.AUST, terra.UST, terra.TERRASWAP_AUSTUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init AUSTUST pair")
	}
	AUSTVUST, err := NewPair(querier, "terra1gkdudgg2a5wt70cneyx5rtehjls4dvhhcmlptv", terra.AUST, terra.VUST, terra.TERRASWAP_AUSTVUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init AUSTVUST pair")
	}
	WHALEVUST, err := NewPair(querier, "terra12arl49w7t4xpq7krtv43t3dg6g8kn2xxyaav56", terra.WHALE, terra.VUST, terra.TERRASWAP_WHALEVUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init WHALEVUST pair")
	}
	BETHUST, err := NewPair(querier, "terra1c0afrdc5253tkp5wt7rxhuj42xwyf2lcre0s7c", terra.BETH, terra.UST, terra.TERRASWAP_BETHUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init BETHUST pair")
	}
	WHALEUST, err := NewPair(querier, "terra1v4kpj65uq63m4x0mqzntzm27ecpactt42nyp5c", terra.WHALE, terra.UST, terra.TERRASWAP_WHALEUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init WHALEUST pair")
	}
	SPECUST, err := NewPair(querier, "terra1tn8ejzw8kpuc87nu42f6qeyen4c7qy35tl8t20", terra.SPEC, terra.UST, terra.TERRASWAP_SPECUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init SPECUST pair")
	}
	STTUST, err := NewPair(querier, "terra19pg6d7rrndg4z4t0jhcd7z9nhl3p5ygqttxjll", terra.STT, terra.UST, terra.TERRASWAP_STTUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init STTUST pair")
	}
	TWDUST, err := NewPair(querier, "terra1etdkg9p0fkl8zal6ecp98kypd32q8k3ryced9d", terra.TWD, terra.UST, terra.TERRASWAP_TWDUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init TWDUST pair")
	}
	PSIUST, err := NewPair(querier, "terra163pkeeuwxzr0yhndf8xd2jprm9hrtk59xf7nqf", terra.PSI, terra.UST, terra.TERRASWAP_PSIUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PSIUST pair")
	}
	PLYUST, err := NewPair(querier, "terra19fjaurx28dq4wgnf9fv3qg0lwldcln3jqafzm6", terra.PLY, terra.UST, terra.TERRASWAP_PLYUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init PLYUST pair")
	}
	LOTAUST, err := NewPair(querier, "terra1pn20mcwnmeyxf68vpt3cyel3n57qm9mp289jta", terra.LOTA, terra.UST, terra.TERRASWAP_LOTAUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LOTAUST pair")
	}
	APOLLOUST, err := NewPair(querier, "terra1xj2w7w8mx6m2nueczgsxy2gnmujwejjeu2xf78", terra.APOLLO, terra.UST, terra.TERRASWAP_APOLLOUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init APOLLOUST pair")
	}
	VKRUST, err := NewPair(querier, "terra1e59utusv5rspqsu8t37h5w887d9rdykljedxw0", terra.VKR, terra.UST, terra.TERRASWAP_VKRUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init VKRUST pair")
	}
	ORIONUST, err := NewPair(querier, "terra1z6tp0ruxvynsx5r9mmcc2wcezz9ey9pmrw5r8g", terra.ORION, terra.UST, terra.TERRASWAP_ORIONUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ORIONUST pair")
	}
	ATLOUST, err := NewPair(querier, "terra1ycp5lnn0qu4sq4wq7k63zax9f05852xt9nu3yc", terra.ATLO, terra.UST, terra.TERRASWAP_ATLOUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ATLOUST pair")
	}
	GLOWUST, err := NewPair(querier, "terra1p44kn7l233p7gcj0v3mzury8k7cwf4zt6gsxs5", terra.GLOW, terra.UST, terra.TERRASWAP_GLOWUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init GLOWUST pair")
	}
	TNSUST, err := NewPair(querier, "terra1hqnk9expq3k4la2ruzdnyapgndntec4fztdyln", terra.TNS, terra.UST, terra.TERRASWAP_TNSUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init TNSUST pair")
	}
	LUVUST, err := NewPair(querier, "terra1hmcd4kwafyydd4mjv2rzhcuuwnfuqc2prkmlhj", terra.LUV, terra.UST, terra.TERRASWAP_LUVUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init LUVUST pair")
	}
	ROBOUST, err := NewPair(querier, "terra1sprg4sv9dwnk78ahxdw78asslj8upyv9lerjhm", terra.ROBO, terra.UST, terra.TERRASWAP_ROBOUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init ROBOUST pair")
	}
	XSDWHSD, err := NewPair(querier, "terra1cmehvqwvglg08clmqn66zfuv5cuxgxwrt3jz2u", terra.XSD, terra.WHSD, terra.TERRASWAP_XSDWHSDLP)
	if err != nil {
		return nil, errors.Wrap(err, "init XSDWHSD pair")
	}
	WHSDUST, err := NewPair(querier, "terra1upuslwv5twc8l7hrwlka4wju9z97q8ju63a6jt", terra.WHSD, terra.UST, terra.TERRASWAP_WHSDUSTLP)
	if err != nil {
		return nil, errors.Wrap(err, "init WHSDUST pair")
	}
	NLUNAPSI, err := NewPair(querier, "terra1zvn8z6y8u2ndwvsjhtpsjsghk6pa6ugwzxp6vx", terra.NLUNA, terra.PSI, terra.TERRASWAP_NLUNAPSILP)
	if err != nil {
		return nil, errors.Wrap(err, "init NLUNAPSI pair")
	}
	XASTROASTRO, err := NewPair(querier, "terra14q2h9nce4spj8n74g6kppj3yf86qx8hsrqngfh", terra.XASTRO, terra.ASTRO, terra.TERRASWAP_XASTROASTROLP)
	if err != nil {
		return nil, errors.Wrap(err, "init NLUNAPSI pair")
	}

	r.SetPairs(
		LUNAUST,
		BLUNALUNA,
		LUNALUNAX,
		LUNAXBLUNA,
		LUNAXUST,
		BLUNAUST,
		KUJIUST,
		PLUNAUST,
		STLUNAUST,
		ANCUST,
		MIRUST,
		LOOPUST,
		LOOPRUST,
		MINEUST,
		SKUJIKUJI,
		MARSUST,
		PRISMXPRISM,
		CLUNALUNA,
		ASTROUST,
		AUSTUST,
		AUSTVUST,
		WHALEVUST,
		BETHUST,
		WHALEUST,
		SPECUST,
		STTUST,
		TWDUST,
		PSIUST,
		PLYUST,
		LOTAUST,
		APOLLOUST,
		VKRUST,
		ORIONUST,
		ATLOUST,
		GLOWUST,
		TNSUST,
		LUVUST,
		ROBOUST,
		XSDWHSD,
		WHSDUST,
		PRISMUST,
		NLUNAPSI,
		XASTROASTRO,
	)

	return &Router{r}, nil
}

type swap struct {
	OfferAssetInfo terra.AssetInfo `json:"offer_asset_info"`
	AskAssetInfo   terra.AssetInfo `json:"ask_asset_info"`
}
type operation struct {
	Swap swap `json:"terra_swap"`
}

func newOperation(aiFactory terra.AssetInfoFactory, offer terra.Token, ask terra.Token) interface{} {
	var res operation
	res.Swap.OfferAssetInfo = aiFactory.NewFromToken(offer)
	res.Swap.AskAssetInfo = aiFactory.NewFromToken(ask)
	return res
}

func (r Router) String() string {
	return "terraswap"
}
