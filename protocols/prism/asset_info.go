package prism

import (
	"encoding/json"
	"errors"

	"github.com/galacticship/terra"
)

//goland:noinspection GoNameStartsWithPackageName
type PrismAssetInfo struct {
	Token       string `json:"cw20,omitempty"`
	NativeToken string `json:"native,omitempty"`
}

func (ai PrismAssetInfo) IsNative() bool {
	return ai.NativeToken != ""
}

func (ai PrismAssetInfo) Id() string {
	if ai.IsNative() {
		return ai.NativeToken
	} else if ai.Token != "" {
		return ai.Token
	} else {
		panic(errors.New("invalid asset info"))
	}
}

type assetInfoFactory struct {
}

func (a assetInfoFactory) NewFromToken(token terra.Token) terra.AssetInfo {
	var res PrismAssetInfo
	if token.IsNative() {
		res.NativeToken = token.Id()
	} else {
		res.Token = token.Id()
	}
	return res
}

func (a assetInfoFactory) DecodeFromJson(raw json.RawMessage) (terra.AssetInfo, error) {
	var res PrismAssetInfo
	err := json.Unmarshal(raw, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewAssetInfoFactory() terra.AssetInfoFactory {
	return &assetInfoFactory{}
}
