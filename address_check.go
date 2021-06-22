package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func checkIfAddressIsValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.Match([]byte(address))
}

func getCodeAt(client *ethclient.Client, account string) ([]byte, error) {
	address := common.HexToAddress(account)
	return client.CodeAt(context.Background(), address, nil)
}

func checkIfIsContractOrEthAddress(client *ethclient.Client, address string) {
	byteCode, err := getCodeAt(client, address)
	if err != nil {
		log.Fatalf("get CodeAt error: %s", err.Error())
	}
	fmt.Printf("is contract: %v\n", len(byteCode) > 0)
}

func main() {
	fmt.Println("is valid: ", checkIfAddressIsValid("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Println("is valid: ", checkIfAddressIsValid("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("ethclient error: %s", err.Error())
	}

	// both is false because ethClient is a ganache-cli :)
	checkIfIsContractOrEthAddress(client, "0xe41d2489571d322189246dafa5ebde1f4699f498")
	checkIfIsContractOrEthAddress(client, "0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
}
