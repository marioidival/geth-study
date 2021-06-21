package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
