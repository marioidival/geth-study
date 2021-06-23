package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RINKEBY_INFURA_URL"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatalf("try get private key error: %s", err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("try get peding nonce error: %s", err.Error())
	}

	// in wei (1 eth)
	value := big.NewInt(1000000000000000000)
	// in units
	gasLimit := uint64(21000)
	// get avarage gas price: 1000000000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("try get SuggestGasPrice error: %s", err.Error())
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("try get networkID error: %s", err.Error())
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("try signed transaction error: %s", err.Error())
	}

	// Insufficient funds for gas * price + value
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("try send transaction error: %s", err.Error())
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
