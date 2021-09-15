package eosrpc

import (
	"net/http"

	"github.com/TheBoringDude/go-urljoin"
)

type EOS_RPC struct {
	ApiUrl string
	Client *http.Client
}

type ChainAPI EOS_RPC

type TraceAPI EOS_RPC

type ProducerAPI EOS_RPC

type NetAPI EOS_RPC

type DBSizeAPI EOS_RPC

// New creates a new instance of EOS_RPC
func New(apiUrl string) *EOS_RPC {
	return &EOS_RPC{
		ApiUrl: urljoin.UrlJoin(apiUrl, "v1"),
		Client: &http.Client{},
	}
}

// https://developers.eos.io/manuals/eos/latest/nodeos/plugins/chain_api_plugin/api-reference/index
func (e *EOS_RPC) NewChainAPI() *ChainAPI {
	return &ChainAPI{urljoin.UrlJoin(e.ApiUrl, CHAIN_API), e.Client}
}

// https://developers.eos.io/manuals/eos/latest/nodeos/plugins/trace_api_plugin/api-reference/index
func (e *EOS_RPC) NewTraceAPI() *TraceAPI {
	return &TraceAPI{urljoin.UrlJoin(e.ApiUrl, TRACE_API), e.Client}
}

// https://developers.eos.io/manuals/eos/latest/nodeos/plugins/producer_api_plugin/api-reference/index#operation/create_snapshot
func (e *EOS_RPC) NewProducerAPI() *ProducerAPI {
	return &ProducerAPI{urljoin.UrlJoin(e.ApiUrl, PRODUCER_API), e.Client}
}

// https://developers.eos.io/manuals/eos/latest/nodeos/plugins/net_api_plugin/api-reference/index
func (e *EOS_RPC) NewNetAPI() *NetAPI {
	return &NetAPI{urljoin.UrlJoin(e.ApiUrl, NET_API), e.Client}
}

// https://developers.eos.io/manuals/eos/latest/nodeos/plugins/db_size_api_plugin/api-reference/index
func (e *EOS_RPC) NewDBSizeAPI() *DBSizeAPI {
	return &DBSizeAPI{urljoin.UrlJoin(e.ApiUrl, DBSIZE_API), e.Client}
}
