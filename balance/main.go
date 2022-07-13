package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraUrl = "https://mainnet.infura.io/v3/a88b378452d94764af81d2ac741cefa7"
var ganacheUrl = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraUrl)
	if err != nil {
		log.Fatalf("cant create eth client:%s", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("cant get latest block from client:%s", err)
	}
	fmt.Println("latest block", block.Number())

	var addrWallet = "0x0dB82039be0c3d77908Ec54B73dA8b6f1D26B66D"
	addressHex := common.HexToAddress(addrWallet)
	balance, err := client.BalanceAt(context.Background(), addressHex, nil)
	if err != nil {
		log.Fatalf("cant get balance:%s", err)
	}
	fmt.Println("the balance: ", balance)
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println("the fBalance: ", fBalance)
	val := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("the val: ", val)
}
