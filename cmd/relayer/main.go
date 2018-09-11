package main

import (
	"log"

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
}
