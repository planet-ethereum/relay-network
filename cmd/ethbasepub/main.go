package main

import (
	"context"
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
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	emitterAddress := common.HexToAddress("0x8d4c2eea61d3694534f6b8d6433d1d7edfe82b34")
	registryAddress := common.HexToAddress("0xababd383f1debf1434193bb4a44e7476f3a0bd2c")
	//subscriberAddress := common.HexToAddress("0x95adff3bb970f6548254331c8f481e857b48a5d9")

	instance, err := ethbase.NewEthbase(client, registryAddress)
	if err != nil {
		log.Fatalf("failed to create ethbase instance: %v\n", err)
	}

	if err = instance.Initialize(); err != nil {
		log.Fatalf("failed to initialize ethbase: %v\n", err)
	}

	errCh := make(chan error)
	msgCh := make(chan []byte)
	if err := instance.Listen(context.Background(), msgCh, errCh, []common.Address{emitterAddress}); err != nil {
		log.Fatalf("failed to listen for incoming events")
	}

	for {
		select {
		case msg := <-msgCh:
			if err = ps.Publish(string(msg)); err != nil {
				log.Fatalf("failed to publish to pubsub: %v\n", err)
			}
		case err := <-errCh:
			log.Fatalf("err from listener: %v\n", err)
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
