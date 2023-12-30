package create_table

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/database"
)

type createTableService struct {
}

var CreateTableService *createTableService

func init() {
	CreateTableService = &createTableService{}
}

func (s *createTableService) CreateInit(chainId int64, db *database.DB) {
	createL2BlockHeaders(chainId, db)
	createL1ToL2(chainId, db)
	createL2ToL1(chainId, db)
	createStateRoot(chainId, db)
}

func createL2BlockHeaders(chainId int64, db *database.DB) {
	tableName := "base_l2_block_headers"
	tableNameByChainId := fmt.Sprintf("l2_block_headers_%d", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createL1ToL2(chainId int64, db *database.DB) {
	tableName := "base_l1_to_l2"
	tableNameByChainId := fmt.Sprintf("l1_to_l2_%d", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createL2ToL1(chainId int64, db *database.DB) {
	tableName := "base_l2_to_l1"
	tableNameByChainId := fmt.Sprintf("l2_to_l1_%d", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createStateRoot(chainId int64, db *database.DB) {
	tableName := "base_state_root"
	tableNameByChainId := fmt.Sprintf("state_root_%d", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}
