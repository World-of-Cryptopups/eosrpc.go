package eosrpc

import "github.com/TheBoringDude/go-urljoin"

const NET_API = "net"

// Connections returns an array of all peer connection statuses.
func (n *NetAPI) Connections() (NetConnectionsResponse, error) {
	var r = NetConnectionsResponse{}
	err := request(n.Client, urljoin.UrlJoin(n.ApiUrl, "connections"), nil, &r)

	return r, err
}

type NetConnectProps struct {
	Endpoint string `json:"endpoint"`
}

// Connect initiates a connection to a specified peer.
func (n *NetAPI) Connect(props NetConnectProps) (string, error) {
	var r = ""
	err := request(n.Client, urljoin.UrlJoin(n.ApiUrl, "connect"), props, &r)

	return r, err
}

// Disconnect initiates a diconnection to a specified peer.
func (n *NetAPI) Disconnect(props NetConnectProps) (string, error) {
	var r = ""
	err := request(n.Client, urljoin.UrlJoin(n.ApiUrl, "disconnect"), props, &r)

	return r, err
}

// Status retrieves the connection status for a specified peer.
func (n *NetAPI) Status(props NetConnectProps) (NetConnectionsResponseElement, error) {
	var r = NetConnectionsResponseElement{}
	err := request(n.Client, urljoin.UrlJoin(n.ApiUrl, "disconnect"), props, &r)

	return r, err
}
