package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	store "go-ganache/store"
)

func deploy(connection Connection) {
	fmt.Println("Deploying....")
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

	input := "2.0"
	address, tx, instance, err := store.DeployStore(auth, connection.ethClient, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address ", address.Hex())
	fmt.Println("TX hash", tx.Hash().Hex())

	_ = instance
}
