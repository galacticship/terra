# Terra

---

A terra client with some protocol partial implementations (anchor, prism, terraswap type routers, ...)

To be able to compile, you need to add the following replace in your go.mod :

```
replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/cosmos/cosmos-sdk => github.com/terra-money/cosmos-sdk v0.44.5-terra.2
	github.com/cosmos/ledger-cosmos-go => github.com/terra-money/ledger-terra-go v0.11.2
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tecbot/gorocksdb => github.com/cosmos/gorocksdb v1.2.0
	github.com/tendermint/tendermint => github.com/terra-money/tendermint v0.34.14-terra.2
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
```


## Examples

---

### Querying Anchor borrow limit

```
    ctx := context.Background()
	
    querier := terra.NewQuerier(
		httpClient := &http.Client{
            Timeout: 30 * time.Second,
		}, 
		"https://lcd.terra.dev")
	
    walletAddress, err := cosmos.AccAddressFromBech32("walletAddress")

    anc, err := anchor.NewAnchor(querier)
    if err != nil {
        panic(err)
    }
    borrowLimit, err := anc.Overseer.BorrowLimit(ctx, walletAddress)
    if err != nil {
        panic(err)
    }
	
```


### Depositing UST to Anchor
```
    ctx := context.Background()
	
    querier := terra.NewQuerier(
		httpClient := &http.Client{
            Timeout: 30 * time.Second,
		}, 
		"https://lcd.terra.dev")
	
    wallet, err := terra.NewWalletFromMnemonic(
        querier,
        "mnemonic",
        0,
        0)
    if err != nil {
        panic(err)
    }

    anc, err := anchor.NewAnchor(querier)
    if err != nil {
        panic(err)
    }
    err = terra.NewTransaction(querier).
        Message(func() (cosmos.Msg, error) {
            return anc.Market.NewDepositUSTMessage(wallet.Address(), decimal.NewFromInt(100))
        }).
        ExecuteAndWaitFor(ctx, wallet)
	
	
```



