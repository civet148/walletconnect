package proto

type Sequence struct {
	Topic string `json:"topic"`
}

type AppMetadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Icons       []string `json:"icons"`
}

type SessionPeerMeta struct {
	PeerId   string      `json:"peerId"`
	PeerMeta AppMetadata `json:"peerMeta"`
	ChainId  int64       `json:"chainId"`
}
