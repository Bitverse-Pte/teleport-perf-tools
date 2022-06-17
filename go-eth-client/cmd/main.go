package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/testutil"
	"github.com/umbracle/ethgo/wallet"
	"math/big"
	"time"
)

var (
	walletFile = "wallet.json"
	password   = "123"
)

// call a contract
func main() {
	//
	c, err := jsonrpc.NewClient("http://172.17.0.1:8545")
	handleErr(err)

	key, _ := importWallet()
	fmt.Println(key.Address())
	found, err := c.Eth().GetBalance(key.Address(), ethgo.Latest)
	handleErr(err)
	fmt.Println(found.String())

	balance, err := c.Eth().GetBalance(testutil.DummyAddr, ethgo.Latest)
	fmt.Printf("before balance is: %d\n", balance)

	nonce, _ := c.Eth().GetNonce(key.Address(), ethgo.Latest)
	fmt.Println(nonce)

	hash, err := sendBatchTx(key, c, nonce, 20)
	_, err = WaitForReceipt(c, hash)

	balance, err = c.Eth().GetBalance(testutil.DummyAddr, ethgo.Latest)
	fmt.Printf("after balance is: %d\n", balance)
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func sendBatchTx(key ethgo.Key, c *jsonrpc.Client, nonce uint64, number uint64) (ethgo.Hash, error) {
	var i uint64
	var hash ethgo.Hash
	for i = nonce; i < nonce+number; i++ {
		txn := &ethgo.Transaction{
			From:     key.Address(),
			Nonce:    i,
			GasPrice: testutil.DefaultGasPrice,
			Gas:      testutil.DefaultGasLimit,
			To:       &testutil.DummyAddr,
			Value:    big.NewInt(10),
		}

		chainID, err := c.Eth().ChainID()
		handleErr(err)
		signer := wallet.NewEIP155Signer(chainID.Uint64())
		txn, err = signer.SignTx(txn, key)
		data, err := txn.MarshalRLPTo(nil)

		hash, err = c.Eth().SendRawTransaction(data)
		fmt.Printf("send transaction hash:%s\n", hash)
	}
	return hash, nil
}

func WaitForReceipt(c *jsonrpc.Client, hash ethgo.Hash) (*ethgo.Receipt, error) {
	var receipt *ethgo.Receipt
	var count uint64
	for {
		err := c.Call("eth_getTransactionReceipt", &receipt, hash)
		if err != nil {
			if err.Error() != "not found" {
				return nil, err
			}
		}
		if receipt != nil {
			break
		}
		if count > 100 {
			return nil, fmt.Errorf("timeout")
		}
		time.Sleep(500 * time.Millisecond)
		count++
	}
	return receipt, nil
}
