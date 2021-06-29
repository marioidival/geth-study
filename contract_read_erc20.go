package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	token "github.com/marioidival/geth-study/contracts_erc20"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("MAINNET_INFURA_URL"))
	if err != nil {
		log.Fatalf("ethclient connection error: %s", err.Error())
	}

	// Golem (GNT) address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatalf("try get new token error: %s", err.Error())
	}

	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatalf("try get balance from smart contract error: %s", err.Error())
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("try get name from smart contract error: %s", err.Error())
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("try get symbol from smart contract error: %s", err.Error())
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("try get decimals from smart contract error: %s", err.Error())
	}

	fmt.Println("Name: ", name)
	fmt.Println("Symbol: ", symbol)
	fmt.Println("decimals: ", decimals)
	fmt.Println("WEI (unformatted): ", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("Balance: %f\n", value)

}
