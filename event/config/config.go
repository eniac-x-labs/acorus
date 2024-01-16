package config

import (
	"context"
	"fmt"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
)

type UnpackConfig struct {
	Db             *database.DB
	CfgRpc         *config.RPC
	ResourceCtx    context.Context
	ResourceCancel context.CancelFunc
	Tasks          tasks.Group
}

func New(db *database.DB,
	cfg *config.RPC, shutdown context.CancelCauseFunc) *UnpackConfig {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &UnpackConfig{
		Db:             db,
		CfgRpc:         cfg,
		ResourceCtx:    resCtx,
		ResourceCancel: resCancel,
		Tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}
}
