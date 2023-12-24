package main

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/params"

	"github.com/cornerstone-labs/acorus"
	"github.com/cornerstone-labs/acorus/common/cliapp"
	oplog "github.com/cornerstone-labs/acorus/common/log"
	"github.com/cornerstone-labs/acorus/common/opio"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	flag2 "github.com/cornerstone-labs/acorus/flag"
	"github.com/cornerstone-labs/acorus/service"
)

func runIndexer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "acorus")
	oplog.SetGlobalLogHandler(log.GetHandler())
	log.Info("running indexer...")

	cfg, err := config.LoadConfig(log, ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}

	return acorus.NewAcorus(ctx.Context, log, &cfg, shutdown)
}

func runApi(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "api")
	oplog.SetGlobalLogHandler(log.GetHandler())
	log.Info("running api...")
	cfg, err := config.LoadConfig(log, ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return service.NewApi(ctx.Context, log, &cfg)
}

func runExporter(ctx *cli.Context) error {
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "exporter")
	oplog.SetGlobalLogHandler(log.GetHandler())
	cfg, err := config.LoadConfig(log, ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	log.Info("running exporter...", "cfg", cfg)
	// todo please write exporter here
	return nil
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log := oplog.NewLogger(oplog.AppOut(ctx), oplog.ReadCLIConfig(ctx)).New("role", "migrations")
	oplog.SetGlobalLogHandler(log.GetHandler())
	log.Info("running migrations...")
	cfg, err := config.LoadConfig(log, ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, log, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer db.Close()
	return db.ExecuteSQLMigration(cfg.Migrations)
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := oplog.CLIFlags("ACORUS")
	flags = append(flags, flag2.Flags...)
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Runs the api service",
				Action:      cliapp.LifecycleCmd(runApi),
			},
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runIndexer),
			},
			{
				Name:        "migrate",
				Flags:       flags,
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
