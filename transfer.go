package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func getAuth(connection Connection) *bind.TransactOpts {
	nonce, err := connection.ethClient.PendingNonceAt(context.Background(), connection.address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce -", nonce)

	gasPrice, err := connection.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gas Price -", gasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(connection.privateKey, connection.chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fmt.Println(auth.Nonce, auth.Value, auth.GasLimit, auth.GasPrice)
	return auth
}

// func transferFunds() {
// 	fmt.Println("Transfering....")
// 	targetAddress := common.HexToAddress(targetPublicKey)
// 	getBalance(targetAddress)

// 	transfer(targetAddress)

// 	fmt.Println("Transfer done!")

// 	getBalance(myAddress)
// 	getBalance(targetAddress)
// }

// todo - test with getAuth
func transfer(connection Connection, target common.Address) {
	nonce, err := connection.ethClient.PendingNonceAt(context.Background(), connection.address)
	if err != nil {
		log.Fatal(err)
	}

	eth := big.NewInt(1000000000000000000)
	value := big.NewInt(10)
	transferAmount := big.NewInt(0)
	transferAmount.Mul(eth, value)
	fmt.Println("Transfer Amount", convertToETH(transferAmount))

	gasLimit := uint64(21000)
	gasPrice, err := connection.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Gas price", gasPrice)

	var data []byte
	tx := types.NewTransaction(nonce, target, transferAmount, gasLimit, gasPrice, data)
	fmt.Println("Transaction is", tx)

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, connection.privateKey)
	if err != nil {
		log.Fatalln("Error in signing tx", err)
	}

	fmt.Println("Sending tx")
	err = connection.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
