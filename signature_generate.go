package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatalf("try load private key: %s", err.Error())
	}

	data := []byte("hello boy")
	hash := crypto.Keccak256Hash(data)
	fmt.Printf("Hash: %s\n", hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatalf("try sign the data: %s", err.Error())
	}

	fmt.Println(hexutil.Encode(signature))
}
