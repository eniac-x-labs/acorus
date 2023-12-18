package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/utils"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/synchronizer/retry"
)

type DB struct {
	gorm *gorm.DB

	Blocks         common.BlocksDB
	ContractEvents event.ContractEventsDB
	L2ToL1         worker.L2ToL1DB
	L1ToL2         worker.L1ToL2DB
	StateRoots     worker.StateRootDB
}

func NewDB(log log.Logger, dbConfig config.DBConfig) (*DB, error) {
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 utils.NewLogger(log),
	}

	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after multiple retries: %w", err)
	}

	db := &DB{
		gorm:           gorm,
		Blocks:         common.NewBlocksDB(gorm),
		ContractEvents: event.NewContractEventsDB(gorm),
		L1ToL2:         worker.NewL1ToL2DB(gorm),
		L2ToL1:         worker.NewL21ToL1DB(gorm),
	}

	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		return fn(dbFromGormTx(tx))
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}

	return sql.Close()
}

func dbFromGormTx(tx *gorm.DB) *DB {
	return &DB{
		gorm:           tx,
		Blocks:         common.NewBlocksDB(tx),
		ContractEvents: event.NewContractEventsDB(tx),
		L1ToL2:         worker.NewL1ToL2DB(tx),
		L2ToL1:         worker.NewL21ToL1DB(tx),
	}
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		// Check for any walking error
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read the migration file content
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}

		// Execute the migration
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})

	return err
}
