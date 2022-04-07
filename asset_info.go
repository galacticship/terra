package terra

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

type AssetInfo interface {
	IsNative() bool
	Id() string
}

func GetTokenFromAssetInfo(ctx context.Context, querier *Querier, ai AssetInfo) (Token, error) {
	var res Token
	if ai.IsNative() {
		return NativeTokenFromDenom(ai.Id()), nil
	}
	var err error
	res, err = Cw20TokenFromAddress(ctx, querier, ai.Id())
	if err != nil {
		return nil, errors.Wrapf(err, "invalid token %s", ai.Id())
	}
	return res, nil
}

type AssetInfoFactory interface {
	DecodeFromJson(raw json.RawMessage) (AssetInfo, error)
	NewFromToken(token Token) AssetInfo
}

type assetInfoFactory struct {
}

func (a assetInfoFactory) NewFromToken(token Token) AssetInfo {
	var res StandardAssetInfo
	if token.IsNative() {
		res.NativeToken = &nativeTokenAssetInfo{
			Denom: token.Id(),
		}
	} else {
		res.Token = &cw20TokenAssetInfo{
			ContractAddr: token.Id(),
		}
	}
	return res
}

func (a assetInfoFactory) DecodeFromJson(raw json.RawMessage) (AssetInfo, error) {
	var res StandardAssetInfo
	err := json.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewAssetInfoFactory() AssetInfoFactory {
	return &assetInfoFactory{}
}

type nativeTokenAssetInfo struct {
	Denom string `json:"denom"`
}

type cw20TokenAssetInfo struct {
	ContractAddr string `json:"contract_addr"`
}

type StandardAssetInfo struct {
	Token       *cw20TokenAssetInfo   `json:"token,omitempty"`
	NativeToken *nativeTokenAssetInfo `json:"native_token,omitempty"`
}

func (ai StandardAssetInfo) IsNative() bool {
	return ai.NativeToken != nil
}

func (ai StandardAssetInfo) Id() string {
	if ai.IsNative() {
		return ai.NativeToken.Denom
	} else if ai.Token != nil {
		return ai.Token.ContractAddr
	} else {
		panic(errors.New("invalid asset info"))
	}
}
