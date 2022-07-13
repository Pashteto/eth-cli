package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ethUrl = "https://mainnet.infura.io/v3/a88b378452d94764af81d2ac741cefa7"

func main() {
	client, err := ethclient.DialContext(context.Background(), ethUrl)
	if err != nil {
		log.Fatalf("cant create eth client:%s", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("cant get latest block from client:%s", err)
	}
	fmt.Println(block.Number())
}
