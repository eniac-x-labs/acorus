package main

import (
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/params"

	"github.com/cornerstone-labs/acorus"
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
		EnvVars: []string{"ACORUS_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"ACORUS_MIGRATIONS_DIR"},
	}
	ChainBridgeFlag = &cli.StringFlag{
		Name:    "chain-bridge",
		Value:   "optimism",
		Usage:   "sync chain bridge event",
		EnvVars: []string{"ACORUS_CHAIN_BRIDGE"},
	}
)

func runAcorus(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "acorus")
	oplog.SetGlobalLogHandler(log.GetHandler())
	chainBridge := ctx.String(ChainBridgeFlag.Name)

	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name), chainBridge)
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
	if err != nil {
		log.Error("failed to connect to redis", "err", err)
		return err
	}
	defer redis.Close()

	newAcorus, err := acorus.NewAcorus(log, db, redis, cfg.Chain, cfg.OpContracts, cfg.HTTPServer, cfg.MetricsServer, chainBridge)
	if err != nil {
		log.Error("failed to create acorus", "err", err)
		return err
	}
	return newAcorus.Run(ctx.Context)
}

func runApi(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "service")
	oplog.SetGlobalLogHandler(log.GetHandler())
	chainBridge := ctx.String(ChainBridgeFlag.Name)
	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name), chainBridge)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	var db *database.DB
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
	chainBridge := ctx.String(ChainBridgeFlag.Name)
	cfg, err := config.LoadConfig(log, ctx.String(ConfigFlag.Name), chainBridge)
	migrationsDir := ctx.String(MigrationsFlag.Name)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}

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
	flags = append(flags, oplog.CLIFlags("ACORUS")...)
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	migrationFlags = append(migrationFlags, oplog.CLIFlags("ACORUS")...)
	IndexFlags := []cli.Flag{ChainBridgeFlag, ConfigFlag}
	IndexFlags = append(IndexFlags, oplog.CLIFlags("ACORUS")...)
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
				Flags:       IndexFlags,
				Description: "Runs the indexing service",
				Action:      runAcorus,
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
