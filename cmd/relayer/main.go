package main

import (
	"log"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ipfs "github.com/ipfs/go-ipfs-api"

	safe "github.com/planet-ethereum/relay-network/safe"
)

func main() {
	// Connect to IPFS
	sh := ipfs.NewShell("localhost:5001")
	sub, err := sh.PubSubSubscribe("local-relay-network")
	if err != nil {
		log.Fatalf("failed to subscribe to pubsub: %v", err)
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("failed to connect to ethclient: %v", err)
	}

	// Run pubsub event loop
	ch := make(chan ipfs.PubSubRecord)
	go listen(sub, ch)

	safeAddrHex := "0x254dffcd3277c0b1660f6d42efbb754edababc2b"
	safeAddr := common.HexToAddress(safeAddrHex)
	instance, err := safe.NewSafe(safeAddr, client)
	if err != nil {
		log.Fatalf("failed to create safe service: %v\n", err)
	}

	for {
		rec := <-ch
		log.Println(rec)

		nonce, err := instance.Nonce(nil)
		if err != nil {
			log.Fatalf("failed to get safe nonce: %v\n", err)
		}

		log.Printf("Safe nonce is: %v\n", nonce)
	}
}

func listen(sub *ipfs.PubSubSubscription, ch chan ipfs.PubSubRecord) {
	log.Println("Starting to listen to pubsub topic")

	for {
		rec, err := sub.Next()
		if err != nil {
			log.Fatalf("failed to wait for msg: %v", err)
		}

		ch <- rec
	}
}
