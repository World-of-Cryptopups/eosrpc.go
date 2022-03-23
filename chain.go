package eosrpc

import (
	"fmt"

	"github.com/TheBoringDude/go-urljoin"
)

const CHAIN_API = "chain"

type AccountProps struct {
	AccountName string `json:"account_name"`
}

// GetAccount returns an object containing various details about a specific account on the blockchain.
func (c *ChainAPI) GetAccount(props AccountProps) (GetAccountResponse, error) {
	var r = GetAccountResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_account"), props, &r)

	return r, err
}

type GetBlockProps struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

// GetBlock returns an object containing various details about a specific block on the blockchain.
func (c *ChainAPI) GetBlock(props GetBlockProps) (GetBlockResponse, error) {
	var r = GetBlockResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_block"), props, &r)

	return r, err
}

type GetBlockInfoProps struct {
	BlockNum int `json:"block_num"`
}

// GetBlockInfo is similar to get_block but returns a fixed-size smaller subset of the block data.
func (c *ChainAPI) GetBlockInfo(props GetBlockInfoProps) (GetBlockInfoResponse, error) {
	var r = GetBlockInfoResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_block_info"), props, &r)

	return r, err
}

// GetInfo returns an object containing various details about a specific block on the blockchain.
func (c *ChainAPI) GetInfo() (GetInfoResponse, error) {
	var r = GetInfoResponse{}
	fmt.Println(urljoin.UrlJoin(c.ApiUrl, "get_info"))

	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_info"), emptyMap, &r)

	return r, err
}

type TransactionProps struct {
	Signatures            []string `json:"signatures,omitempty"`
	Compression           bool     `json:"compression,omitempty"`
	PackedContextFreeData string   `json:"packed_context_free_data,omitempty"`
	PackedTrx             string   `json:"packed_trx,omitempty"`
}

// PushTransaction expects a transaction in JSON format and will attempt to apply it to the blockchain.
func (c *ChainAPI) PushTransaction(props TransactionProps) error {
	return request(c.Client, urljoin.UrlJoin(c.ApiUrl, "push_transaction"), props, nil)
}

// SendTransaction expects a transaction in JSON format and will attempt to apply it to the blockchain.
func (c *ChainAPI) SendTransaction(props TransactionProps) error {
	return request(c.Client, urljoin.UrlJoin(c.ApiUrl, "send_transaction"), props, nil)
}

type PushTransactionProps []PushTransactionObject

type PushTransactionObject struct {
	Expiration            string              `json:"expiration"`
	RefBlockNum           int64               `json:"ref_block_num"`
	RefBlockPrefix        int64               `json:"ref_block_prefix"`
	MaxNetUsageWords      string              `json:"max_net_usage_words"`
	MaxCPUUsageMS         string              `json:"max_cpu_usage_ms"`
	DelaySEC              int64               `json:"delay_sec"`
	ContextFreeActions    []TransactionAction `json:"context_free_actions"`
	Actions               []TransactionAction `json:"actions"`
	TransactionExtensions [][]int64           `json:"transaction_extensions,omitempty"`
}

type TransactionAction struct {
	Account       string                 `json:"account"`
	Name          string                 `json:"name"`
	Authorization []Authorization        `json:"authorization"`
	Data          map[string]interface{} `json:"data"`
	HexData       string                 `json:"hex_data"`
}

// PushTransactions expects a transaction in JSON format and will attempt to apply it to the blockchain.
func (c *ChainAPI) PushTransactions(props PushTransactionProps) error {
	return request(c.Client, urljoin.UrlJoin(c.ApiUrl, "push_transactions"), props, nil)
}

type GetBlockHeaderStateProps struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

// GetBlockHeaderState retrieves the block header state.
func (c *ChainAPI) GetBlockHeaderState(props GetBlockHeaderStateProps) (GetBlockHeaderStateResponse, error) {
	var r = GetBlockHeaderStateResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_block_header_state"), props, &r)

	return r, err
}

// GetABI retrieves the ABI for a contract based on its account name.
func (c *ChainAPI) GetABI(props AccountProps) (GetABIResponse, error) {
	var r = GetABIResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_abi"), props, &r)

	return r, err
}

type GetCurrencyBalanceProps struct {
	Code    string `json:"code"`
	Account string `json:"account"`
	Symbol  string `json:"symbol,omitempty"`
}

// GetCurrencyBalance returns the current balance.
func (c *ChainAPI) GetCurrencyBalance(props GetCurrencyBalanceProps) (GetCurrencyBalanceResponse, error) {
	var r = GetCurrencyBalanceResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_currency_balance"), props, &r)

	return r, err
}

type GetCurrencyStatsProps struct {
	Code   string `json:"code,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}

// GetCurrencyStats retrieves currency stats.
func (c *ChainAPI) GetCurrencyStats(props GetCurrencyStatsProps) (GetCurrencyStatsResponse, error) {
	var r interface{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_currency_stats"), props, &r)

	return r, err
}

type GetRequiredKeysProps struct {
	Transaction   Transaction `json:"transaction"`
	AvailableKeys []string    `json:"available_keys"`
}

// GetRequiredKeys returns the required keys needed to sign a transaction.
func (c *ChainAPI) GetRequiredKeys(props GetRequiredKeysProps) (GetRequiredKeysResponse, error) {
	var r = GetRequiredKeysResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_required_keys"), props, &r)

	return r, err
}

type GetProducersProps struct {
	Limit      string `json:"limit"`
	LowerBound string `json:"lower_bound"`
	JSON       bool   `json:"json,omitempty"`
}

// GetProducers retrieves producers list.
func (c *ChainAPI) GetProducers(props GetProducersProps) (GetProducersResponse, error) {
	var r = GetProducersResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_producers"), props, &r)

	return r, err
}

// GetRawCodeAndABI retrieves raw code and ABI for a contract based on account name.
func (c *ChainAPI) GetRawCodeAndABI(props AccountProps) (GetRawCodeAndABIResponse, error) {
	var r = GetRawCodeAndABIResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_raw_code_and_abi"), props, &r)

	return r, err
}

type GetScheduledTransactionProps struct {
	LowerBound string `json:"lower_bound,omitempty"`
	Limit      int64  `json:"limit"`
	JSON       bool   `json:"json"`
}

// GetScheduledTransaction retrieves the scheduled transaction.
func (c *ChainAPI) GetScheduledTransaction(props GetScheduledTransactionProps) (GetScheduledTransactionResponse, error) {
	var r = GetScheduledTransactionResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_scheduled_transaction"), props, &r)

	return r, err
}

type GetTableByScopeProps struct {
	Code       string `json:"code"`
	Table      string `json:"table,omitempty"`
	LowerBound string `json:"lower_bound,omitempty"`
	UpperBound string `json:"upper_bound,omitempty"`
	Limit      int32  `json:"limit,omitempty"`
	Reverse    bool   `json:"reverse,omitempty"`
	ShowPayer  bool   `json:"show_payer,omitempty"`
}

// GetTableByScope retrieves table scope.
func (c *ChainAPI) GetTableByScope(props GetTableByScopeProps) (GetTableByScopeResponse, error) {
	var r = GetTableByScopeResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_table_by_scope"), props, &r)

	return r, err
}

type GetTableRowsProps struct {
	Code          string `json:"code"`
	Table         string `json:"table"`
	Scope         string `json:"scope"`
	IndexPosition string `json:"index_position"`
	KeyType       string `json:"key_type"`
	EncodeType    string `json:"encode_type"`
	LowerBound    string `json:"lower_bound"`
	UpperBound    string `json:"upper_bound"`
	Limit         int32  `json:"limit"`
	Reverse       bool   `json:"reverse"`
	ShowPayer     bool   `json:"show_payer"`
}

// GetTableRows returns an object containing rows from the specified table.
func (c *ChainAPI) GetTableRows(props GetTableRowsProps) (GetTableRowsResponse, error) {
	var r = GetTableRowsResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_table_rows"), props, &r)

	return r, err
}

type GetKVTableRowsProps struct {
	Code       string      `json:"code"`
	Table      string      `json:"table"`
	IndexName  string      `json:"index_name"`
	EncodeType interface{} `json:"encode_type"`
	IndexValue string      `json:"index_value"`
	LowerBound string      `json:"lower_bound"`
	UpperBound string      `json:"upper_bound"`
	Limit      int64       `json:"limit"`
	Reverse    bool        `json:"reverse"`
}

// GetKVTableRows returns an object containing rows from the specified table.
func (c *ChainAPI) GetKVTableRows(props GetKVTableRowsProps) (GetKVTableRowsResponse, error) {
	var r = GetKVTableRowsResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_kv_table_rows"), props, &r)

	return r, err
}

type ABIJSONToBinProps struct {
	Code   string      `json:"code,omitempty"`
	Action string      `json:"action,omitempty"`
	Args   interface{} `json:"args,omitempty"`
}

// ABIJsonToBin returns an object containing the serialized action data.
func (c *ChainAPI) ABIJsonToBin(props ABIJSONToBinProps) (ABIJsonToBinResponse, error) {
	var r = ABIJsonToBinResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "abi_json_to_bin"), props, &r)

	return r, err
}

type ABIBinToJSONProps struct {
	Code    string `json:"code"`
	Action  string `json:"action"`
	Binargs string `json:"binargs"`
}

// ABIBinToJson returns an object containing the deserialized action data.
func (c *ChainAPI) ABIBinToJson(props ABIBinToJSONProps) (ABIBinToJsonResponse, error) {
	var r = ABIBinToJsonResponse("")
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "abi_bin_to_json"), props, &r)

	return r, err
}

type GetCodeProps struct {
	AccountName string `json:"account_name"`
	CodeAsWASM  int64  `json:"code_as_wasm"`
}

// GetCode returns an object containing the smart contract WASM code.
func (c *ChainAPI) GetCode(props GetCodeProps) (GetCodeResponse, error) {
	var r = GetCodeResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_code"), props, &r)

	return r, err
}

// GetRawABI returns an object containing smart contract abi.
func (c *ChainAPI) GetRawABI(props AccountProps) (GetRawABIResponse, error) {
	var r = GetRawABIResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_raw_abi"), props, &r)

	return r, err
}

type GetActivatedProtocolFeaturesProps struct {
	Params Params `json:"params"`
}

type Params struct {
	LowerBound       int64 `json:"lower_bound,omitempty"`
	UpperBound       int64 `json:"upper_bound,omitempty"`
	Limit            int64 `json:"limit,omitempty"`
	SearchByBlockNum bool  `json:"search_by_block_num"`
	Reverse          bool  `json:"reverse"`
}

// GetActivatedProtocolFeatures retreives the activated protocol features for producer node.
func (c *ChainAPI) GetActivatedProtocolFeatures(props GetActivatedProtocolFeaturesProps) (GetActivatedProtocolFeaturesResponse, error) {
	var r = GetActivatedProtocolFeaturesResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_activated_protocol_features"), props, &r)

	return r, err
}

type GetAccountsByAuthorizersProps struct {
	Accounts []string `json:"accounts"`
	Keys     []string `json:"keys"`
}

// GetAccountsByAuthorizers, given a set of account names and public keys, find all account permission authorities that are, in part or whole, satisfiable.
func (c *ChainAPI) GetAccountsByAuthorizers(props GetAccountsByAuthorizersProps) (GetAccountsByAuthorizersResponse, error) {
	var r = GetAccountsByAuthorizersResponse{}
	err := request(c.Client, urljoin.UrlJoin(c.ApiUrl, "get_accounts_by_authorizers"), props, &r)

	return r, err
}
