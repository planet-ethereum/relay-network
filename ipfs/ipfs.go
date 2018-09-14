package ipfs

import (
	"log"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type IPFSPubSub struct {
	shell        *ipfs.Shell
	subscription *ipfs.PubSubSubscription
	topic        string
}

func NewIPFSPub(endpoint, topic string) *IPFSPubSub {
	ps := new(IPFSPubSub)
	ps.shell = ipfs.NewShell(endpoint)
	ps.topic = topic
	return ps
}

func NewIPFSSub(endpoint, topic string) (*IPFSPubSub, error) {
	ps := new(IPFSPubSub)

	shell := ipfs.NewShell(endpoint)
	ps.shell = shell
	ps.topic = topic

	sub, err := shell.PubSubSubscribe(topic)
	if err != nil {
		return nil, err
	}
	ps.subscription = sub

	return ps, nil
}

func (ps *IPFSPubSub) Publish(data string) error {
	return ps.shell.PubSubPublish(ps.topic, data)
}

func (ps *IPFSPubSub) Listen(ch chan []byte) {
	for {
		rec, err := ps.subscription.Next()
		if err != nil {
			log.Fatalf("failed to wait for msg: %v", err)
		}

		ch <- rec.Data()
	}
}
