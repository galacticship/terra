package terra

var (
	LUNA = NewNativeToken("LUNA", "uluna")
	UST  = NewNativeToken("UST", "uusd")
)

var (
	SKUJI, _    = NewCw20Token("terra188w26t95tf4dz77raftme8p75rggatxjxfeknw", "sKUJI", 6)
	KUJI, _     = NewCw20Token("terra1xfsdgcemqwxp4hhnyk4rle6wr22sseq7j07dnn", "KUJI", 6)
	XPRISM, _   = NewCw20Token("terra1042wzrwg2uk6jqxjm34ysqquyr9esdgm5qyswz", "xPRISM", 6)
	YLUNA, _    = NewCw20Token("terra17wkadg0tah554r35x6wvff0y5s7ve8npcjfuhz", "yLUNA", 6)
	PLUNA, _    = NewCw20Token("terra1tlgelulz9pdkhls6uglfn5lmxarx7f2gxtdzh2", "pLUNA", 6)
	PRISM, _    = NewCw20Token("terra1dh9478k2qvqhqeajhn75a2a7dsnf74y5ukregw", "PRISM", 6)
	CLUNA, _    = NewCw20Token("terra13zaagrrrxj47qjwczsczujlvnnntde7fdt0mau", "cLUNA", 6)
	ASTRO, _    = NewCw20Token("terra1xj49zyqrwpv5k928jwfpfy2ha668nwdgkwlrg3", "ASTRO", 6)
	APOLLO, _   = NewCw20Token("terra100yeqvww74h4yaejj6h733thgcafdaukjtw397", "APOLLO", 6)
	ANC, _      = NewCw20Token("terra14z56l0fp2lsf86zy3hty2z47ezkhnthtr9yq76", "ANC", 6)
	BLUNA, _    = NewCw20Token("terra1kc87mu460fwkqte29rquh4hc20m54fxwtsx7gp", "bLUNA", 6)
	AUST, _     = NewCw20Token("terra1hzh9vpxhsk8253se0vv5jj6etdvxu3nv8z07zu", "aUST", 6)
	BETH, _     = NewCw20Token("terra1dzhzukyezv0etz22ud940z7adyv7xgcjkahuun", "bETH", 6)
	MIR, _      = NewCw20Token("terra15gwkyepfc6xgca5t5zefzwy42uts8l2m4g40k6", "MIR", 6)
	MINE, _     = NewCw20Token("terra1kcthelkax4j9x8d3ny6sdag0qmxxynl3qtcrpy", "MINE", 6)
	STT, _      = NewCw20Token("terra13xujxcrc9dqft4p9a8ls0w3j0xnzm6y2uvve8n", "STT", 6)
	PSI, _      = NewCw20Token("terra12897djskt9rge8dtmm86w654g7kzckkd698608", "PSI", 6)
	VKR, _      = NewCw20Token("terra1dy9kmlm4anr92e42mrkjwzyvfqwz66un00rwr5", "VKR", 6)
	SPEC, _     = NewCw20Token("terra1s5eczhe0h0jutf46re52x5z4r03c8hupacxmdr", "SPEC", 6)
	ORION, _    = NewCw20Token("terra1mddcdx0ujx89f38gu7zspk2r2ffdl5enyz2u03", "ORION", 8)
	GLOW, _     = NewCw20Token("terra13zx49nk8wjavedjzu8xkk95r3t0ta43c9ptul7", "GLOW", 6)
	HALO, _     = NewCw20Token("terra1w8kvd6cqpsthupsk4l0clwnmek4l3zr7c84kwq", "HALO", 6)
	LOOP, _     = NewCw20Token("terra1nef5jf6c7js9x6gkntlehgywvjlpytm7pcgkn4", "LOOP", 6)
	PLY, _      = NewCw20Token("terra13awdgcx40tz5uygkgm79dytez3x87rpg4uhnvu", "PLY", 6)
	WHALE, _    = NewCw20Token("terra1php5m8a6qd68z02t3zpw4jv2pj4vgw4wz0t8mz", "WHALE", 6)
	MARS, _     = NewCw20Token("terra12hgwnpupflfpuual532wgrxu2gjp0tcagzgx4n", "MARS", 6)
	ATLO, _     = NewCw20Token("terra1cl7whtrqmz5ldr553q69qahck8xvk80fm33qjx", "ATLO", 6)
	LOTA, _     = NewCw20Token("terra1ez46kxtulsdv07538fh5ra5xj8l68mu8eg24vr", "LOTA", 6)
	TWD, _      = NewCw20Token("terra19djkaepjjswucys4npd5ltaxgsntl7jf0xz7w6", "TWD", 6)
	LUNAX, _    = NewCw20Token("terra17y9qkl8dfkeg4py7n0g5407emqnemc3yqk5rup", "LunaX", 6)
	VUST, _     = NewCw20Token("terra1w0p5zre38ecdy3ez8efd5h9fvgum5s206xknrg", "vUST", 6)
	STLUNA, _   = NewCw20Token("terra1yg3j2s986nyp5z7r2lvt0hx3r0lnd7kwvwwtsc", "stLUNA", 6)
	NLUNA, _    = NewCw20Token("terra10f2mt82kjnkxqj2gepgwl637u2w4ue2z5nhz5j", "nLUNA", 6)
	WEWSTETH, _ = NewCw20Token("terra133chr09wu8sakfte5v7vd8qzq9vghtkv4tn0ur", "wewstETH", 8)
	NETH, _     = NewCw20Token("terra178v546c407pdnx5rer3hu8s2c0fc924k74ymnn", "nETH", 6)
	XDEFI, _    = NewCw20Token("terra169edevav3pdrtjcx35j6pvzuv54aevewar4nlh", "XDEFI", 8)
	LUART, _    = NewCw20Token("terra1vwz7t30q76s7xx6qgtxdqnu6vpr3ak3vw62ygk", "XDEFI", 6)
	ORNE, _     = NewCw20Token("terra1hnezwjqlhzawcrfysczcxs6xqxu2jawn729kkf", "ORNE", 6)
	LOOPR, _    = NewCw20Token("terra1jx4lmmke2srcvpjeereetc9hgegp4g5j0p9r2q", "LOOPR", 6)
	TNS, _      = NewCw20Token("terra14vz4v8adanzph278xyeggll4tfww7teh0xtw2y", "TNS", 6)
	LUV, _      = NewCw20Token("terra15k5r9r8dl8r7xlr29pry8a9w7sghehcnv5mgp6", "LUV", 6)
	ROBO, _     = NewCw20Token("terra1f62tqesptvmhtzr8sudru00gsdtdz24srgm7wp", "ROBO", 6)
	XSD, _      = NewCw20Token("terra1ln2z938phz0nc2wepxpzfkwp6ezn9yrz9zv9ep", "XSD", 8)
	WHSD, _     = NewCw20Token("terra1ustvnmngueq0p4jd7gfnutgvdc6ujpsjhsjd02", "WHSD", 8)
	XASTRO, _   = NewCw20Token("terra14lpnyzc9z4g3ugr4lhm8s4nle0tq8vcltkhzh7", "xASTRO", 6)
)

var (
	ASTRO_LUNAUSTLP, _   = NewCw20Token("terra1m24f7k4g66gnh9f7uncp32p722v0kyt3q4l3u5", "uLP", 6)
	ASTRO_BLUNAUSTLP, _  = NewCw20Token("terra1aaqmlv4ajsg9043zrhsd44lk8dqnv2hnakjv97", "uLP", 6)
	ASTRO_ANCUSTLP, _    = NewCw20Token("terra1wmaty65yt7mjw6fjfymkd9zsm6atsq82d9arcd", "uLP", 6)
	ASTRO_MIRUSTLP, _    = NewCw20Token("terra17trxzqjetl0q6xxep0s2w743dhw2cay0x47puc", "uLP", 6)
	ASTRO_MINEUSTLP, _   = NewCw20Token("terra16unvjel8vvtanxjpw49ehvga5qjlstn8c826qe", "uLP", 6)
	ASTRO_SKUJIKUJILP, _ = NewCw20Token("terra1kp4n4tms5w4tvvypya7589zswssqqahtjxy6da", "uLP", 6)
	ASTRO_MARSUSTLP, _   = NewCw20Token("terra1ww6sqvfgmktp0afcmvg78st6z89x5zr3tmvpss", "uLP", 6)
	ASTRO_ASTROUSTLP, _  = NewCw20Token("terra17n5sunn88hpy965mzvt3079fqx3rttnplg779g", "uLP", 6)
	ASTRO_ASTROLUNALP, _ = NewCw20Token("terra1ryxkslm6p04q0nl046quwz8ctdd5llkjnaccpa", "uLP", 6)
	ASTRO_LUNABLUNALP, _ = NewCw20Token("terra1htw7hm40ch0hacm8qpgd24sus4h0tq3hsseatl", "uLP", 6)

	TERRASWAP_LUNAUSTLP, _     = NewCw20Token("terra17dkr9rnmtmu7x4azrpupukvur2crnptyfvsrvr", "uLP", 6)
	TERRASWAP_BLUNALUNALP, _   = NewCw20Token("terra1nuy34nwnsh53ygpc4xprlj263cztw7vc99leh2", "uLP", 6)
	TERRASWAP_LUNAXLUNALP, _   = NewCw20Token("terra1halhfnaul7c0u9t5aywj430jnlu2hgauftdvdq", "uLP", 6)
	TERRASWAP_LUNAXBLUNALP, _  = NewCw20Token("terra1spagsh9rgcpdgl5pj6lfyftmhtz9elugurfl92", "uLP", 6)
	TERRASWAP_LUNAXUSTLP, _    = NewCw20Token("terra1ah6vn794y3fjvn5jvcv0pzrzky7gar3tp8zuyu", "uLP", 6)
	TERRASWAP_BLUNAUSTLP, _    = NewCw20Token("terra1qmr8j3m9x53dhws0yxhymzsvnkjq886yk8k93m", "uLP", 6)
	TERRASWAP_KUJIUSTLP, _     = NewCw20Token("terra1cmqv3sjew8kcm3j907x2026e4n0ejl2jackxlx", "uLP", 6)
	TERRASWAP_PLUNAUSTLP, _    = NewCw20Token("terra1t5tg7jrmsk6mj9xs3fk0ey092wfkqqlapuevwr", "uLP", 6)
	TERRASWAP_STLUNAUSTLP, _   = NewCw20Token("terra1cksuxx4ryfyhkk2c6lw3mpnn4hahkxlkml82rp", "uLP", 6)
	TERRASWAP_ANCUSTLP, _      = NewCw20Token("terra1gecs98vcuktyfkrve9czrpgtg0m3aq586x6gzm", "uLP", 6)
	TERRASWAP_MIRUSTLP, _      = NewCw20Token("terra17gjf2zehfvnyjtdgua9p9ygquk6gukxe7ucgwh", "uLP", 6)
	TERRASWAP_LOOPUSTLP, _     = NewCw20Token("terra12v03at235nxnmsyfg09akh4tp02mr60ne6flry", "uLP", 6)
	TERRASWAP_LOOPRUSTLP, _    = NewCw20Token("terra17mau5a2q453vf4e33ffaa4cvtn0twle5vm0zuf", "uLP", 6)
	TERRASWAP_MINEUSTLP, _     = NewCw20Token("terra1rqkyau9hanxtn63mjrdfhpnkpddztv3qav0tq2", "uLP", 6)
	TERRASWAP_SKUJIKUJILP, _   = NewCw20Token("terra1qf5xuhns225e6xr3mnjv3z8qwlpzyzf2c8we82", "uLP", 6)
	TERRASWAP_MARSUSTLP, _     = NewCw20Token("terra175xghpferetqhnx0hlp3e0um0wyfknxzv8h42q", "uLP", 6)
	TERRASWAP_PRISMXPRISMLP, _ = NewCw20Token("terra1pc6fvx7vzk220uj840kmkrnyjhjwxcrneuffnk", "uLP", 6)
	TERRASWAP_PRISMUSTLP, _    = NewCw20Token("terra1tragr8vkyx52rzy9f8n42etl6la42zylhcfkwa", "uLP", 6)
	TERRASWAP_CLUNALUNALP, _   = NewCw20Token("terra18cul84v9tt4nmxmmyxm2z74vpgjmrj6py73pus", "uLP", 6)
	TERRASWAP_ASTROUSTLP, _    = NewCw20Token("terra1xwyhu8geetx2mv8429a3z5dyzr0ajqnmmn4rtr", "uLP", 6)
	TERRASWAP_AUSTUSTLP, _     = NewCw20Token("terra1umup8qvslkayek0af8u7x2r3r5ndhk2fwhdxz5", "uLP", 6)
	TERRASWAP_AUSTVUSTLP, _    = NewCw20Token("terra14d33ndaanjc802ural7uq8ck3n6smsy2e4r0rt", "uLP", 6)
	TERRASWAP_WHALEVUSTLP, _   = NewCw20Token("terra1hg3tr0gx2jfaw38m80s83c7khr4wgfvzyh5uak", "uLP", 6)
	TERRASWAP_BETHUSTLP, _     = NewCw20Token("terra1jvewsf7922dm47wr872crumps7ktxd7srwcgte", "uLP", 6)
	TERRASWAP_WHALEUSTLP, _    = NewCw20Token("terra17pqpurglgfqnvkhypask28c3llnf69cstaquke", "uLP", 6)
	TERRASWAP_SPECUSTLP, _     = NewCw20Token("terra1y9kxxm97vu4ex3uy0rgdr5h2vt7aze5sqx7jyl", "uLP", 6)
	TERRASWAP_STTUSTLP, _      = NewCw20Token("terra1uwhf02zuaw7grj6gjs7pxt5vuwm79y87ct5p70", "uLP", 6)
	TERRASWAP_TWDUSTLP, _      = NewCw20Token("terra1c9wr85y8p8989tr58flz5gjkqp8q2r6murwpm9", "uLP", 6)
	TERRASWAP_PSIUSTLP, _      = NewCw20Token("terra1q6r8hfdl203htfvpsmyh8x689lp2g0m7856fwd", "uLP", 6)
	TERRASWAP_PLYUSTLP, _      = NewCw20Token("terra1h69kvcmg8jnq7ph2r45k6md4afkl96ugg73amc", "uLP", 6)
	TERRASWAP_LOTAUSTLP, _     = NewCw20Token("terra1t4xype7nzjxrzttuwuyh9sglwaaeszr8l78u6e", "uLP", 6)
	TERRASWAP_APOLLOUSTLP, _   = NewCw20Token("terra1n3gt4k3vth0uppk0urche6m3geu9eqcyujt88q", "uLP", 6)
	TERRASWAP_VKRUSTLP, _      = NewCw20Token("terra17fysmcl52xjrs8ldswhz7n6mt37r9cmpcguack", "uLP", 6)
	TERRASWAP_ORIONUSTLP, _    = NewCw20Token("terra14ffp0waxcck733a9jfd58d86h9rac2chf5xhev", "uLP", 6)
	TERRASWAP_ATLOUSTLP, _     = NewCw20Token("terra1l0wqwge0vtfmukx028pluxsr7ee2vk8gl5mlxr", "uLP", 6)
	TERRASWAP_GLOWUSTLP, _     = NewCw20Token("terra1khm4az2cjlzl76885x2n7re48l9ygckjuye0mt", "uLP", 6)
	TERRASWAP_TNSUSTLP, _      = NewCw20Token("terra1kg9vmu4e43d3pz0dfsdg9vzwgnnuf6uf3z9jwj", "uLP", 6)
	TERRASWAP_LUVUSTLP, _      = NewCw20Token("terra1qq6g0kds9zn97lvrukf2qxf6w4sjt0k9jhcdty", "uLP", 6)
	TERRASWAP_ROBOUSTLP, _     = NewCw20Token("terra19ryu7a586s75ncw3ddc8julkszjht4ahwd7zja", "uLP", 6)
	TERRASWAP_XSDWHSDLP, _     = NewCw20Token("terra1z0vaks4wkehncztu2a3j2z4fj2gjsnyk2ng9xu", "uLP", 6)
	TERRASWAP_WHSDUSTLP, _     = NewCw20Token("terra13m7t5z9zvx2phtpa0k6lxht3qtjjhj68u0t0jz", "uLP", 6)
	TERRASWAP_NLUNAPSILP, _    = NewCw20Token("terra1tuw46dwfvahpcwf3ulempzsn9a0vhazut87zec", "uLP", 6)
	TERRASWAP_XASTROASTROLP, _ = NewCw20Token("terra1h5egnh0uu4qcjx359fgr5jfytjsazsynhm7lw7", "uLP", 6)

	PRISM_PRISMUSTLP, _    = NewCw20Token("terra1wkv9htanake4yerrrjz8p5n40lyrjg9md28tg3", "uLP", 6)
	PRISM_PRISMLUNALP, _   = NewCw20Token("terra1af7hyx4ek8vqr8asmtujsyv7s3z6py3jgtsgh8", "uLP", 6)
	PRISM_PRISMPLUNALP, _  = NewCw20Token("terra1rjm3ca2xh2cfm6l6nsnvs6dqzed0lgzdydy7wf", "uLP", 6)
	PRISM_PRISMXPRISMLP, _ = NewCw20Token("terra1zuv05w52xvtn3td2lpfl3q9jj807533ew54f0x", "uLP", 6)
	PRISM_PRISMCLUNALP, _  = NewCw20Token("terra1vn5c4yf70aasrq50k2xdy3vn2s8vm40wmngljh", "uLP", 6)
	PRISM_PRISMYLUNALP, _  = NewCw20Token("terra1argcazqn3ukpyp0vmldxnf9qymnm6vfjaar94g", "uLP", 6)
)

var (
	Cw20TokenMap = map[string]Cw20Token{
		"terra188w26t95tf4dz77raftme8p75rggatxjxfeknw": SKUJI,
		"terra1xfsdgcemqwxp4hhnyk4rle6wr22sseq7j07dnn": KUJI,
		"terra1042wzrwg2uk6jqxjm34ysqquyr9esdgm5qyswz": XPRISM,
		"terra17wkadg0tah554r35x6wvff0y5s7ve8npcjfuhz": YLUNA,
		"terra1tlgelulz9pdkhls6uglfn5lmxarx7f2gxtdzh2": PLUNA,
		"terra1dh9478k2qvqhqeajhn75a2a7dsnf74y5ukregw": PRISM,
		"terra13zaagrrrxj47qjwczsczujlvnnntde7fdt0mau": CLUNA,
		"terra1xj49zyqrwpv5k928jwfpfy2ha668nwdgkwlrg3": ASTRO,
		"terra1f68wt2ch3cx2g62dxtc8v68mkdh5wchdgdjwz7": XASTRO,
		"terra100yeqvww74h4yaejj6h733thgcafdaukjtw397": APOLLO,
		"terra14z56l0fp2lsf86zy3hty2z47ezkhnthtr9yq76": ANC,
		"terra1kc87mu460fwkqte29rquh4hc20m54fxwtsx7gp": BLUNA,
		"terra1hzh9vpxhsk8253se0vv5jj6etdvxu3nv8z07zu": AUST,
		"terra1dzhzukyezv0etz22ud940z7adyv7xgcjkahuun": BETH,
		"terra15gwkyepfc6xgca5t5zefzwy42uts8l2m4g40k6": MIR,
		"terra1kcthelkax4j9x8d3ny6sdag0qmxxynl3qtcrpy": MINE,
		"terra13xujxcrc9dqft4p9a8ls0w3j0xnzm6y2uvve8n": STT,
		"terra12897djskt9rge8dtmm86w654g7kzckkd698608": PSI,
		"terra1dy9kmlm4anr92e42mrkjwzyvfqwz66un00rwr5": VKR,
		"terra1s5eczhe0h0jutf46re52x5z4r03c8hupacxmdr": SPEC,
		"terra1mddcdx0ujx89f38gu7zspk2r2ffdl5enyz2u03": ORION,
		"terra13zx49nk8wjavedjzu8xkk95r3t0ta43c9ptul7": GLOW,
		"terra1w8kvd6cqpsthupsk4l0clwnmek4l3zr7c84kwq": HALO,
		"terra1nef5jf6c7js9x6gkntlehgywvjlpytm7pcgkn4": LOOP,
		"terra13awdgcx40tz5uygkgm79dytez3x87rpg4uhnvu": PLY,
		"terra1php5m8a6qd68z02t3zpw4jv2pj4vgw4wz0t8mz": WHALE,
		"terra12hgwnpupflfpuual532wgrxu2gjp0tcagzgx4n": MARS,
		"terra1cl7whtrqmz5ldr553q69qahck8xvk80fm33qjx": ATLO,
		"terra1ez46kxtulsdv07538fh5ra5xj8l68mu8eg24vr": LOTA,
		"terra19djkaepjjswucys4npd5ltaxgsntl7jf0xz7w6": TWD,
		"terra17y9qkl8dfkeg4py7n0g5407emqnemc3yqk5rup": LUNAX,
		"terra1w0p5zre38ecdy3ez8efd5h9fvgum5s206xknrg": VUST,
		"terra1yg3j2s986nyp5z7r2lvt0hx3r0lnd7kwvwwtsc": STLUNA,
		"terra133chr09wu8sakfte5v7vd8qzq9vghtkv4tn0ur": WEWSTETH,
		"terra178v546c407pdnx5rer3hu8s2c0fc924k74ymnn": NETH,
		"terra169edevav3pdrtjcx35j6pvzuv54aevewar4nlh": XDEFI,
		"terra1vwz7t30q76s7xx6qgtxdqnu6vpr3ak3vw62ygk": LUART,
		"terra1hnezwjqlhzawcrfysczcxs6xqxu2jawn729kkf": ORNE,
		"terra1jx4lmmke2srcvpjeereetc9hgegp4g5j0p9r2q": LOOPR,
		"terra14vz4v8adanzph278xyeggll4tfww7teh0xtw2y": TNS,
		"terra15k5r9r8dl8r7xlr29pry8a9w7sghehcnv5mgp6": LUV,
		"terra1f62tqesptvmhtzr8sudru00gsdtdz24srgm7wp": ROBO,
		"terra1ln2z938phz0nc2wepxpzfkwp6ezn9yrz9zv9ep": XSD,
		"terra1ustvnmngueq0p4jd7gfnutgvdc6ujpsjhsjd02": WHSD,
	}
	NativeTokenMap = map[string]NativeToken{
		"uusd":  UST,
		"uluna": LUNA,
	}
)
