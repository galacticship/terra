package terra

import (
	"context"
	"fmt"
	"time"

	"github.com/galacticship/terra/cosmos"
	"github.com/galacticship/terra/crypto"
	"github.com/galacticship/terra/terra"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type Transaction struct {
	builder cosmos.TxBuilder
	config  cosmos.TxConfig

	q      *Querier
	errors *multierror.Error

	gasLimit      uint64
	feeAmount     cosmos.Coins
	signMode      cosmos.SignMode
	accountNumber uint64
	sequence      uint64
	messages      []cosmos.Msg
}

func NewTransaction(q *Querier) *Transaction {
	return &Transaction{
		builder: terra.EncodingConfig.TxConfig.NewTxBuilder(),
		config:  terra.EncodingConfig.TxConfig,
		q:       q,
	}
}

func (t *Transaction) Error() error {
	return t.errors
}

func (t *Transaction) Message(message func() (cosmos.Msg, error)) *Transaction {
	m, err := message()
	if err != nil {
		t.errors = multierror.Append(t.errors, errors.Wrap(err, "generating message"))
	}
	t.messages = append(t.messages, m)
	return t
}

func (t *Transaction) Memo(memo string) *Transaction {
	t.builder.SetMemo(memo)
	return t
}

func (t *Transaction) FeeGranter(feeGranter cosmos.AccAddress) *Transaction {
	t.builder.SetFeeGranter(feeGranter)
	return t
}

func (t *Transaction) TimeoutHeight(timeoutHeight uint64) *Transaction {
	t.builder.SetTimeoutHeight(timeoutHeight)
	return t
}

func (t *Transaction) GasLimit(gasLimit uint64) *Transaction {
	t.gasLimit = gasLimit
	return t
}

func (t *Transaction) FeeAmount(feeAmount cosmos.Coins) *Transaction {
	t.feeAmount = feeAmount
	return t
}

func (t *Transaction) SignMode(signMode cosmos.SignMode) *Transaction {
	t.signMode = signMode
	return t
}

func (t *Transaction) AccountNumber(accountNumber uint64) *Transaction {
	t.accountNumber = accountNumber
	return t
}

func (t *Transaction) Sequence(sequence uint64) *Transaction {
	t.sequence = sequence
	return t
}

func (t *Transaction) simulate(ctx context.Context) (*cosmos.SimulateResponse, error) {
	sig := cosmos.SignatureV2{
		PubKey: &crypto.PubKey{},
		Data: &cosmos.SingleSignatureData{
			SignMode: t.signMode,
		},
		Sequence: t.sequence,
	}
	if err := t.builder.SetSignatures(sig); err != nil {
		return nil, err
	}

	txBytes, err := t.GetTxBytes()
	if err != nil {
		return nil, err
	}

	var res cosmos.SimulateResponse
	err = t.q.POSTProto(ctx, "cosmos/tx/v1beta1/simulate", cosmos.NewSimulateRequest(txBytes), &res)

	if err != nil {
		return nil, errors.Wrap(err, "querying")
	}
	return &res, nil
}

func (t *Transaction) computeTax(ctx context.Context) (*terra.ComputeTaxResponse, error) {
	txBytes, err := t.GetTxBytes()
	if err != nil {
		return nil, errors.Wrap(err, "getting transaction bytes")
	}
	var res terra.ComputeTaxResponse
	err = t.q.POSTProto(ctx, "terra/tx/v1beta1/compute_tax", terra.NewComputeTaxRequest(txBytes), &res)
	if err != nil {
		return nil, errors.Wrap(err, "querying")
	}
	return &res, nil
}

func (t *Transaction) validate(ctx context.Context, wallet *Wallet) error {
	err := t.builder.SetMsgs(t.messages...)
	if err != nil {
		t.errors = multierror.Append(t.errors, errors.Wrap(err, "setting messages"))
	}

	if t.errors.ErrorOrNil() != nil {
		return t.errors.ErrorOrNil()
	}

	if t.accountNumber == 0 || t.sequence == 0 {
		state, err := wallet.State(ctx)
		if err != nil {
			return errors.Wrap(err, "getting wallet state")
		}
		t.accountNumber = state.AccountNumber
		t.sequence = state.Sequence
	}

	if t.signMode == cosmos.SignModeUnspecified {
		t.signMode = cosmos.SignModeDirect
	}

	gasLimit := int64(t.gasLimit)
	if gasLimit == 0 {
		simulateRes, err := t.simulate(ctx)
		if err != nil {
			return errors.Wrap(err, "simulating transaction for gas limit")
		}
		gasLimit = wallet.GasAdjustment().MulInt64(int64(simulateRes.GasInfo.GasUsed)).Ceil().RoundInt64()
	}
	t.builder.SetGasLimit(uint64(gasLimit))

	feeAmount := t.feeAmount
	if feeAmount.IsZero() {
		//computeTaxRes, err := t.computeTax(ctx)
		//if err != nil {
		//	return errors.Wrap(err, "computing taxes to determine feeAmount")
		//}
		gasPrice := wallet.GasPrice()
		feeAmount = cosmos.NewCoins(cosmos.NewCoin(gasPrice.Denom, gasPrice.Amount.MulInt64(gasLimit).Ceil().RoundInt()))
	}
	t.builder.SetFeeAmount(feeAmount)
	return nil
}

func (t *Transaction) broadcast(ctx context.Context) (*cosmos.TxResponse, error) {
	txBytes, err := t.GetTxBytes()
	if err != nil {
		return nil, err
	}

	var res cosmos.BroadcastTxResponse
	err = t.q.POSTProto(ctx, "cosmos/tx/v1beta1/txs", cosmos.NewBroadcastTxRequest(txBytes, cosmos.BroadcastModeAsync), &res)
	if err != nil {
		return nil, errors.Wrap(err, "querying")
	}
	txResponse := res.TxResponse
	if txResponse.Code != 0 {
		return txResponse, errors.Errorf("tx failed with code %d: %s", txResponse.Code, txResponse.RawLog)
	}
	return txResponse, nil
}

func (t *Transaction) ExecuteAndWaitFor(ctx context.Context, wallet *Wallet) error {
	wallet.lock()
	defer wallet.unlock()
	err := t.validate(ctx, wallet)
	if err != nil {
		return errors.Wrap(err, "validating transaction")
	}
	err = wallet.SignTransaction(t)
	if err != nil {
		return errors.Wrap(err, "signing transaction")
	}
	transresp, err := t.broadcast(ctx)
	if err != nil {
		return errors.Wrap(err, "broadcasting transaction")
	}
	tick := time.NewTicker(2 * time.Second)
	notfoundmax := 10
	notfoundcount := 0
	for {
		select {
		case <-ctx.Done():
			return errors.New("context canceled")
		case <-tick.C:
			var res struct {
				TxResponse struct {
					Height int64  `json:"height,string"`
					Txhash string `json:"txhash"`
					Code   int    `json:"code"`
				} `json:"tx_response"`
			}
			err := t.q.GET(ctx, fmt.Sprintf("cosmos/tx/v1beta1/txs/%s", transresp.TxHash), nil, &res)
			if err != nil {
				if notfoundcount < notfoundmax {
					notfoundcount++
					continue
				}
				return errors.Wrapf(err, "retrieving transaction state for hash %s", transresp.TxHash)
			}
			if res.TxResponse.Code != 0 {
				return errors.Errorf("transaction %s failed with code %d", transresp.TxHash, res.TxResponse.Code)
			}
			t.waitForBlock(ctx, res.TxResponse.Height)
			return nil
		}
	}
}

func (t *Transaction) waitForBlock(ctx context.Context, height int64) {
	checkBlock := func(height int64) error {
		latestBlockHeight, _, err := t.q.LatestBlockInfo(ctx)
		if err != nil {
			return errors.Wrap(err, "querying latest block")
		}
		if latestBlockHeight < height {
			return errors.Wrap(err, "latest block height is less than asked height")
		}
		return nil
	}
	if err := checkBlock(height); err != nil {
		tickHeight := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case <-tickHeight.C:
				err = checkBlock(height)
				if err != nil {
					continue
				}
			}
		}
	}
}

func (t *Transaction) sign(
	signMode cosmos.SignMode, signerData cosmos.SignerData,
	privKey crypto.PrivKey, overwriteSig bool) error {

	sigData := cosmos.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := cosmos.SignatureV2{
		PubKey:   privKey.PubKey(),
		Data:     &sigData,
		Sequence: signerData.Sequence,
	}

	var err error
	var prevSignatures []cosmos.SignatureV2
	if !overwriteSig {
		prevSignatures, err = t.builder.GetTx().GetSignaturesV2()
		if err != nil {
			return err
		}
	}

	if err := t.builder.SetSignatures(sig); err != nil {
		return err
	}

	signature, err := cosmos.SignWithPrivKey(
		signMode,
		signerData,
		t.builder,
		privKey,
		t.config,
		signerData.Sequence,
	)

	if err != nil {
		return err
	}

	if overwriteSig {
		return t.builder.SetSignatures(signature)
	}
	prevSignatures = append(prevSignatures, signature)
	return t.builder.SetSignatures(prevSignatures...)
}

// GetTxBytes return tx bytes for broadcast
func (t Transaction) GetTxBytes() ([]byte, error) {
	return t.config.TxEncoder()(t.builder.GetTx())
}
