package polygon

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	common3 "github.com/cornerstone-labs/acorus/event/processor/common"
	"github.com/cornerstone-labs/acorus/event/processor/polygon/abi"
	"github.com/cornerstone-labs/acorus/event/processor/polygon/bridge"
	"github.com/cornerstone-labs/acorus/event/processor/polygon/utils"
	"github.com/cornerstone-labs/acorus/synchronizer"
)

type PolygonProcessor struct {
	common3.IProcessor
}

func NewBridgeProcessor(log log.Logger, db *database.DB, l1Sync *synchronizer.L1Sync, l2Sync *synchronizer.L2Sync, chainConfig config.ChainConfig, rpcConfig config.RPCsConfig, shutdown context.CancelCauseFunc) (*common3.IProcessor, error) {
	ethClient, _ := ethclient.Dial(rpcConfig.L1RPC)
	zkEVMClient, _ := ethclient.Dial(rpcConfig.L2RPC)
	latestL1Header, latestL2Header,
		latestFinalizedL1Header, latestFinalizedL2Header,
		err := common3.GetLastBlockHeardByChain(log, db, chainConfig)
	if err != nil {
		return nil, err
	}
	log.Info("detected indexed bridge state",
		"l1_block", latestL1Header, "l2_block", latestL2Header,
		"finalized_l1_block", latestFinalizedL1Header, "finalized_l2_block", latestFinalizedL2Header)

	resCtx, resCancel := context.WithCancel(context.Background())

	l1PolygonBridge, err := abi.NewPolygonzkevmbridge(utils.L1PolygonZKEVMBridgeAddr, ethClient)
	if err != nil {
		return nil, err
	}
	l2PolygonBridge, err := abi.NewPolygonzkevmbridge(utils.L2PolygonZKEVMBridgeAddr, zkEVMClient)
	if err != nil {
		return nil, err
	}

	return &common3.IProcessor{
		Log:                   log,
		Db:                    db,
		L1Sync:                l1Sync,
		L2Sync:                l2Sync,
		ResourceCtx:           resCtx,
		ResourceCancel:        resCancel,
		ChainConfig:           chainConfig,
		LastL1Header:          latestL1Header,
		LastL2Header:          latestL2Header,
		LastRollupL1Header:    latestL1Header,
		L1PolygonBridge:       l1PolygonBridge,
		L2PolygonBridge:       l2PolygonBridge,
		LastFinalizedL1Header: latestFinalizedL1Header,
		LastFinalizedL2Header: latestFinalizedL2Header,
		Tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (bp *PolygonProcessor) Start() error {

	bp.Log.Info("starting bridge processor...")
	bp.Tasks.Go(func() error {
		l1SyncUpdates := bp.L1Sync.Notify()
		for range l1SyncUpdates {
			err := bp.onL1Data()
			if err != nil {
				return err
			}
		}
		bp.Log.Info("no more l1 etl updates. shutting down l1 task")
		return nil
	})
	// start L2 worker
	bp.Tasks.Go(func() error {
		l2SyncUpdates := bp.L2Sync.Notify()
		for range l2SyncUpdates {
			err := bp.onL2Data()
			if err != nil {
				return err
			}
		}
		bp.Log.Info("no more l2 etl updates. shutting down l2 task")
		return nil
	})
	return nil

}

func (bp *PolygonProcessor) Close() error {
	bp.ResourceCancel()
	return bp.Tasks.Wait()
}

func (bp *PolygonProcessor) onL1Data() error {
	fromL1Height := new(big.Int).Add(bp.LastFinalizedL1Header.Number, bigint.One)
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(10000))
	if err := bp.Db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := bp.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (bp *PolygonProcessor) onL2Data() error {
	fromL2Height := new(big.Int).Add(bp.LastFinalizedL2Header.Number, bigint.One)
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(10000))
	if err := bp.Db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := bp.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (bp *PolygonProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l1Contracts := bp.ChainConfig.L1Contracts
	batchLog := bp.Log.New("start_number", fromL1Height, "end_number", toL1Height)
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := bp.Db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			batchLog.Error("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		l1ToL2s := make([]worker.L1ToL2, len(events))
		for _, contractEvent := range events {
			unpackEvent, unpackErr := bp.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				batchLog.Error("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l1ToL2s = append(l1ToL2s, *unpackEvent)
			}
		}
		if len(l1ToL2s) > 0 {
			saveErr := bp.Db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s)
			if saveErr != nil {
				batchLog.Error("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (bp *PolygonProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l2Contracts := bp.ChainConfig.L2Contracts
	batchLog := bp.Log.New("start_number", fromL1Height, "end_number", toL1Height)
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		events, err := bp.Db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			batchLog.Error("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		l2ToL1s := make([]worker.L2ToL1, len(events))
		for _, contractEvent := range events {
			unpackEvent, unpackErr := bp.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				batchLog.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l2ToL1s = append(l2ToL1s, *unpackEvent)
			}
		}
		if len(l2ToL1s) > 0 {
			saveErr := bp.Db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s)
			if saveErr != nil {
				batchLog.Error("failed to StoreL2ToL1Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (bp *PolygonProcessor) l1EventUnpack(event event.L1ContractEvent) (*worker.L1ToL2, error) {

	switch event.EventSignature.String() {
	case utils.DepositEventSignatureHash.String():
		l1Deposit, err := bridge.L1Deposit(bp.L1PolygonBridge, event)
		if err != nil {
			return nil, err
		}
		return l1Deposit, nil
	case utils.ClaimEventSignatureHash.String():
		L1WithdrawClaimed, err := bridge.L1WithdrawClaimed(bp.L1PolygonBridge, event, bp.Db)
		if err != nil {
			return nil, err
		}
		return L1WithdrawClaimed, nil
	}
	return nil, nil
}

func (bp *PolygonProcessor) l2EventUnpack(event event.L2ContractEvent) (*worker.L2ToL1, error) {
	switch event.EventSignature.String() {
	case utils.WithdrawEventSignatureHash.String():
		withdraw, err := bridge.L2Withdraw(bp.L2PolygonBridge, event)
		if err != nil {
			return nil, err
		}
		return withdraw, nil
	}
	return nil, nil
}
