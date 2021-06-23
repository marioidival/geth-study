package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("MAINNET_INFURA_URL"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	blockNumber := big.NewInt(12690616)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("try get blockByNumber error: %s", err.Error())
	}

	// Get all transactions from block 12690616
	for _, tx := range block.Transactions() {
		fmt.Println("Hash: ", tx.Hash().Hex())
		fmt.Println("Value :", tx.Value().String())
		fmt.Println("Gas: ", tx.Gas())
		fmt.Println("Gas price: ", tx.GasPrice().Uint64())
		fmt.Println("Nonce: ", tx.Nonce())
		fmt.Println("Data: ", tx.Data())
		fmt.Println("To: ", tx.To().Hex())

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatalf("try get networkID error: %s", err.Error())
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
			fmt.Println(msg.From().Hex())
			fmt.Println(msg.From().String())
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatalf("try get Transaction Receipt error: %s", err.Error())
		}
		fmt.Println("receipt status: ", receipt.Status)
	}

	// Get transaction from blockhash
	blockHash := common.HexToHash("0xb8c3807fbcc0ae7c4aabeb786d6f03c4b9ee28065e1ce83912709a7e6aabeeb6")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatalf("try get transaction count error: %s", err.Error())
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatalf("try get transaction in block error: %s", err.Error())
		}
		fmt.Println("Tx Hash: ", tx.Hash().Hex())
	}

	// Query single transaction given the transaction hash
	txHash := common.HexToHash("0x66bf65c63f7337bcdfc9b16539623aab6821fb93c955f7530cd7e26c2251341a")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("try get transaction by hash error: %s", err.Error())
	}

	fmt.Println("Tx Hash: ", tx.Hash().Hex())
	fmt.Println("Is Pending? ", isPending)
}
