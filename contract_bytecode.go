package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RINKEBY_INFURA_URL"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	// This function CodeAt returns the contract code of the given account
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatalf("try get CodeAt error: %s", err.Error())
	}

	fmt.Println("Contract: ", hex.EncodeToString(bytecode))
}
