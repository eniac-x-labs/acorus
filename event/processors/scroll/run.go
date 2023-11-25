package scroll

import (
	"context"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type ScBridgeProcessor struct {
	log log.Logger
	db  *database.DB

	l1Etl       *synchronizer.L1Sync
	chainConfig config.ChainConfig
	opContracts interface{}

	LatestL1Header *types.Header
	LatestL2Header *types.Header
}

func NewScrollBridgeProcessor(log log.Logger, db *database.DB, L1Syncer *synchronizer.L1Sync, chainConfig config.ChainConfig, opContracts interface{}) (*ScBridgeProcessor, error) {
	log = log.New("processor", "bridge")
	latestL1Header, err := db.BridgeTransactions.L1LatestBlockHeader()
	if err != nil {
		return nil, err
	}
	latestL2Header, err := db.BridgeTransactions.L2LatestBlockHeader()
	if err != nil {
		return nil, err
	}

	var l1Header, l2Header *types.Header
	if latestL1Header == nil && latestL2Header == nil {
		log.Info("no indexed state, starting from rollup genesis")
	} else {
		l1Height, l2Height := bigint.Zero, bigint.Zero
		if latestL1Header != nil {
			l1Height = latestL1Header.Number
			l1Header = latestL1Header.RLPHeader.Header()
		}
		if latestL2Header != nil {
			l2Height = latestL2Header.Number
			l2Header = latestL2Header.RLPHeader.Header()
		}
		log.Info("detected latest indexed bridge state", "l1_block_number", l1Height, "l2_block_number", l2Height)
	}
	return &ScBridgeProcessor{log, db, L1Syncer, chainConfig, opContracts, l1Header, l2Header}, nil
}

func (b *ScBridgeProcessor) Start(ctx context.Context) error {
	done := ctx.Done()
	l1EtlUpdates := b.l1Etl.Notify()
	startup := make(chan interface{}, 1)
	startup <- nil

	b.log.Info("starting bridge processor...")
	for {
		select {
		case <-done:
			b.log.Info("stopping bridge processor")
			return nil

		// Tickers
		case <-startup:
		case <-l1EtlUpdates:
		}

		//b.run()
	}
}
