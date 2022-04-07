package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

type (
	TxBuilder = client.TxBuilder
	TxConfig  = client.TxConfig

	SimulateRequest  = txtypes.SimulateRequest
	SimulateResponse = txtypes.SimulateResponse

	BroadcastTxRequest  = txtypes.BroadcastTxRequest
	BroadcastTxResponse = txtypes.BroadcastTxResponse

	TxResponse = types.TxResponse

	BroadcastMode = txtypes.BroadcastMode
)

const (
	BroadcastModeUnspecified BroadcastMode = 0
	BroadcastModeBlock       BroadcastMode = 1
	BroadcastModeSync        BroadcastMode = 2
	BroadcastModeAsync       BroadcastMode = 3
)

func NewBroadcastTxRequest(txBytes []byte, broadcastMode BroadcastMode) *BroadcastTxRequest {
	return &BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    broadcastMode,
	}
}

func NewSimulateRequest(txBytes []byte) *SimulateRequest {
	return &SimulateRequest{
		TxBytes: txBytes,
	}
}
