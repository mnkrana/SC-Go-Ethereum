package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func getHeader() *big.Int {
// 	header, err := myclient.HeaderByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Block number", header.Number.String())
// 	getFullBlock(myclient, header.Number)
// 	return header.Number
// }

// func getFullBlock(client *ethclient.Client, number *big.Int) {
// 	block, err := client.BlockByNumber(context.Background(), number)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Block Number", block.Number().Uint64())
// 	fmt.Println("Block timestamp", block.Time())
// 	fmt.Println("Block difficulty", block.Difficulty())
// 	fmt.Println("Block hex", block.Hash().Hex())
// 	fmt.Println("Transactions", len(block.Transactions()))

// 	getTransactionCount(client, block.Hash())
// }

// func getTransactionCount(client *ethclient.Client, blockHash common.Hash) {
// 	count, err := client.TransactionCount(context.Background(), blockHash)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Transactoin count", count)
// }
