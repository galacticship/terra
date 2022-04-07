package crypto

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/go-bip39"
)

type (
	PubKey  = secp256k1.PubKey
	PrivKey = types.PrivKey
)

// CreateMnemonic - create new mnemonic
func CreateMnemonic() (string, error) {
	// Default number of words (24): This generates a mnemonic directly from the
	// number of words by reading system entropy.
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(entropy)
}

// CreateHDPath returns BIP 44 object from account and index parameters.
func CreateHDPath(account uint32, index uint32) string {
	return hd.CreateHDPath(330, account, index).String()
}

// DerivePrivKeyBz - derive private key bytes
func DerivePrivKeyBz(mnemonic string, hdPath string) ([]byte, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, errors.New("invalid mnemonic")
	}

	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyring.SigningAlgoList{hd.Secp256k1})
	if err != nil {
		return nil, err
	}

	// create master key and derive first key for keyring
	return algo.Derive()(mnemonic, "", hdPath)
}

// PrivKeyGen is the default PrivKeyGen function in the keybase.
// For now, it only supports Secp256k1
func PrivKeyGen(bz []byte) (types.PrivKey, error) {
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyring.SigningAlgoList{hd.Secp256k1})
	if err != nil {
		return nil, err
	}

	return algo.Generate()(bz), nil
}
