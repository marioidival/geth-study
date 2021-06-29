package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RINKEBY_INFURA_WS"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("try get logs error: %s", err.Error())
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("error when try get logs: %s", err.Error())
		case vLog := <-logs:
			fmt.Println("log: ", vLog)
		}
	}
}
