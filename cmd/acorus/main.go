package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"

	oplog "github.com/cornerstone-labs/acorus/common/log"
	"github.com/cornerstone-labs/acorus/common/opio"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		opio.BlockOnInterrupts()
		cancel()
	}()

	oplog.SetupDefaults()
	app := newCli(GitCommit, GitDate)
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
	}
}
