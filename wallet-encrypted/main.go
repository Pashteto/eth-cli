package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("../walletStored", keystore.StandardScryptN, keystore.StandardScryptP)
	pass := "password123"
	// a, err := key.NewAccount(pass)
	// if err != nil {
	// 	log.Fatalf("cant create new account:%s", err)
	// }
	// fmt.Println("address", a.Address)
	b, err := ioutil.ReadFile("../walletStored/UTC--2022-07-13T10-40-51.695396600Z--a238e06736f441acb9eaf1b651457c46f9758574")
	if err != nil {
		log.Fatalf("cant read file:%s", err)
	}
	key, err := keystore.DecryptKey(b, pass)
	if err != nil {
		log.Fatalf("%v", err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("priv key ", hexutil.Encode(pData))
	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("pub key ", hexutil.Encode(pData))
	fmt.Println("addr ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
