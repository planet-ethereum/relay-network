package main

import (
	"log"

	ipfs "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := ipfs.NewShell("localhost:5001")
	if err := sh.PubSubPublish("local-relay-network", "hi"); err != nil {
		log.Fatalf("failed to publish to pubsub: %v", err)
	}
}
