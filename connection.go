package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Connection struct {
	rpcUrl        string
	ethClient     *ethclient.Client
	chainId       *big.Int
	rawPrivateKey string
	privateKey    *ecdsa.PrivateKey
	address       common.Address
}

func (c *Connection) Print() {
	fmt.Println("RPC - ", c.rpcUrl)
	fmt.Println("Client - ", c.ethClient)
	fmt.Println("ChainId - ", c.chainId)
	fmt.Println("Raw PK - ", c.rawPrivateKey)
	fmt.Println("PK - ", c.privateKey)
	fmt.Println("Address - ", c.address)
}
