package eosrpc

type DBSizeGetResponse struct {
	FreeBytes int64   `json:"free_bytes"`
	UsedBytes int64   `json:"used_bytes"`
	Size      int64   `json:"size"`
	Indices   []Index `json:"indices"`
}

type Index struct {
	Index    string `json:"index"`
	RowCount int64  `json:"row_count"`
}
