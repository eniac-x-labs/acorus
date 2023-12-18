package models

type QueryDWParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryPageListParams struct {
	Page     int
	PageSize int
	Order    string
}

type QueryByIdParams struct {
	Id uint64
}

type QueryByIndexParams struct {
	Index uint64
}
