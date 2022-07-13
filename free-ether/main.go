package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var mainnet = "https://mainnet.infura.io/v3/a88b378452d94764af81d2ac741cefa7"
var rinkeby = "https://rinkeby.infura.io/v3/a88b378452d94764af81d2ac741cefa7"

func main() {
	// key := keystore.NewKeyStore("../walletStored", keystore.StandardScryptN, keystore.StandardScryptP)
	// pass := "qwrferfeqrqfvqfvqeffvqeefv"
	// a1, err := key.NewAccount(pass)
	// if err != nil {
	// 	log.Fatalf("cant create new account:%s", err)
	// }
	// fmt.Println("address", a1.Address)

	// a2, err := key.NewAccount(pass)
	// if err != nil {
	// 	log.Fatalf("cant create new account:%s", err)
	// }
	// fmt.Println("address", a2.Address)

	a1s := common.HexToAddress("0x376Ede682e30406C20F21Dd238de6Ef153bF84c3")
	a2s := common.HexToAddress("0x6e15eDB390C1772178C5cf3beC2a9f729B288fdB")

	cli, err := ethclient.DialContext(context.Background(), rinkeby)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer cli.Close()
	b1, err := cli.BalanceAt(context.Background(), a1s, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	b2, err := cli.BalanceAt(context.Background(), a2s, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("bal1", b1, "\nbal2", b2)
}
