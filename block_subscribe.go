package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("ROPSTEN_INFURA_WS"))
	if err != nil {
		log.Fatalf("ethclient error: %s", err.Error())
	}

	// create a new channel to receive the latest block headers
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatalf("error to create subscribeNewHead: %s", err.Error())
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("try push new block headers", err.Error())
		case header := <-headers:
			fmt.Println("header hash: ", header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatalf("try get block by hash: %s", err.Error())
			}

			fmt.Println("block hash: ", block.Hash().Hex())
			fmt.Println("block number: ", block.Number().Uint64())
			fmt.Println("block time: ", block.Time())
			fmt.Println("block nonce :", block.Nonce())
			fmt.Println("block transactions: ", len(block.Transactions()))
		}
	}
}
