package appchain

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"math/big"
	"time"

	"github.com/cornerstone-labs/acorus/appchain/bindings"
	"github.com/cornerstone-labs/acorus/appchain/unpack"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/appchain"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type L1AppChainListener struct {
	chainId          string
	contracts        []string
	tasks            tasks.Group
	loopInterval     time.Duration
	startHeight      uint64
	epoch            uint64
	db               *database.DB
	bridgeRpcService bridge.BridgeRpcService
}

func NewL1AppChainListener(chainId string,
	contracts []string,
	shutdown context.CancelCauseFunc,
	loopInterval time.Duration,
	startHeight uint64,
	epoch uint64,
	bridgeRpcService bridge.BridgeRpcService,
	db *database.DB) (*L1AppChainListener, error) {
	return &L1AppChainListener{
		chainId:   chainId,
		contracts: contracts,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval:     loopInterval,
		startHeight:      startHeight,
		epoch:            epoch,
		bridgeRpcService: bridgeRpcService,
		db:               db,
	}, nil
}

func (l *L1AppChainListener) Start() error {
	l1ListenerEventOn1 := time.NewTicker(l.loopInterval)
	l.tasks.Go(func() error {
		for range l1ListenerEventOn1.C {
			err := l.onL1Data()
			if err != nil {
				log.Error("L1AppChainListener", "err", err)
				continue
			}
		}
		return nil
	})
	l.tasks.Go(func() error {
		for range l1ListenerEventOn1.C {
			err := l.notifyRelayerBatch()
			if err != nil {
				log.Error("notifyRelayerBatch", "err", err)
				continue
			}
		}
		return nil
	})
	l.tasks.Go(func() error {
		for range l1ListenerEventOn1.C {
			err := l.notifyRelayerSingle()
			if err != nil {
				log.Error("notifyRelayerSingle", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (l *L1AppChainListener) onL1Data() error {
	lastListenBlock, err := l.db.AppChainLastBlock.GetLastBlockNumber(l.chainId)
	if err != nil {
		log.Error("appchain L1  failed to get last block heard", "err", err)
		return err
	}
	fromL2Height := bigint.Zero
	if lastListenBlock != nil {
		fromL2Height = new(big.Int).Add(lastListenBlock.BlockNumber, bigint.One)
	}
	if l.startHeight > 0 {
		if fromL2Height.Cmp(big.NewInt(int64(l.startHeight))) == -1 {
			fromL2Height = big.NewInt(int64(l.startHeight))
		}
	}
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(l.epoch)))
	chainLatestBlockHeader, err := l.db.Blocks.ChainLatestBlockHeader(l.chainId)
	if err != nil {
		return err
	}
	if chainLatestBlockHeader == nil {
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL2Height) == -1 {
		fromL2Height = chainLatestBlockHeader.Number
		toL2Height = chainLatestBlockHeader.Number
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL2Height) == -1 {
			toL2Height = chainLatestBlockHeader.Number
		}
	}
	if err := l.db.Transaction(func(tx *database.DB) error {
		eventsFetchErr := l.EventsFetch(fromL2Height, toL2Height)
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := appchain.AppChainLastBlock{
			ChainId:     l.chainId,
			BlockNumber: toL2Height,
		}
		updateErr := l.db.AppChainLastBlock.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			return updateErr
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (l *L1AppChainListener) EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainIdStr := l.chainId
	contracts := l.contracts
	for _, contract := range contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		log.Info("appchain L1 Fetch", "chainId", chainIdStr, "contract", common.HexToAddress(contract).String(),
			"fromL1Height", fromL1Height, "toL1Height", toL1Height)

		events, err := l.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Error("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := l.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (l *L1AppChainListener) eventUnpack(event event.ContractEvent) error {
	switch event.EventSignature.String() {
	case bindings.AppChainL1Abi.Events["UnstakeSingle"].ID.String():
		err := unpack.UnstakeSingle(event, l.db)
		return err
	case bindings.AppChainL1Abi.Events["UnstakeBatchRequest"].ID.String():
		err := unpack.UnstakeBatchRequest(event, l.db)
		return err
	case bindings.AppChainL1Abi.Events["UnstakeRequestClaimed"].ID.String():
		err := unpack.UnstakeRequestClaimed(event, l.db)
		return err
	}
	return nil
}

func (l *L1AppChainListener) notifyRelayerBatch() error {
	batchList := l.db.AppChainUnStakeBatch.ListAppChainUnStakeBatch()
	log.Info("notifyRelayerBatch", "batchList", batchList)
	for _, batch := range batchList {
		batchId := batch.BatchId
		destChainId := batch.DestChainId
		sourceChainId := batch.SourceChainId
		bridge := batch.Bridge
		txHash := batch.TxHash
		unstakeBatch, err := l.bridgeRpcService.UnstakeBatch(txHash.String(), batchId.String(), bridge.String(), sourceChainId, destChainId)
		if err != nil {
			return err
		}
		success := unstakeBatch.Success
		if success {
			err := l.db.AppChainUnStakeBatch.NotifyAppChainUnStakeBatch(txHash.String())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *L1AppChainListener) notifyRelayerSingle() error {
	singleList := l.db.AppChainUnStakeSingle.ListAppChainUnStakeSingle()
	log.Info("notifyRelayerSingle", "singleList", singleList)
	for _, single := range singleList {
		txHash := single.TxHash
		strategy := single.L2Strategy
		amount := single.LockedAmount
		staker := single.Staker
		unstakeSingle, err := l.bridgeRpcService.UnstakeSingle(staker.String(), strategy.String(), amount.String(), l.chainId, txHash.String())
		if err != nil {
			return err
		}
		success := unstakeSingle.Success
		if success {
			err := l.db.AppChainUnStakeSingle.NotifyAppChainUnStakeSingle(txHash.String())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
