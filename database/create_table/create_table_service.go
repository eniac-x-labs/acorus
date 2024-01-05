package create_table

import (
	"fmt"

	"github.com/cornerstone-labs/acorus/database"
)

func CreateTableFromTemplate(chainId string, db *database.DB) {
	createBlockHeaders(chainId, db)
	createContractEvents(chainId, db)
	createTransactions(chainId, db)
	createL1ToL2(chainId, db)
	createL2ToL1(chainId, db)
	createBatches(chainId, db)
	createStateRoot(chainId, db)
}

func createBlockHeaders(chainId string, db *database.DB) {
	tableName := "template_block_headers"
	tableNameByChainId := fmt.Sprintf("block_headers_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createContractEvents(chainId string, db *database.DB) {
	tableName := "template_contract_events"
	tableNameByChainId := fmt.Sprintf("contract_events_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createTransactions(chainId string, db *database.DB) {
	tableName := "template_transactions"
	tableNameByChainId := fmt.Sprintf("transactions_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createL1ToL2(chainId string, db *database.DB) {
	tableName := "template_l1_to_l2"
	tableNameByChainId := fmt.Sprintf("l1_to_l2_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createL2ToL1(chainId string, db *database.DB) {
	tableName := "template_l2_to_l1"
	tableNameByChainId := fmt.Sprintf("l2_to_l1_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createBatches(chainId string, db *database.DB) {
	tableName := "template_batches"
	tableNameByChainId := fmt.Sprintf("l2_to_l1_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createStateRoot(chainId string, db *database.DB) {
	tableName := "template_state_root"
	tableNameByChainId := fmt.Sprintf("state_root_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}
