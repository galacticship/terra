package cosmos

import (
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type (
	Msg       = sdktypes.Msg
	Send      = banktypes.MsgSend
	MultiSend = banktypes.MsgMultiSend
)

var (
	NewMsgSend      = banktypes.NewMsgSend
	NewMsgMultiSend = banktypes.NewMsgMultiSend
)
