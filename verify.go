package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"regexp"
// )

// func verifyAddress() bool {

// 	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
// 	valid := re.MatchString(myAddress.String())
// 	if valid {
// 		fmt.Println(myAddress, "is valid address")
// 	} else {
// 		fmt.Println(myAddress, "is not a valid address")
// 	}
// 	return valid
// }

// func verifyContract() bool {
// 	bytecode, err := myclient.CodeAt(context.Background(), myAddress, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isContract := len(bytecode) > 0

// 	if isContract {
// 		fmt.Println(myAddress, "is a contract")
// 	} else {
// 		fmt.Println(myAddress, "is not a contract")
// 	}

// 	return isContract
// }
