package eosrpc

type NetConnectionsResponse []NetConnectionsResponseElement

type NetConnectionsResponseElement struct {
	Peer          string        `json:"peer"`
	Connecting    bool          `json:"connecting"`
	Syncing       bool          `json:"syncing"`
	LastHandshake LastHandshake `json:"last_handshake"`
}

type LastHandshake struct {
	NetworkVersion           int64  `json:"network_version"`
	ChainID                  string `json:"chain_id"`
	NodeID                   string `json:"node_id"`
	Key                      string `json:"key"`
	Time                     string `json:"time"`
	Token                    string `json:"token"`
	Sig                      string `json:"sig"`
	P2PAddress               string `json:"p2p_address"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID  string `json:"last_irreversible_block_id"`
	HeadNum                  int64  `json:"head_num"`
	HeadID                   string `json:"head_id"`
	OS                       string `json:"os"`
	Agent                    string `json:"agent"`
	Generation               int64  `json:"generation"`
}
