package main

import (
	"context"
	"encoding/json"
	"log"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	//relaynetwork "github.com/planet-ethereum/relay-network"
	ethbase "github.com/planet-ethereum/relay-network/ethbase"
	ipfs "github.com/planet-ethereum/relay-network/ipfs"
)

func main() {
	// Connect to IPFS
	ps := ipfs.NewIPFSPub("localhost:5001", "local-relay-network")

	// Connect to Ethereum client
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	emitterAddress := common.HexToAddress("0xfc8e2b378a37ea51b90049df56fad5f44cd35daf")
	registryAddress := common.HexToAddress("0x8c42e8e89c4ee591d69fdc95803388d80290c46c")
	//subscriberAddress := common.HexToAddress("0x065e983f32f905afbb26e7dbbcd21141afce3145")

	instance, err := ethbase.NewEthbase(client, registryAddress)
	if err != nil {
		log.Fatalf("failed to create ethbase instance: %v\n", err)
	}

	if err = instance.Initialize(); err != nil {
		log.Fatalf("failed to initialize ethbase: %v\n", err)
	}

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
			m, ok := instance.LogToMessage(vLog)
			if !ok {
				log.Printf("invalid log: %v\n", vLog)
				continue
			}

			marshalized, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("failed to marshalize message: %v\n", err)
			}

			if err = ps.Publish(string(marshalized)); err != nil {
				log.Fatalf("failed to publish to pubsub: %v\n", err)
			}
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
