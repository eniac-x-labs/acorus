package processor

import (
	"context"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
)

type IProcessor struct {
	Log            log.Logger
	Db             *database.DB
	Cfg            *config.Config
	ResourceCtx    context.Context
	ResourceCancel context.CancelFunc
	Tasks          tasks.Group
}

func (ip *IProcessor) Start() {

}
