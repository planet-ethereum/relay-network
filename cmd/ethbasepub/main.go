package main

import (
	"context"
	"encoding/json"
	"log"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/ipfs/go-ipfs-api"

	relaynetwork "github.com/planet-ethereum/relay-network"
	ethbase "github.com/planet-ethereum/relay-network/ethbase"
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
	registryAddress := common.HexToAddress("0x8c42e8e89c4ee591d69fdc95803388d80290c46c")
	subscriberAddress := common.HexToAddress("0x065e983f32f905afbb26e7dbbcd21141afce3145")

	/*registry, err := ethbase.NewRegistry(registryAddress, client)
	if err != nil {
		log.Fatalf("failed to create registry instance: %v\n", err)
	}*/

	subscriptions, err := ethbase.GetSubscriptions(client, registryAddress)
	for _, s := range subscriptions {
		eventId := s.EventId[:]
		log.Printf("Sub: %v\n", common.Bytes2Hex(eventId))
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
			m := relaynetwork.Message{
				Kind: "etherbase",
				To:   subscriberAddress.Hex(),
				Data: vLog.Data,
			}

			marshalized, err := json.Marshal(m)
			if err != nil {
				log.Fatalf("failed to marshalize message: %v\n", err)
			}

			if err = sh.PubSubPublish("local-relay-network", string(marshalized)); err != nil {
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
