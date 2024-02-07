package models

import (
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/cornerstone-labs/acorus/database/worker"
)

type QueryDWParams struct {
	ChainId  string
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QuerySRParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryBRParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryPageParams struct {
	ChainId  string
	Page     int
	PageSize int
	Order    string
}

type QueryIdParams struct {
	ChainId string
	Id      uint64
}

type QueryIndexParams struct {
	ChainId string
	Index   uint64
}

type DepositsResponse struct {
	Current int             `json:"Current"`
	Size    int             `json:"Size"`
	Total   int64           `json:"Total"`
	Records []worker.L1ToL2 `json:"Records"`
}

type WithdrawsResponse struct {
	Current int             `json:"Current"`
	Size    int             `json:"Size"`
	Total   int64           `json:"Total"`
	Records []worker.L2ToL1 `json:"Records"`
}

type StakingResponse struct {
	Current int                     `json:"Current"`
	Size    int                     `json:"Size"`
	Total   int64                   `json:"Total"`
	Records []relayer.StakingRecord `json:"Records"`
}

type BridgeResponse struct {
	Current int                     `json:"Current"`
	Size    int                     `json:"Size"`
	Total   int64                   `json:"Total"`
	Records []relayer.BridgeRecords `json:"Records"`
}
