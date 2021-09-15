package eosrpc

// Response for GetAccount
type GetAccountResponse struct {
	AccountName            string              `json:"account_name"`
	HeadBlockNum           int64               `json:"head_block_num"`
	HeadBlockTime          string              `json:"head_block_time"`
	Privileged             bool                `json:"privileged"`
	LastCodeUpdate         string              `json:"last_code_update"`
	Created                string              `json:"created"`
	CoreLiquidBalance      string              `json:"core_liquid_balance"`
	RAMQuota               int64               `json:"ram_quota"`
	NetWeight              int64               `json:"net_weight"`
	CPUWeight              int64               `json:"cpu_weight"`
	NetLimit               Limit               `json:"net_limit"`
	CPULimit               Limit               `json:"cpu_limit"`
	RAMUsage               int64               `json:"ram_usage"`
	Permissions            []PermissionElement `json:"permissions"`
	TotalResources         TotalResources      `json:"total_resources"`
	SelfDelegatedBandwidth interface{}         `json:"self_delegated_bandwidth"`
	RefundRequest          interface{}         `json:"refund_request"`
	VoterInfo              VoterInfo           `json:"voter_info"`
	RexInfo                interface{}         `json:"rex_info"`
	SubjectiveCPUBillLimit Limit               `json:"subjective_cpu_bill_limit"`
}

type Limit struct {
	Used      int64 `json:"used"`
	Available int64 `json:"available"`
	Max       int64 `json:"max"`
}

type PermissionElement struct {
	PermName     string       `json:"perm_name"`
	Parent       string       `json:"parent"`
	RequiredAuth RequiredAuth `json:"required_auth"`
}

type RequiredAuth struct {
	Threshold int64         `json:"threshold"`
	Keys      []interface{} `json:"keys"`
	Accounts  []Account     `json:"accounts"`
	Waits     []interface{} `json:"waits"`
}

type Account struct {
	Permission AccountPermission `json:"permission"`
	Weight     int64             `json:"weight"`
}

type AccountPermission struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

type TotalResources struct {
	Owner     string `json:"owner"`
	NetWeight string `json:"net_weight"`
	CPUWeight string `json:"cpu_weight"`
	RAMBytes  int64  `json:"ram_bytes"`
}

type VoterInfo struct {
	Owner                      string        `json:"owner"`
	Proxy                      string        `json:"proxy"`
	Producers                  []interface{} `json:"producers"`
	Staked                     string        `json:"staked"`
	UnpaidVoteshare            string        `json:"unpaid_voteshare"`
	UnpaidVoteshareLastUpdated string        `json:"unpaid_voteshare_last_updated"`
	UnpaidVoteshareChangeRate  string        `json:"unpaid_voteshare_change_rate"`
	LastClaimTime              string        `json:"last_claim_time"`
	LastVoteWeight             string        `json:"last_vote_weight"`
	ProxiedVoteWeight          string        `json:"proxied_vote_weight"`
	IsProxy                    int64         `json:"is_proxy"`
	Flags1                     int64         `json:"flags1"`
	Reserved2                  int64         `json:"reserved2"`
	Reserved3                  string        `json:"reserved3"`
}

// Response for GetBlock
type GetBlockResponse struct {
	Timestamp           string               `json:"timestamp"`
	Producer            string               `json:"producer"`
	Confirmed           int64                `json:"confirmed"`
	Previous            string               `json:"previous"`
	TransactionMroot    string               `json:"transaction_mroot"`
	ActionMroot         string               `json:"action_mroot"`
	ScheduleVersion     int64                `json:"schedule_version"`
	NewProducers        NewProducers         `json:"new_producers"`
	HeaderExtensions    []int64              `json:"header_extensions"`
	NewProtocolFeatures []NewProtocolFeature `json:"new_protocol_features"`
	ProducerSignature   string               `json:"producer_signature"`
	Transactions        []BlockTransaction   `json:"transactions"`
	BlockExtensions     []int64              `json:"block_extensions"`
	ID                  string               `json:"id"`
	BlockNum            int64                `json:"block_num"`
	RefBlockPrefix      int64                `json:"ref_block_prefix"`
}

type NewProducers struct {
	Version   int64      `json:"version"`
	Producers []Producer `json:"producers"`
}

type Producer struct {
	ProducerName    string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}

type NewProtocolFeature []interface{}

type BlockTransaction struct {
	Status        string `json:"status"`
	CPUUsageUs    int64  `json:"cpu_usage_us"`
	NetUsageWords int64  `json:"net_usage_words"`
	Trx           string `json:"trx"`
}

// Response for GetBlockInfo
type GetBlockInfoResponse struct {
	BlockNum          int64  `json:"block_num"`
	RefBlockNum       int64  `json:"ref_block_num"`
	ID                string `json:"id"`
	Timestamp         string `json:"timestamp"`
	Producer          string `json:"producer"`
	Confirmed         int64  `json:"confirmed"`
	Previous          string `json:"previous"`
	TransactionMroot  string `json:"transaction_mroot"`
	ActionMroot       string `json:"action_mroot"`
	ScheduleVersion   int64  `json:"schedule_version"`
	ProducerSignature string `json:"producer_signature"`
	RefBlockPrefix    int64  `json:"ref_block_prefix"`
}

// Response for GetInfo
type GetInfoResponse struct {
	ServerVersion             string `json:"server_version"`
	ChainID                   string `json:"chain_id"`
	HeadBlockNum              int64  `json:"head_block_num"`
	LastIrreversibleBlockNum  int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID   string `json:"last_irreversible_block_id"`
	HeadBlockID               string `json:"head_block_id"`
	HeadBlockTime             string `json:"head_block_time"`
	HeadBlockProducer         string `json:"head_block_producer"`
	VirtualBlockCPULimit      int64  `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit      int64  `json:"virtual_block_net_limit"`
	BlockCPULimit             int64  `json:"block_cpu_limit"`
	BlockNetLimit             int64  `json:"block_net_limit"`
	ServerVersionString       string `json:"server_version_string"`
	ForkDBHeadBlockNum        int64  `json:"fork_db_head_block_num"`
	ForkDBHeadBlockID         string `json:"fork_db_head_block_id"`
	ServerFullVersionString   string `json:"server_full_version_string"`
	LastIrreversibleBlockTime string `json:"last_irreversible_block_time"`
}

// Response for GetBlockHeaderState
type GetBlockHeaderStateResponse struct {
	ID                               string          `json:"id"`
	BlockNum                         int64           `json:"block_num"`
	Header                           Header          `json:"header"`
	DposProposedIrreversibleBlocknum string          `json:"dpos_proposed_irreversible_blocknum"`
	DposIrreversibleBlocknum         string          `json:"dpos_irreversible_blocknum"`
	BFTIrreversibleBlocknum          string          `json:"bft_irreversible_blocknum"`
	PendingScheduleLIBNum            string          `json:"pending_schedule_lib_num"`
	PendingScheduleHash              string          `json:"pending_schedule_hash"`
	PendingSchedule                  ActiveSchedule  `json:"pending_schedule"`
	ActiveSchedule                   ActiveSchedule  `json:"active_schedule"`
	BlockrootMerkle                  BlockrootMerkle `json:"blockroot_merkle"`
	ProducerToLastProduced           [][]string      `json:"producer_to_last_produced"`
	ProducerToLastImpliedIrb         [][]string      `json:"producer_to_last_implied_irb"`
	BlockSigningKey                  string          `json:"block_signing_key"`
	ConfirmCount                     []string        `json:"confirm_count"`
	Confirmations                    []interface{}   `json:"confirmations"`
}

type ActiveSchedule struct {
	Version   int64      `json:"version"`
	Producers []Producer `json:"producers"`
}

type BlockrootMerkle struct {
	ActiveNodes []string `json:"_active_nodes"`
	NodeCount   string   `json:"_node_count"`
}

type Header struct {
	Timestamp           string               `json:"timestamp"`
	Producer            string               `json:"producer"`
	Confirmed           int64                `json:"confirmed"`
	Previous            string               `json:"previous"`
	TransactionMroot    string               `json:"transaction_mroot"`
	ActionMroot         string               `json:"action_mroot"`
	ScheduleVersion     int64                `json:"schedule_version"`
	NewProducers        ActiveSchedule       `json:"new_producers"`
	HeaderExtensions    []int64              `json:"header_extensions"`
	NewProtocolFeatures []NewProtocolFeature `json:"new_protocol_features"`
	ProducerSignature   string               `json:"producer_signature"`
	Transactions        []BlockTransaction   `json:"transactions"`
	BlockExtensions     []int64              `json:"block_extensions"`
	ID                  string               `json:"id"`
	BlockNum            int64                `json:"block_num"`
	RefBlockPrefix      int64                `json:"ref_block_prefix"`
}

// Response for GetABI
type GetABIResponse struct {
	Version          string      `json:"version"`
	Types            []Type      `json:"types"`
	Structs          []Struct    `json:"structs"`
	Actions          []ABIAction `json:"actions"`
	Tables           []ABITable  `json:"tables"`
	ABIExtensions    [][]int64   `json:"abi_extensions"`
	ErrorMessages    []string    `json:"error_messages"`
	RicardianClauses []string    `json:"ricardian_clauses"`
	Variants         []string    `json:"variants"`
}

type ABIAction struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	RicardianContract string `json:"ricardian_contract"`
}

type Struct struct {
	Name   string  `json:"name"`
	Base   string  `json:"base"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ABITable struct {
	Name      string   `json:"name"`
	IndexType string   `json:"index_type"`
	KeyNames  []string `json:"key_names"`
	KeyTypes  []string `json:"key_types"`
	Type      string   `json:"type"`
}

type Type struct {
	NewTypeName string `json:"new_type_name"`
	Type        string `json:"type"`
}

// Response for GetCurrencyBalance
type GetCurrencyBalanceResponse []string

// Response for GetCurrencyStats
type GetCurrencyStatsResponse interface{}

// Response for GetRequiredKeys
type GetRequiredKeysResponse struct {
	Transaction   Transaction `json:"transaction"`
	AvailableKeys []string    `json:"available_keys"`
}

type Transaction struct {
	Expiration            string    `json:"expiration"`
	RefBlockNum           int64     `json:"ref_block_num"`
	RefBlockPrefix        int64     `json:"ref_block_prefix"`
	MaxNetUsageWords      string    `json:"max_net_usage_words"`
	MaxCPUUsageMS         string    `json:"max_cpu_usage_ms"`
	DelaySEC              int64     `json:"delay_sec"`
	ContextFreeActions    []Action  `json:"context_free_actions"`
	Actions               []Action  `json:"actions"`
	TransactionExtensions [][]int64 `json:"transaction_extensions"`
}

type Action struct {
	Account       string                 `json:"account"`
	Name          string                 `json:"name"`
	Authorization []Authorization        `json:"authorization"`
	Data          map[string]interface{} `json:"data"`
	HexData       string                 `json:"hex_data"`
}

type Authorization struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

// Response for GetProducers
type GetProducersResponse struct {
	Active   []Active `json:"active"`
	Pending  []Active `json:"pending"`
	Proposed []Active `json:"proposed"`
}

type Active struct {
	Version   int64      `json:"version"`
	Producers []Producer `json:"producers"`
}

// Response for GetRawCodeAndABI
type GetRawCodeAndABIResponse struct {
	AccountName string `json:"account_name"`
	WASM        string `json:"wasm"`
	ABI         string `json:"abi"`
}

// Response for GetScheduledTransaction
type GetScheduledTransactionResponse struct {
	Transactions []BlockTransaction `json:"transactions"`
}

// Response for GetTableByScope
type GetTableByScopeResponse struct {
	Rows []Row  `json:"rows"`
	More string `json:"more"`
}

type Row struct {
	Code  string `json:"code"`
	Scope string `json:"scope"`
	Table string `json:"table"`
	Payer string `json:"payer"`
	Count int64  `json:"count"`
}

// Response for GetTableRows
type GetTableRowsResponse struct {
	Rows []interface{} `json:"rows"`
}

// Response for GetKVTableRows
type GetKVTableRowsResponse struct {
	Rows []interface{} `json:"rows"`
}

// Response for ABIJsonToBin
type ABIJsonToBinResponse struct {
	Binargs string `json:"binargs"`
}

// Response for ABIBinToJson
type ABIBinToJsonResponse = string

// Response for GetCode
type GetCodeResponse struct {
	Name     string `json:"name"`
	CodeHash string `json:"code_hash"`
	Wast     string `json:"wast"`
	WASM     string `json:"wasm"`
	ABI      ABI    `json:"abi"`
}

type ABI struct {
	Version          string      `json:"version"`
	Types            []Type      `json:"types"`
	Structs          []Struct    `json:"structs"`
	Actions          []ABIAction `json:"actions"`
	Tables           []Table     `json:"tables"`
	ABIExtensions    [][]int64   `json:"abi_extensions"`
	ErrorMessages    []string    `json:"error_messages"`
	RicardianClauses []string    `json:"ricardian_clauses"`
	Variants         []string    `json:"variants"`
}

type Table struct {
	Name      string   `json:"name"`
	IndexType string   `json:"index_type"`
	KeyNames  []string `json:"key_names"`
	KeyTypes  []string `json:"key_types"`
	Type      string   `json:"type"`
}

// Response for GetRawABI
type GetRawABIResponse struct {
	AccountName string `json:"account_name"`
	CodeHash    string `json:"code_hash"`
	ABIHash     string `json:"abi_hash"`
	ABI         string `json:"abi"`
}

// Response for GetActivatedProtocolFeatures
type GetActivatedProtocolFeaturesResponse struct {
	ActivatedProtocolFeatures []string `json:"activated_protocol_features"`
	More                      int64    `json:"more"`
}

// Response for GetAccountsByAuthorizers
type GetAccountsByAuthorizersResponse struct {
	Accounts []AuthorizerAccount `json:"accounts"`
}

type AuthorizerAccount struct {
	AccountName    string `json:"account_name"`
	PermissionName string `json:"permission_name"`
	Authorizer     string `json:"authorizer"`
	Weight         int64  `json:"weight"`
	Threshold      int64  `json:"threshold"`
}
