package models

type DataStoreListItem struct {
	ID        uint64 `json:"data_store_id"`
	DataSize  uint64 `json:"data_size"`
	Status    bool   `json:"status"`
	Timestamp uint64 `json:"timestamp"`
	DaHash    string `json:"da_hash"`
}

type DataStoreListResponse struct {
	Current int                 `json:"Current"`
	Size    int                 `json:"Size"`
	Total   int64               `json:"Total"`
	Records []DataStoreListItem `json:"Records"`
}
