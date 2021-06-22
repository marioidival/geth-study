package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// create a private key with goeth-crypto
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("generate key error: %s", err.Error())
	}

	// convert privatekey into bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// convert into hexadecimal
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// makes a public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// create a public address based on publick key ECDSA
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

	// create a public address with other approach
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
