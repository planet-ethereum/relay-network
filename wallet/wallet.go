package wallet

import (
	accounts "github.com/ethereum/go-ethereum/accounts"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type HDWallet struct {
	wallet  *hdwallet.Wallet
	account accounts.Account
}

func NewHDWallet(mnemonic string, derivationPath string) (*HDWallet, error) {
	hdw := new(HDWallet)

	w, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath(derivationPath)
	a, err := w.Derive(path, true)
	if err != nil {
		return nil, err
	}

	hdw.wallet = w
	hdw.account = a

	return hdw, nil
}

func (w *HDWallet) Address() (common.Address, error) {
	return w.wallet.Address(w.account)
}

func (w *HDWallet) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return w.wallet.SignTx(w.account, tx, nil)
}

func (w *HDWallet) NewKeyedTransactor() (*bind.TransactOpts, error) {
	priv, err := w.wallet.PrivateKey(w.account)
	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactor(priv), nil
}
