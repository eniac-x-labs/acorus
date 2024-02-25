package database

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/database/relation"
	"github.com/cornerstone-labs/acorus/database/relayer"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"time"

	"github.com/cornerstone-labs/acorus/common/logs"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	_ "github.com/cornerstone-labs/acorus/database/utils/serializers"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type DB struct {
	gorm                *gorm.DB
	CreateTable         worker.CreateTableDB
	Blocks              common.BlocksDB
	ContractEvents      event.ContractEventsDB
	WithdrawProven      event.WithdrawProvenDB
	WithdrawFinalized   event.WithdrawFinalizedDB
	StateRoots          worker.StateRootDB
	L2ToL1              worker.L2ToL1DB
	L1ToL2              worker.L1ToL2DB
	MsgSentRelationD    relation.MsgSentRelationDB
	MsgHashRelationD    relation.MsgHashRelationDB
	RelayRelationD      relation.RelayRelationDB
	RelayMessage        event.RelayMessageDB
	StakeRecord         relayer.StakingRecordDB
	BridgeRecord        relayer.BridgeRecordDB
	BridgeMsgSent       relayer.BridgeMsgSentDB
	BridgeMsgHash       relayer.BridgeMsgHashDB
	BridgeClaim         relayer.BridgeClaimDB
	BridgeFinalize      relayer.BridgeFinalizeDB
	BridgeBlockListener relayer.BridgeBlockListenerDB
}

func NewDB(ctx context.Context, dbConfig config.Database) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.DbHost, dbConfig.DbName)
	if dbConfig.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.DbPort)
	}
	if dbConfig.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.DbUser)
	}
	if dbConfig.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.DbPassword)
	}
	DbLogger := logger.New(
		logs.New(), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	gormConfig := gorm.Config{
		Logger:                 DbLogger,
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm:                gorm,
		CreateTable:         worker.NewCreateTableDB(gorm),
		Blocks:              common.NewBlocksDB(gorm),
		WithdrawProven:      event.NewWithdrawProvenDB(gorm),
		WithdrawFinalized:   event.NewWithdrawFinalizedDB(gorm),
		ContractEvents:      event.NewContractEventsDB(gorm),
		StateRoots:          worker.NewStateRootDB(gorm),
		L1ToL2:              worker.NewL1ToL2DB(gorm),
		L2ToL1:              worker.NewL21ToL1DB(gorm),
		MsgSentRelationD:    relation.NewMsgSentRelationStructDB(gorm),
		MsgHashRelationD:    relation.NewMsgHashRelationDB(gorm),
		RelayRelationD:      relation.NewEvmRelayRelationDB(gorm),
		RelayMessage:        event.NewRelayMessageDB(gorm),
		StakeRecord:         relayer.NewStakingRecordDB(gorm),
		BridgeRecord:        relayer.NewBridgeRecordDB(gorm),
		BridgeMsgSent:       relayer.NewBridgeMsgSentDB(gorm),
		BridgeMsgHash:       relayer.NewBridgeMsgHashDB(gorm),
		BridgeClaim:         relayer.NewBridgeClaimDB(gorm),
		BridgeFinalize:      relayer.NewBridgeFinalizeDB(gorm),
		BridgeBlockListener: relayer.NewBlockListenerDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(gorm *gorm.DB) error {
		txDB := &DB{
			gorm:                gorm,
			Blocks:              common.NewBlocksDB(gorm),
			WithdrawProven:      event.NewWithdrawProvenDB(gorm),
			WithdrawFinalized:   event.NewWithdrawFinalizedDB(gorm),
			ContractEvents:      event.NewContractEventsDB(gorm),
			L1ToL2:              worker.NewL1ToL2DB(gorm),
			L2ToL1:              worker.NewL21ToL1DB(gorm),
			StateRoots:          worker.NewStateRootDB(gorm),
			MsgSentRelationD:    relation.NewMsgSentRelationStructDB(gorm),
			MsgHashRelationD:    relation.NewMsgHashRelationDB(gorm),
			RelayRelationD:      relation.NewEvmRelayRelationDB(gorm),
			RelayMessage:        event.NewRelayMessageDB(gorm),
			StakeRecord:         relayer.NewStakingRecordDB(gorm),
			BridgeRecord:        relayer.NewBridgeRecordDB(gorm),
			BridgeMsgSent:       relayer.NewBridgeMsgSentDB(gorm),
			BridgeMsgHash:       relayer.NewBridgeMsgHashDB(gorm),
			BridgeClaim:         relayer.NewBridgeClaimDB(gorm),
			BridgeFinalize:      relayer.NewBridgeFinalizeDB(gorm),
			BridgeBlockListener: relayer.NewBlockListenerDB(gorm),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
