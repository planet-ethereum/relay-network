package main

import (
	"context"
	"log"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := ipfs.NewShell("localhost:5001")
	sub, err := sh.PubSubSubscribe("relay-network")
	if err != nil {
		log.Fatalf("failed to subscribe to pubsub: %v", err)
	}

	rec, err := sub.Next()
	if err != nil {
		log.Fatalf("failed to wait for msg: %v", err)
	}

	log.Printf("record: %v\n", rec)

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	account := common.HexToAddress("0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1")
	log.Printf("Addr: %s", account.Hex())
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("failed to get account balance: %v", err)
	}

	log.Println(balance)
}
