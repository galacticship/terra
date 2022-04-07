package terra

import (
	"encoding/json"

	"github.com/galacticship/terra/cosmos"
	markettypes "github.com/terra-money/core/x/market/types"
	wasmtypes "github.com/terra-money/core/x/wasm/types"
)

type (
	MsgSwap                = markettypes.MsgSwap
	MsgSwapSend            = markettypes.MsgSwapSend
	MsgStoreCode           = wasmtypes.MsgStoreCode
	MsgMigrateCode         = wasmtypes.MsgMigrateCode
	MsgInstantiateContract = wasmtypes.MsgInstantiateContract
	MsgExecuteContract     = wasmtypes.MsgExecuteContract
	MsgMigrateContract     = wasmtypes.MsgMigrateContract
)

var (
	NewMsgSwap                = markettypes.NewMsgSwap
	NewMsgSwapSend            = markettypes.NewMsgSwapSend
	NewMsgStoreCode           = wasmtypes.NewMsgStoreCode
	NewMsgMigrateCode         = wasmtypes.NewMsgMigrateCode
	NewMsgInstantiateContract = wasmtypes.NewMsgInstantiateContract
	NewMsgMigrateContract     = wasmtypes.NewMsgMigrateContract
	NewMsgExecuteContract     = func(sender cosmos.AccAddress, contract cosmos.AccAddress, execMsg interface{}, coins cosmos.Coins) (*MsgExecuteContract, error) {
		jsonq, err := json.Marshal(execMsg)
		if err != nil {
			return nil, err
		}
		return wasmtypes.NewMsgExecuteContract(sender, contract, jsonq, coins), nil
	}
)
