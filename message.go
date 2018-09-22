package relaynetwork

type Message struct {
	Kind string `json:"kind"`
	To   string `json:"to"`
	Data []byte `json:"data"` // Protocol-specific message
}
