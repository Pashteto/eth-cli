package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// var mainnet = "https://mainnet.infura.io/v3/a88b378452d94764af81d2ac741cefa7"
var rinkeby = "https://rinkeby.infura.io/v3/a88b378452d94764af81d2ac741cefa7"

func main() {
	addr1 := common.HexToAddress("0x376Ede682e30406C20F21Dd238de6Ef153bF84c3")
	addr2 := common.HexToAddress("0x6e15eDB390C1772178C5cf3beC2a9f729B288fdB")

	ctx := context.Background()
	cli, err := ethclient.DialContext(ctx, rinkeby)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer cli.Close()

	b1, err := cli.BalanceAt(ctx, addr1, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	b2, err := cli.BalanceAt(ctx, addr2, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println("bal1", b1, "\nbal2", b2)

	nonce, err := cli.PendingNonceAt(ctx, addr1)
	if err != nil {
		log.Fatalf("%s", err)
	}

	amount := big.NewInt(1000000000000000000 / 10) //1 eth/10
	gasPrice, err := cli.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("%s", err)
	}

	tx := types.NewTransaction(nonce, addr2, amount, 21000, gasPrice, nil)
	chainID, err := cli.NetworkID(ctx)
	if err != nil {
		log.Fatalf("%s", err)
	}

	b, err := ioutil.ReadFile("../walletStored/UTC--2022-07-13T14-27-10.756206300Z--376ede682e30406c20f21dd238de6ef153bf84c3")
	if err != nil {
		log.Fatalf("cant read file:%s", err)
	}
	pass := "qwrferfeqrqfvqfvqeffvqeefv"

	key, err := keystore.DecryptKey(b, pass)
	if err != nil {
		log.Fatalf("%v", err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("priv key ", hexutil.Encode(pData))

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = cli.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
