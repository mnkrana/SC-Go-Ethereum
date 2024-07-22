package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(rpcUrl string) *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	return client
}

func getChainId(connection Connection) *big.Int {
	chainId, err := connection.ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return chainId
}
