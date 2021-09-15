package eosrpc

import "github.com/TheBoringDude/go-urljoin"

const TRACE_API = "trace_api"

// GetBlock returns a block trace object containing retired actions and related metadata.
func (t *TraceAPI) GetBlock(props GetBlockInfoProps) (TraceGetBlockResponse, error) {
	var r = TraceGetBlockResponse{}
	err := request(t.Client, urljoin.UrlJoin(t.ApiUrl, "get_block"), props, &r)

	return r, err
}
