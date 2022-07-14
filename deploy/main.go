package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	todo "github.com/Pashteto/eth-cli/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var rinkeby = "https://rinkeby.infura.io/v3/a88b378452d94764af81d2ac741cefa7"

func main() {
	b, err := ioutil.ReadFile("../walletStored/UTC--2022-07-13T14-27-10.756206300Z--376ede682e30406c20f21dd238de6ef153bf84c3")
	if err != nil {
		log.Fatal(err)
	}

	pass := "qwrferfeqrqfvqfvqeffvqeefv"

	key, err := keystore.DecryptKey(b, pass)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(rinkeby)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}
