package eosrpc

type GetRuntimeOptionsResponse struct {
	MaxTransactionTime                    int64 `json:"max_transaction_time"`
	MaxIrreversibleBlockAge               int64 `json:"max_irreversible_block_age"`
	ProduceTimeOffsetUs                   int64 `json:"produce_time_offset_us"`
	LastBlockTimeOffsetUs                 int64 `json:"last_block_time_offset_us"`
	MaxScheduledTransactionTimePerBlockMS int64 `json:"max_scheduled_transaction_time_per_block_ms"`
	SubjectiveCPULeewayUs                 int64 `json:"subjective_cpu_leeway_us"`
	IncomingDeferRatio                    int64 `json:"incoming_defer_ratio"`
}

type GetWhitelistBlacklist struct {
	ActorWhitelist    []string   `json:"actor_whitelist"`
	ActorBlacklist    []string   `json:"actor_blacklist"`
	ContractWhitelist []string   `json:"contract_whitelist"`
	ContractBlacklist []string   `json:"contract_blacklist"`
	ActionBlacklist   [][]string `json:"action_blacklist"`
	KeyBlacklist      []string   `json:"key_blacklist"`
}

type GetIntegrityHashResponse struct {
	IntegrityHash string `json:"integrity_hash"`
	HeadBlockID   string `json:"head_block_id"`
}
