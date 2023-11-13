package main

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"

	acorus "github.com/cornerstone-labs/acorus"
	oplog "github.com/cornerstone-labs/acorus/common/log"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/service"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./acorus.toml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"INDEXER_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"INDEXER_MIGRATIONS_DIR"},
	}
)

func runIndexer(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "acorus")
	oplog.SetGlobalLogHandler(log.GetHandler())
	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}

	db, err := database.NewDB(log, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer db.Close()

	redis, err := database.NewRedis(cfg.Redis)

	acorus, err := acorus.Newacorus(log, db, redis, cfg.Chain, cfg.RPCs, cfg.HTTPServer, cfg.MetricsServer)
	if err != nil {
		log.Error("failed to create acorus", "err", err)
		return err
	}

	log.Info("running acorus...")
	return acorus.Run(ctx.Context)
}

func runApi(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "service")
	oplog.SetGlobalLogHandler(log.GetHandler())
	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	var db *database.DB
	// slave database for service services
	if !cfg.SlaveDbEnable {
		db, err = database.NewDB(log, cfg.MasterDB)
		if err != nil {
			log.Error("failed to connect to master database", "err", err)
			return err
		}
	} else {
		db, err = database.NewDB(log, cfg.SlaveDB)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
	}
	defer db.Close()

	log.Info("running service...")
	api := service.NewApi(log, db.BridgeTransfers, db.L1ToL2, db.L2ToL1, db.Blocks, db.StateRoots, cfg.HTTPServer, cfg.MetricsServer)
	return api.Run(ctx.Context)
}

func runMigrations(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "migrations")
	oplog.SetGlobalLogHandler(log.GetHandler())
	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name))
	migrationsDir := ctx.String(MigrationsFlag.Name)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}

	// If slave database not enable, use master database
	if cfg.SlaveDbEnable {
		slaveDB, err := database.NewDB(log, cfg.SlaveDB)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
		defer slaveDB.Close()
		err = slaveDB.ExecuteSQLMigration(migrationsDir)
		if err != nil {
			log.Error("migrate slave database fail", "err", err)
			return err
		}
	}

	masterDB, err := database.NewDB(log, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to master database", "err", err)
		return err
	}
	defer masterDB.Close()

	err = masterDB.ExecuteSQLMigration(migrationsDir)
	if err != nil {
		log.Error("migrate master database fail", "err", err)
		return err
	}
	return nil
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := []cli.Flag{ConfigFlag}
	flags = append(flags, oplog.CLIFlags("INDEXER")...)
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	migrationFlags = append(migrationFlags, oplog.CLIFlags("INDEXER")...)
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "An acorus of all optimism events with a serving service layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "service",
				Flags:       flags,
				Description: "Runs the service service",
				Action:      runApi,
			},
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      runIndexer,
			},
			{
				Name:        "migrate",
				Flags:       migrationFlags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
