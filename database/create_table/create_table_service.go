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
	createWithdrawProven(chainId, db)
	createWithdrawFinalized(chainId, db)
	createMsgSent(chainId, db)
	createRelay(chainId, db)
	createMshHash(chainId, db)
	createRelayMessage(chainId, db)
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

func createWithdrawProven(chainId string, db *database.DB) {
	tableName := "template_withdraw_proven"
	tableNameByChainId := fmt.Sprintf("withdraw_proven_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createWithdrawFinalized(chainId string, db *database.DB) {
	tableName := "template_withdraw_finalized"
	tableNameByChainId := fmt.Sprintf("withdraw_finalized_%s", chainId)
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

func createMsgSent(chainId string, db *database.DB) {
	tableName := "template_msg_sent"
	tableNameByChainId := fmt.Sprintf("msg_sent_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createRelay(chainId string, db *database.DB) {
	tableName := "template_relay"
	tableNameByChainId := fmt.Sprintf("relay_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createMshHash(chainId string, db *database.DB) {
	tableName := "template_msg_hash"
	tableNameByChainId := fmt.Sprintf("msg_hash_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createRelayMessage(chainId string, db *database.DB) {
	tableName := "template_relay_message"
	tableNameByChainId := fmt.Sprintf("relay_message_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}
