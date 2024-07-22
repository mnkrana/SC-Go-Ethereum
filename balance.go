package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func getBalance(connection Connection, address common.Address) big.Float {
	balance, err := connection.ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error fetching balance %v", err)
	}
	ethValue := convertToETH(balance)
	fmt.Println(address, "has", ethValue, "ETH")
	return *ethValue
}

func convertToETH(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}
