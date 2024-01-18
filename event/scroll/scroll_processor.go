package scroll

import (
	"context"
	"fmt"
	"time"

	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/scroll/bridge"
)

type ScrollEventProcessor struct {
	db             *database.DB
	cfgRpc         *config.RPC
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
	loopInterval   time.Duration
	l1StartHeight  *big.Int
	l2StartHeight  *big.Int
	epoch          uint64
}

func NewBridgeProcessor(db *database.DB,
	cfg *config.RPC, shutdown context.CancelCauseFunc,
	loopInterval time.Duration, epoch uint64) (*ScrollEventProcessor, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &ScrollEventProcessor{
		db:             db,
		cfgRpc:         cfg,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
		loopInterval: loopInterval,
		epoch:        epoch,
	}, nil
}

func (sp *ScrollEventProcessor) StartUnpack() error {
	tickerSyncer := time.NewTicker(sp.loopInterval)
	log.Println("starting scroll bridge processor...")
	sp.tasks.Go(func() error {
		for range tickerSyncer.C {
			err := sp.onL1Data()
			if err != nil {
				log.Println("no more l1 etl updates. shutting down l1 task")
				return err
			}
		}
		return nil
	})
	// start L2 worker
	sp.tasks.Go(func() error {
		for range tickerSyncer.C {
			err := sp.onL2Data()
			if err != nil {
				log.Println("no more l2 etl updates. shutting down l2 task")
				return err
			}
		}
		return nil
	})
	return nil
}

func (sp *ScrollEventProcessor) ClosetUnpack() error {
	sp.resourceCancel()
	return sp.tasks.Wait()
}

func (sp *ScrollEventProcessor) onL1Data() error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	if sp.l1StartHeight == nil {
		lastBlockHeard, err := sp.db.L1ToL2.LatestBlockHeader(chainIdStr)
		if err != nil {
			log.Println("l1 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		sp.l1StartHeight = lastBlockHeard.Number
		if sp.l1StartHeight.Cmp(big.NewInt(int64(sp.cfgRpc.L1EventUnpackBlock))) == -1 {
			sp.l1StartHeight = big.NewInt(int64(sp.cfgRpc.L1EventUnpackBlock))
		}
	} else {
		sp.l1StartHeight = new(big.Int).Add(sp.l1StartHeight, bigint.One)
	}

	fromL1Height := sp.l1StartHeight
	toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(int64(sp.epoch)))
	if err := sp.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := sp.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	sp.l1StartHeight = toL1Height
	return nil
}

func (sp *ScrollEventProcessor) onL2Data() error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	if sp.l2StartHeight == nil {
		lastBlockHeard, err := sp.db.L2ToL1.LatestBlockHeader(chainIdStr)
		if err != nil {
			log.Println("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			lastBlockHeard = &common2.BlockHeader{
				Number: big.NewInt(0),
			}
		}
		sp.l2StartHeight = lastBlockHeard.Number
	} else {
		sp.l2StartHeight = new(big.Int).Add(sp.l2StartHeight, bigint.One)
	}
	fromL2Height := sp.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(sp.epoch)))
	if err := sp.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := sp.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}
	sp.l2StartHeight = toL2Height
	return nil
}

func (sp *ScrollEventProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	l1Contracts := sp.cfgRpc.Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := sp.db.ContractEvents.ContractEventsWithFilter("1", contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		l1ToL2s := make([]worker.L1ToL2, 0)
		for _, contractEvent := range events {
			unpackEvent, unpackErr := sp.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l1ToL2s = append(l1ToL2s, *unpackEvent)
			}
		}
		if len(l1ToL2s) > 0 {
			saveErr := sp.db.L1ToL2.StoreL1ToL2Transactions(chainIdStr, l1ToL2s)
			if saveErr != nil {
				log.Println("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (sp *ScrollEventProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	l2Contracts := sp.cfgRpc.EventContracts
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
		events, err := sp.db.ContractEvents.ContractEventsWithFilter(chainIdStr, contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		l2ToL1s := make([]worker.L2ToL1, 0)
		for _, contractEvent := range events {
			unpackEvent, unpackErr := sp.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l2ToL1s = append(l2ToL1s, *unpackEvent)
			}
		}
		if len(l2ToL1s) > 0 {
			saveErr := sp.db.L2ToL1.StoreL2ToL1Transactions(chainIdStr, l2ToL1s)
			if saveErr != nil {
				log.Println("failed to StoreL2ToL1Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (sp *ScrollEventProcessor) l1EventUnpack(event event.ContractEvent) (*worker.L1ToL2, error) {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L1DepositETHSig.String():
		l1DepositETH, err := bridge.L1DepositETH(event)
		if err != nil {
			return nil, err
		}
		return l1DepositETH, nil
	case abi.L1DepositERC20Sig.String():
		L1DepositERC20, err := bridge.L1DepositERC20(event)
		if err != nil {
			return nil, err
		}
		return L1DepositERC20, nil
	case abi.L1DepositERC721Sig.String():
		L1DepositERC721, err := bridge.L1DepositERC721(event)
		if err != nil {
			return nil, err
		}
		return L1DepositERC721, nil
	case abi.L1DepositERC1155Sig.String():
		L1DepositERC1155, err := bridge.L1DepositERC1155(event)
		if err != nil {
			return nil, err
		}
		return L1DepositERC1155, nil
	case abi.L1BatchDepositERC721Sig.String():
		batchDepositERC721, err := bridge.L1BatchDepositERC721(event)
		if err != nil {
			return nil, err
		}
		return batchDepositERC721, nil
	case abi.L1BatchDepositERC1155Sig.String():
		L1BatchDepositERC1155, err := bridge.L1BatchDepositERC1155(event)
		if err != nil {
			return nil, err
		}
		return L1BatchDepositERC1155, nil
	case abi.L1SentMessageEventSignature.String():
		_, err := bridge.L1SentMessageEvent(chainIdStr, event, sp.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case abi.L1RelayedMessageEventSignature.String():
		_, err := bridge.L1RelayedMessageEvent(chainIdStr, event, sp.db)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (sp *ScrollEventProcessor) l2EventUnpack(event event.ContractEvent) (*worker.L2ToL1, error) {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L2WithdrawETHSig.String():
		withdrawETH, err := bridge.L2WithdrawETH(event)
		if err != nil {
			return nil, err
		}
		return withdrawETH, nil
	case abi.L2WithdrawERC20Sig.String():
		L2WithdrawERC20, err := bridge.L2WithdrawERC20(event)
		if err != nil {
			return nil, err
		}
		return L2WithdrawERC20, nil
	case abi.L2WithdrawERC721Sig.String():
		L2WithdrawERC721, err := bridge.L2WithdrawERC721(event)
		if err != nil {
			return nil, err
		}
		return L2WithdrawERC721, nil
	case abi.L2WithdrawERC1155Sig.String():
		L2WithdrawERC1155, err := bridge.L2WithdrawERC1155(event)
		if err != nil {
			return nil, err
		}
		return L2WithdrawERC1155, nil
	case abi.L2BatchWithdrawERC721Sig.String():
		L2BatchWithdrawERC721, err := bridge.L2BatchWithdrawERC721(event)
		if err != nil {
			return nil, err
		}
		return L2BatchWithdrawERC721, nil
	case abi.L2BatchWithdrawERC1155Sig.String():
		L2BatchWithdrawERC1155, err := bridge.L2BatchWithdrawERC1155(event)
		if err != nil {
			return nil, err
		}
		return L2BatchWithdrawERC1155, nil
	case abi.L2SentMessageEventSignature.String():
		_, err := bridge.L2SentMessageEvent(chainIdStr, event, sp.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case abi.L2RelayedMessageEventSignature.String():
		_, err := bridge.L2RelayedMessageEvent(chainIdStr, event, sp.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, nil
}
