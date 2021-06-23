package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("MAINNET_INFURA_URL"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	// Get last block number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("try get block number error: %s", err.Error())
	}
	fmt.Println(header.Number.String()) // 12690616

	// Get block using blockNumber
	blockNumber := big.NewInt(12690616)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("try get block number returned error: %s", err.Error())
	}

	fmt.Println("block number: ", block.Number().Uint64())
	fmt.Println("block time", block.Time())
	fmt.Println("block difficult", block.Difficulty().Uint64())
	fmt.Println("block hash", block.Hash().Hex())
	fmt.Println("quantity transaction", len(block.Transactions()))

	// Get TransactionCount using block hash
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatalf("try get transaction count by block hash: %s", err.Error())
	}
	fmt.Println("count by transaction count", count)
}
