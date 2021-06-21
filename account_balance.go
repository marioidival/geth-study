package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("connection error: %s ", err.Error())
	}

	// reading balance of an account
	account := common.HexToAddress("0xc203E29ecBEc4a1abAC5a1A5f0266235d1b7eF82")
	// third arg from BalanceAt is a block number
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("balanceAt error: %s ", err.Error())
	}
	fmt.Println(balance)

	blockNumber := big.NewInt(0)
	// using ganachi-cli, the block number to this account is 0
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatalf("balanceAt error with block number %s", err.Error())
	}
	fmt.Println(balanceAt)

	// do conversion to smallest possible unit, wei
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatalf("pendingBalance error %s", err.Error())
	}
	fmt.Println(pendingBalance)
}
