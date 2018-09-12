package relaynetwork

type Message struct {
	Kind    string   `json:"kind"`
	To      string   `json:"to"`
	EventId [32]byte `json:"eventId"`
	Data    []byte   `json:"data"`
}
