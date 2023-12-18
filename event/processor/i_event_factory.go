package processor

import (
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type IEventFactory interface {
	GetChainUnpackService(eventParam GetEventFactoryParams) IEventUnpack
}

type GetEventFactoryParams struct {
	Log         log.Logger
	Db          *database.DB
	L1Syncer    *synchronizer.L1Sync
	ChainConfig config.ChainConfig
	ChainBridge string
}
