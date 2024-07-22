package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"

	store "go-ganache/store"
)

// hardhat
var contractAddress = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"

// sepolia
// var contractAddress = "0xe0B3a0c579671E6593270E566C8eA6dFB5028cbF"

func load(connection Connection) *store.Store {
	fmt.Println("Loading contract")
	address := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(address, connection.ethClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	return instance
}

func write(connection Connection) {
	fmt.Println("Writing into contract")

	auth := getAuth(connection)
	instance := load(connection)

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

}

func getVersion(connection Connection) {
	version, err := load(connection).Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func getItems(connection Connection) {
	instance := load(connection)
	key := [32]byte{}
	copy(key[:], []byte("foo"))
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(string(result[:]))
}
