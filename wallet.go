package terra

import (
	"context"
	"fmt"
	"sync"

	"github.com/galacticship/terra/cosmos"
	"github.com/galacticship/terra/crypto"
	"github.com/pkg/errors"
)

type Wallet struct {
	q               *Querier
	privKey         crypto.PrivKey
	transactionLock *sync.Mutex

	gasAdjustment cosmos.Dec
	gasPrice      cosmos.DecCoin
}

type WalletOption func(w *Wallet) *Wallet

func WithGasAdjustment(gasAdjustment cosmos.Dec) WalletOption {
	return func(w *Wallet) *Wallet {
		w.gasAdjustment = gasAdjustment
		return w
	}
}

func WithGasPrice(gasPrice cosmos.DecCoin) WalletOption {
	return func(w *Wallet) *Wallet {
		w.gasPrice = gasPrice
		return w
	}
}

func NewWalletFromMnemonic(querier *Querier, mnemonic string, account uint32, index uint32, options ...WalletOption) (*Wallet, error) {
	privKeyBz, err := crypto.DerivePrivKeyBz(mnemonic, crypto.CreateHDPath(account, index))
	if err != nil {
		return nil, errors.Wrap(err, "deriving private key bytes")
	}
	privKey, err := crypto.PrivKeyGen(privKeyBz)
	if err != nil {
		return nil, errors.Wrap(err, "generating private key")
	}

	return NewWalletFromPrivateKey(querier, privKey, options...), nil
}

func NewWalletFromPrivateKey(querier *Querier, privateKey crypto.PrivKey, options ...WalletOption) *Wallet {
	w := &Wallet{
		q:               querier,
		privKey:         privateKey,
		transactionLock: &sync.Mutex{},

		gasAdjustment: cosmos.NewDecFromIntWithPrec(cosmos.NewInt(14), 1),
		gasPrice:      cosmos.NewDecCoinFromDec("uusd", cosmos.NewDecFromIntWithPrec(cosmos.NewInt(15), 2)),
	}
	for _, option := range options {
		w = option(w)
	}
	return w
}

func (a Wallet) GasAdjustment() cosmos.Dec {
	return a.gasAdjustment
}
func (a Wallet) GasPrice() cosmos.DecCoin {
	return a.gasPrice
}

func (a Wallet) Address() cosmos.AccAddress {
	return cosmos.AccAddress(a.privKey.PubKey().Address())
}

type WalletState struct {
	AccountNumber uint64 `json:"account_number,string"`
	Sequence      uint64 `json:"sequence,string"`
}

func (a Wallet) State(ctx context.Context) (WalletState, error) {
	var response struct {
		AccountInfo WalletState `json:"account"`
	}
	err := a.q.GET(ctx, fmt.Sprintf("cosmos/auth/v1beta1/accounts/%s", a.Address().String()), nil, &response)
	if err != nil {
		return WalletState{}, errors.Wrap(err, "querying lcd")
	}
	return response.AccountInfo, nil
}

func (a Wallet) SignTransaction(transaction *Transaction) error {
	err := transaction.sign(transaction.signMode, cosmos.SignerData{
		AccountNumber: transaction.accountNumber,
		ChainID:       a.q.ChainId(),
		Sequence:      transaction.sequence,
	}, a.privKey, true)
	if err != nil {
		return errors.Wrap(err, "signing transaction")
	}
	return nil
}

func (a Wallet) lock() {
	a.transactionLock.Lock()
}

func (a Wallet) unlock() {
	a.transactionLock.Unlock()
}
