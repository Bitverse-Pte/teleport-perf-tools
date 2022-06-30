package main

import (
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/wallet"
)

func importWallet() (ethgo.Key, error) {
	key, err := wallet.NewJSONWalletFromFile(walletFile, password)
	handleErr(err)
	return key, nil
}
