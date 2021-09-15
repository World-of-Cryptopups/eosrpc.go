package eosrpc

import "github.com/TheBoringDude/go-urljoin"

const PRODUCER_API = "producer"

// Pause a producer node.
func (p *ProducerAPI) Pause() error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "pause"), nil, nil)
}

// Resume a producer node.
func (p *ProducerAPI) Resume() error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "resume"), nil, nil)
}

// Paused retrieves a paused status for a producer node.
func (p *ProducerAPI) Paused() (bool, error) {
	var r bool
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "paused"), nil, &r)

	return r, err
}

// GetRuntimeOptions retrieves run time options for producer node.
func (p *ProducerAPI) GetRuntimeOptions() (GetRuntimeOptionsResponse, error) {
	var r = GetRuntimeOptionsResponse{}
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "get_runtime_options"), nil, &r)

	return r, err
}

type UpdateRuntimeProps struct {
	Options Options `json:"options"`
}

type Options struct {
	MaxTransactionTime                    int64 `json:"max_transaction_time"`
	MaxIrreversibleBlockAge               int64 `json:"max_irreversible_block_age"`
	ProduceTimeOffsetUs                   int64 `json:"produce_time_offset_us"`
	LastBlockTimeOffsetUs                 int64 `json:"last_block_time_offset_us"`
	MaxScheduledTransactionTimePerBlockMS int64 `json:"max_scheduled_transaction_time_per_block_ms"`
	SubjectiveCPULeewayUs                 int64 `json:"subjective_cpu_leeway_us"`
	IncomingDeferRatio                    int64 `json:"incoming_defer_ratio"`
}

// Update run time options for producer node.
func (p *ProducerAPI) UpdateRuntimeOptions(props UpdateRuntimeProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "update_runtime_options"), props, nil)
}

// GetGreylist retrieves the greylist for producer node.
func (p *ProducerAPI) GetGreylist() ([]string, error) {
	var r []string
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "get_greylist"), nil, &r)

	return r, err
}

type GreylistActionProps struct {
	Params GreylistParams `json:"params"`
}

type GreylistParams struct {
	Accounts []string `json:"accounts"`
}

// AddGreylistAccounts adds accounts to grey list for producer node.
func (p *ProducerAPI) AddGreylistAccounts(props GreylistActionProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "add_greylist_accounts"), props, nil)
}

// RemoveGreylistAccounts removes accounts to grey list for producer node.
func (p *ProducerAPI) RemoveGreylistAccounts(props GreylistActionProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "add_greylist_accounts"), props, nil)
}

// GetWhitelistBlacklist retrieves the whitelist and blacklist for producer node.
func (p *ProducerAPI) GetWhitelistBlacklist() (GetWhitelistBlacklist, error) {
	var r = GetWhitelistBlacklist{}
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "get_whitelist_blacklist"), nil, &r)

	return r, err
}

type SetWhitelistBlacklistProps struct {
	Params WhitelistBlacklistParams `json:"params"`
}

type WhitelistBlacklistParams struct {
	ActorWhitelist    []string   `json:"actor_whitelist"`
	ActorBlacklist    []string   `json:"actor_blacklist"`
	ContractWhitelist []string   `json:"contract_whitelist"`
	ContractBlacklist []string   `json:"contract_blacklist"`
	ActionBlacklist   [][]string `json:"action_blacklist"`
	KeyBlacklist      []string   `json:"key_blacklist"`
}

// SetWhitelistBlacklist sets the whitelist and blacklist for producer node.
func (p *ProducerAPI) SetWhitelistBlacklist(props SetWhitelistBlacklistProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "set_whitelist_blacklist"), props, nil)
}

type CreateSnapshotProps struct {
	Next Next `json:"next"`
}

type Next struct {
	HeadBlockID   string `json:"head_block_id"`
	HeadBlockNum  int64  `json:"head_block_num"`
	HeadBlockTime string `json:"head_block_time"`
	Version       int64  `json:"version"`
	SnapshotName  string `json:"snapshot_name"`
}

// Creates a snapshot for producer node.
func (p *ProducerAPI) CreateSnapshot(props CreateSnapshotProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "create_snapshot"), props, nil)
}

// GetIntegrityHash retrieves the integrity hash for producer node.
func (p *ProducerAPI) GetIntegrityHash(props CreateSnapshotProps) (GetIntegrityHashResponse, error) {
	var r = GetIntegrityHashResponse{}
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "get_integrity_hash"), props, nil)

	return r, err
}

type ScheduleProtoFeatActiProps struct {
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	ProtocolFeaturesToActivate []string `json:"protocol_features_to_activate"`
}

// Schedule protocol feature activation for producer node.
func (p *ProducerAPI) ScheduleProtocolFeatureActivations(props CreateSnapshotProps) error {
	return request(p.Client, urljoin.UrlJoin(p.ApiUrl, "schedule_protocol_feature_activations"), props, nil)
}

type GetSupportedProtoFeatProps struct {
	Params SuppProtoParams `json:"params"`
}

type SuppProtoParams struct {
	ExcludeDisabled      bool `json:"exclude_disabled"`
	ExcludeUnactivatable bool `json:"exclude_unactivatable"`
}

// GetSupportedProtocolFeatures retreives supported protocol features for producer node.
func (p *ProducerAPI) GetSupportedProtocolFeatures(props GetSupportedProtoFeatProps) ([]string, error) {
	var r []string
	err := request(p.Client, urljoin.UrlJoin(p.ApiUrl, "get_supported_protocol_features"), props, &r)

	return r, err
}
