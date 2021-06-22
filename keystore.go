package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// CreateKeystore create a file containing an encrypted wallet and private key
func CreateKeystore() {
	ks := keystore.NewKeyStore("/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "thenewkeystore"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalf("create account error %s", err.Error())
	}

	fmt.Println("account created: ", account.Address.Hex())
}

// ImportKeystore load a file containing an encrypted account.
func ImportKeystore(filepath, password string) {
	ks := keystore.NewKeyStore("/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("import key error: %s", err.Error())
	}

	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatalf("import key error: %s", err.Error())
	}

	fmt.Println("account: ", account.Address.Hex())

	if err := os.Remove(filepath); err != nil {
		log.Fatalf("remove file error: %s", err.Error())
	}
}

func main() {
	CreateKeystore()
	ImportKeystore("/tmp/some-file", "thnewkeystore")
}
