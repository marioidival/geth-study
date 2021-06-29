package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/marioidival/geth-study/contracts"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RINKEBY_INFURA_WS"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("try get logs error: %s", err.Error())
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatalf("try parse ABI interface error: %s", err.Error())
	}

	for _, vLog := range logs {
		fmt.Println("Block hash: ", vLog.BlockHash.Hex())
		fmt.Println("Block number: ", vLog.BlockNumber)
		fmt.Println("Tx Hash: ", vLog.TxHash.Hex())

		event, err := contractAbi.Unpack("ItemSet", vLog.Data)
		if err != nil {
			log.Fatalf("try unpack error: %s", err.Error())
		}

		fmt.Printf("Event: %q\n", event)

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0])
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
}
