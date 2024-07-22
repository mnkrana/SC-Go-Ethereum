package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := Connection{
		rpcUrl:        os.Getenv("GANACHE_RPC_URL"),
		rawPrivateKey: os.Getenv("PRIVATE_KEY"),
	}
	if os.Getenv("TESTNET") == "1" {
		fmt.Println("Ovveride with testnet config")
		connection.rpcUrl = os.Getenv("SEPOLIA_RPC_URL")
		connection.rawPrivateKey = os.Getenv("SEPOLIA_WALLET_PRIVATE_KEY")
	}

	connection.ethClient = getClient(connection.rpcUrl)
	connection.chainId = getChainId(connection)
	defer connection.ethClient.Close()

	connection.privateKey = getPrivateKeyFromString(connection.rawPrivateKey)
	connection.address = getAddressFromPrivateKey(connection.privateKey)
	// connection.Print()

	// if verifyAddress() {
	// verifyContract()
	// getBalance(myAddress)
	// getHeader()
	// transferFunds()
	// deploy(connection)
	// load(connection)
	// write(connection)
	getVersion(connection)
	getItems(connection)
	// }
}
