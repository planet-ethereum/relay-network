package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	relaynetwork "github.com/planet-ethereum/relay-network"
	ethbase "github.com/planet-ethereum/relay-network/ethbase"
	ipfs "github.com/planet-ethereum/relay-network/ipfs"
	wallet "github.com/planet-ethereum/relay-network/wallet"
	//safe "github.com/planet-ethereum/relay-network/safe"
)

func main() {
	// Connect to IPFS
	ps, err := ipfs.NewIPFSSub("localhost:5001", "local-relay-network")
	if err != nil {
		log.Fatalf("failed to subscribe to pubsub: %v", err)
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	// Set up wallet
	wal, err := wallet.NewHDWallet(os.Getenv("MNEMONIC"), "m/44'/60'/0'/0/2")
	if err != nil {
		log.Fatalf("failed to create wallet: %v\n", err)
	}

	fromAddress, err := wal.Address()
	if err != nil {
		log.Fatalf("failed to get account address: %v\n", err)
	}

	log.Printf("Account: %s\n", fromAddress.Hex())

	// Set up ethbase
	registryAddress := common.HexToAddress("0xababd383f1debf1434193bb4a44e7476f3a0bd2c")
	ethbase_, err := ethbase.NewEthbase(client, registryAddress)
	if err != nil {
		log.Fatalf("failed to instantiate ethbase: %v\n", err)
	}

	if err = ethbase_.Initialize(); err != nil {
		log.Fatalf("failed to initialize ethbase: %v\n", err)
	}

	log.Println("Starting to listen to pubsub topic")

	// Run pubsub event loop
	ch := make(chan []byte)
	go ps.Listen(ch)

	for {
		raw := <-ch

		m := relaynetwork.Message{}
		err = json.Unmarshal(raw, &m)

		if m.Kind == "ethbase" {
			log.Printf("Submitting: Kind: %s\tTo: %s\tData: %v\n", m.Kind, m.To, m.Data)
			if err = ethbase_.SubmitTx(context.Background(), wal, m); err != nil {
				log.Printf("error: failed to submit tx: %v\n", err)
			}
		}
	}
}
