package terra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/galacticship/terra/cosmos"
	"github.com/galacticship/terra/terra"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type tokenInfo struct {
	Name        string          `json:"name"`
	Symbol      string          `json:"symbol"`
	Decimals    uint8           `json:"decimals"`
	TotalSupply decimal.Decimal `json:"total_supply"`
}

type Token interface {
	Id() string
	Address() cosmos.AccAddress
	Symbol() string
	Decimals() uint8
	Balance(context.Context, *Querier, cosmos.AccAddress) (decimal.Decimal, error)

	IsNative() bool

	ValueFromTerra(value decimal.Decimal) decimal.Decimal
	ValueToTerra(value decimal.Decimal) decimal.Decimal

	NewMsgSendExecute(sender cosmos.AccAddress, contract *Contract, amount decimal.Decimal, execMsg interface{}) (cosmos.Msg, error)

	Equals(Token) bool

	String() string
}

type Cw20Token struct {
	address  cosmos.AccAddress
	symbol   string
	decimals uint8
}

func NewCw20Token(contractAddress string, symbol string, decimals uint8) (Cw20Token, error) {
	accaddress, err := cosmos.AccAddressFromBech32(contractAddress)
	if err != nil {
		return Cw20Token{}, errors.Wrap(err, "validating token contract address")
	}
	return Cw20Token{
		address:  accaddress,
		symbol:   symbol,
		decimals: decimals,
	}, nil
}

func Cw20TokenFromAddress(ctx context.Context, querier *Querier, contractAddress string) (Cw20Token, error) {
	if t, ok := Cw20TokenMap[contractAddress]; ok {
		return t, nil
	}
	accaddress, err := cosmos.AccAddressFromBech32(contractAddress)
	if err != nil {
		return Cw20Token{}, errors.Wrap(err, "validating token contract address")
	}
	t := Cw20Token{
		address: accaddress,
	}
	ti, err := t.getTokenInfo(ctx, querier)
	if err != nil {
		return Cw20Token{}, errors.Wrap(err, "getting token info")
	}
	t.symbol = ti.Symbol
	t.decimals = ti.Decimals

	return t, nil
}

func (t Cw20Token) getTokenInfo(ctx context.Context, querier *Querier) (tokenInfo, error) {
	var ti tokenInfo
	query := struct {
		TokenInfo struct{} `json:"token_info"`
	}{}
	contract, err := NewContract(querier, t.address.String())
	if err != nil {
		return tokenInfo{}, errors.Wrap(err, "creating contract object")
	}
	err = contract.QueryStore(ctx, query, &ti)
	if err != nil {
		return tokenInfo{}, errors.Wrap(err, "calling token_info contract method")
	}
	return ti, nil
}

func (t Cw20Token) Id() string {
	return t.Address().String()
}

func (t Cw20Token) Decimals() uint8 {
	return t.decimals
}

func (t Cw20Token) DecimalsAsInt32() int32 {
	return int32(t.decimals)
}

func (t Cw20Token) Symbol() string {
	return t.symbol
}

func (t Cw20Token) Address() cosmos.AccAddress {
	return t.address
}

func (t Cw20Token) ValueFromTerra(value decimal.Decimal) decimal.Decimal {
	return value.Shift(-t.DecimalsAsInt32())
}

func (t Cw20Token) ValueToTerra(value decimal.Decimal) decimal.Decimal {
	return value.Shift(t.DecimalsAsInt32())
}

func (t Cw20Token) Balance(ctx context.Context, querier *Querier, address cosmos.AccAddress) (decimal.Decimal, error) {
	type query struct {
		Balance struct {
			Address string `json:"address"`
		} `json:"balance"`
	}
	type response struct {
		Balance decimal.Decimal `json:"balance"`
	}
	var q query
	q.Balance.Address = address.String()
	var r response
	contract, err := NewContract(querier, t.address.String())
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "creating contract object")
	}
	err = contract.QueryStore(ctx, q, &r)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "querying contract store")
	}
	return t.ValueFromTerra(r.Balance), nil
}

func (t Cw20Token) IsNative() bool {
	return false
}

func (t Cw20Token) Equals(token Token) bool {
	if !token.IsNative() && t.symbol == token.Symbol() {
		return true
	}
	return false
}

func (t Cw20Token) String() string {
	return t.symbol
}

func (t Cw20Token) NewMsgSendExecute(sender cosmos.AccAddress, contract *Contract, amount decimal.Decimal, execMsg interface{}) (cosmos.Msg, error) {
	jsonexecmsg, err := json.Marshal(execMsg)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling execMsg")
	}
	type query struct {
		Send struct {
			Contract string          `json:"contract"`
			Amount   decimal.Decimal `json:"amount"`
			Message  interface{}     `json:"msg"`
		} `json:"send"`
	}
	var q query
	q.Send.Contract = contract.Address().String()
	q.Send.Amount = t.ValueToTerra(amount)
	q.Send.Message = jsonexecmsg
	return terra.NewMsgExecuteContract(sender, t.Address(), q, nil)
}

type NativeToken struct {
	symbol string
	denom  string
}

func NewNativeToken(symbol string, denom string) NativeToken {
	return NativeToken{
		denom:  denom,
		symbol: symbol,
	}
}

func NativeTokenFromDenom(denom string) NativeToken {
	if t, ok := NativeTokenMap[denom]; ok {
		return t
	}
	return NewNativeToken("", denom)
}

func (n NativeToken) Id() string {
	return n.denom
}

func (n NativeToken) Address() cosmos.AccAddress {
	return cosmos.AccAddress{}
}

func (n NativeToken) Symbol() string {
	return n.symbol
}
func (n NativeToken) Denom() string {
	return n.denom
}

func (n NativeToken) Decimals() uint8 {
	return 6
}

func (n NativeToken) Balance(ctx context.Context, querier *Querier, address cosmos.AccAddress) (decimal.Decimal, error) {
	var response struct {
		Balance struct {
			Denom  string          `json:"denom"`
			Amount decimal.Decimal `json:"amount"`
		} `json:"balance"`
	}
	params := url.Values{}
	params.Set("denom", n.Denom())
	err := querier.GET(ctx, fmt.Sprintf("cosmos/bank/v1beta1/balances/%s/by_denom", address.String()), params, &response)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "executing get request")
	}
	return n.ValueFromTerra(response.Balance.Amount), nil
}

func (n NativeToken) DecimalsAsInt32() int32 {
	return int32(n.Decimals())
}

func (n NativeToken) ValueFromTerra(value decimal.Decimal) decimal.Decimal {
	return value.Shift(-n.DecimalsAsInt32())
}

func (n NativeToken) ValueToTerra(value decimal.Decimal) decimal.Decimal {
	return value.Shift(n.DecimalsAsInt32())
}

func (n NativeToken) IsNative() bool {
	return true
}

func (n NativeToken) Equals(token Token) bool {
	if token.IsNative() && n.symbol == token.Symbol() {
		return true
	}
	return false
}

func (n NativeToken) String() string {
	return n.symbol
}

func (n NativeToken) NewMsgSendExecute(sender cosmos.AccAddress, contract *Contract, amount decimal.Decimal, execMsg interface{}) (cosmos.Msg, error) {
	return terra.NewMsgExecuteContract(sender, contract.contractAddress, execMsg, cosmos.NewCoins(cosmos.NewInt64Coin(n.Id(), n.ValueToTerra(amount).IntPart())))
}
