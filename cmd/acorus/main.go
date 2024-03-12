package main

import (
	"context"
	"os"

	"github.com/cornerstone-labs/acorus/common/opio"
	"github.com/ethereum/go-ethereum/log"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := newCli(GitCommit, GitDate)
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
		os.Exit(1)
	}
}
