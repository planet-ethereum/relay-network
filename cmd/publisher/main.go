package main

import (
	"log"

	ipfs "github.com/planet-ethereum/relay-network/ipfs"
)

func main() {
	ps := ipfs.NewIPFSPub("localhost:5001", "local-relay-network")
	if err := ps.Publish("hi"); err != nil {
		log.Fatalf("failed to publish to pubsub: %v", err)
	}
}
