package main

import (
	"context"
	"log"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	// Connect to IPFS
	sh := ipfs.NewShell("localhost:5001")

	// Connect to Ethereum client
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	emitterAddress := common.HexToAddress("0xfc8e2b378a37ea51b90049df56fad5f44cd35daf")
	// registryAddress := common.HexToAddress("0x8c42e8e89c4ee591d69fdc95803388d80290c46c")

	// Test client
	balance, err := client.BalanceAt(context.Background(), emitterAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(balance)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			emitterAddress,
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("failed to subscribe to logs: %v\n", err)
	}

	log.Println("Subscribed to logs")

	for {
		select {
		case err := <-sub.Err():
			sub.Unsubscribe()
			log.Fatal(err)
		case vLog := <-logs:
			log.Printf("incoming log: %s\n", vLog.Address.Hex())
			log.Printf("data: %v\n", vLog.Data)
			err = sh.PubSubPublish("local-relay-network", vLog.Address.Hex())
		}
	}
}

func getPastLogs(client *ethclient.Client, emitterAddress common.Address) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			emitterAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	return logs, nil
}
