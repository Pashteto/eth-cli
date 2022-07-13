package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("cant create priv key:%s", err)
	}
	privateData := crypto.FromECDSA(privKey)
	fmt.Println("privateData", privateData)
	fmt.Println("privKey", privKey)
	encodedPrivateData := hexutil.Encode(privateData)
	fmt.Println("encodedPrivateData", encodedPrivateData)

	pubData := crypto.FromECDSAPub(&privKey.PublicKey)
	fmt.Println("pubData", pubData)

	encodedPubData := hexutil.Encode(pubData)
	fmt.Println("encodedPubData", encodedPubData)

	oi := crypto.PubkeyToAddress(privKey.PublicKey).Hex()
	fmt.Println("Address", oi)

}
