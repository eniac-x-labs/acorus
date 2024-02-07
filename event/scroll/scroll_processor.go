package scroll

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/common/global_const"
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
	tickerEventOn1 := time.NewTicker(sp.loopInterval)
	tickerEventOn2 := time.NewTicker(sp.loopInterval)
	tickerEventRel := time.NewTicker(sp.loopInterval)
	log.Println("starting scroll bridge processor...")
	sp.tasks.Go(func() error {
		for range tickerEventOn1.C {
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
		for range tickerEventOn2.C {
			err := sp.onL2Data()
			if err != nil {
				log.Println("no more l2 etl updates. shutting down l2 task")
				return err
			}
		}
		return nil
	})

	// start relation worker
	sp.tasks.Go(func() error {
		for range tickerEventRel.C {
			err := sp.relationL1L2()
			if err != nil {
				log.Println("shutting down relation task")
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
		lastBlockHeard, err := sp.db.L1ToL2.L1LatestBlockHeader(chainIdStr)
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

	chainLatestBlockHeader, err := sp.db.Blocks.ChainLatestBlockHeader(strconv.FormatUint(global_const.EthereumChainId, 10))
	if err != nil {
		sp.l1StartHeight = new(big.Int).Sub(sp.l1StartHeight, bigint.One)
		return err
	}
	if chainLatestBlockHeader == nil {
		sp.l1StartHeight = new(big.Int).Sub(sp.l1StartHeight, bigint.One)
		return nil
	}
	if chainLatestBlockHeader.Number.Cmp(fromL1Height) == -1 {
		fromL1Height = chainLatestBlockHeader.Number
		toL1Height = chainLatestBlockHeader.Number
	} else {
		if chainLatestBlockHeader.Number.Cmp(toL1Height) == -1 {
			toL1Height = chainLatestBlockHeader.Number
		}
	}
	if err := sp.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := sp.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		return nil
	}); err != nil {
		sp.l1StartHeight = new(big.Int).Sub(sp.l1StartHeight, bigint.One)

		return err
	}
	sp.l1StartHeight = toL1Height
	return nil
}

func (sp *ScrollEventProcessor) onL2Data() error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	if sp.l2StartHeight == nil {
		lastBlockHeard, err := sp.db.L2ToL1.L2LatestBlockHeader(chainIdStr)
		if err != nil {
			log.Println("l2 failed to get last block heard", "err", err)
			return err
		}
		if lastBlockHeard == nil {
			if sp.cfgRpc.StartBlock > 0 {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(int64(sp.cfgRpc.StartBlock)),
				}
			} else {
				lastBlockHeard = &common2.BlockHeader{
					Number: big.NewInt(0),
				}
			}
		}
		sp.l2StartHeight = lastBlockHeard.Number
	} else {
		sp.l2StartHeight = new(big.Int).Add(sp.l2StartHeight, bigint.One)
	}
	fromL2Height := sp.l2StartHeight
	toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(int64(sp.epoch)))
	chainLatestBlockHeader, err := sp.db.Blocks.ChainLatestBlockHeader(chainIdStr)
	if err != nil {
		sp.l2StartHeight = new(big.Int).Sub(sp.l2StartHeight, bigint.One)
		return err
	}
	if chainLatestBlockHeader == nil {
		sp.l2StartHeight = new(big.Int).Sub(sp.l2StartHeight, bigint.One)

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
	if err := sp.db.Transaction(func(tx *database.DB) error {
		l2EventsFetchErr := sp.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		sp.l2StartHeight = new(big.Int).Sub(sp.l2StartHeight, bigint.One)
		return err
	}
	sp.l2StartHeight = toL2Height
	return nil
}

func (sp *ScrollEventProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l1Contracts := sp.cfgRpc.Contracts
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
		events, err := sp.db.ContractEvents.ContractEventsWithFilter("1", contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			log.Println("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := sp.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
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
		for _, contractEvent := range events {
			unpackErr := sp.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				log.Println("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (sp *ScrollEventProcessor) l1EventUnpack(event event.ContractEvent) error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L1DepositETHSig.String():
		err := bridge.L1DepositETH(chainIdStr, event, sp.db)
		return err
	case abi.L1DepositERC20Sig.String():
		err := bridge.L1DepositERC20(chainIdStr, event, sp.db)
		return err
	case abi.L1DepositERC721Sig.String():
		err := bridge.L1DepositERC721(chainIdStr, event, sp.db)
		return err
	case abi.L1DepositERC1155Sig.String():
		err := bridge.L1DepositERC1155(chainIdStr, event, sp.db)
		return err
	case abi.L1BatchDepositERC721Sig.String():
		err := bridge.L1BatchDepositERC721(chainIdStr, event, sp.db)
		return err
	case abi.L1BatchDepositERC1155Sig.String():
		err := bridge.L1BatchDepositERC1155(chainIdStr, event, sp.db)
		return err
	case abi.L1SentMessageEventSignature.String():
		err := bridge.L1SentMessageEvent(chainIdStr, event, sp.db)
		return err
	case abi.L1RelayedMessageEventSignature.String():
		err := bridge.L1RelayedMessageEvent(chainIdStr, event, sp.db)
		return err
	}
	return nil
}

func (sp *ScrollEventProcessor) l2EventUnpack(event event.ContractEvent) error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))
	switch event.EventSignature.String() {
	case abi.L2WithdrawETHSig.String():
		err := bridge.L2WithdrawETH(chainIdStr, event, sp.db)
		return err
	case abi.L2WithdrawERC20Sig.String():
		err := bridge.L2WithdrawERC20(chainIdStr, event, sp.db)
		return err
	case abi.L2WithdrawERC721Sig.String():
		err := bridge.L2WithdrawERC721(chainIdStr, event, sp.db)
		return err
	case abi.L2WithdrawERC1155Sig.String():
		err := bridge.L2WithdrawERC1155(chainIdStr, event, sp.db)
		return err
	case abi.L2BatchWithdrawERC721Sig.String():
		err := bridge.L2BatchWithdrawERC721(chainIdStr, event, sp.db)
		return err
	case abi.L2BatchWithdrawERC1155Sig.String():
		err := bridge.L2BatchWithdrawERC1155(chainIdStr, event, sp.db)
		return err
	case abi.L2SentMessageEventSignature.String():
		err := bridge.L2SentMessageEvent(chainIdStr, event, sp.db)
		return err
	case abi.L2RelayedMessageEventSignature.String():
		err := bridge.L2RelayedMessageEvent(chainIdStr, event, sp.db)
		return err
	}
	return nil
}

func (sp *ScrollEventProcessor) relationL1L2() error {
	chainId := sp.cfgRpc.ChainId
	chainIdStr := strconv.Itoa(int(chainId))

	if err := sp.db.Transaction(func(tx *database.DB) error {
		// step 1
		if err := sp.db.MsgSentRelation.MsgHashRelation(chainIdStr); err != nil {
			return err
		}
		// step 2
		if err := sp.db.MsgSentRelation.RelayRelation(chainIdStr); err != nil {
			return err
		}
		// step 3
		if canSaveDataList, err := sp.db.MsgSentRelation.GetCanSaveDataList(chainIdStr); err != nil {
			return err
		} else {
			l1ToL2s := make([]worker.L1ToL2, 0)
			l2ToL1s := make([]worker.L2ToL1, 0)
			for _, data := range canSaveDataList {
				l1l2Data := data.Data
				if data.LayerType == global_const.LayerTypeOne {
					var l1Tol2 worker.L1ToL2
					if unMarErr := json.Unmarshal([]byte(l1l2Data), &l1Tol2); unMarErr != nil {
						return unMarErr
					}
					l1Tol2.MessageHash = data.MsgHash
					l1Tol2.L2BlockNumber = data.LayerBlockNumber
					l1Tol2.L2TransactionHash = data.LayerHash
					l1Tol2.Status = 1
					l1ToL2s = append(l1ToL2s, l1Tol2)
				}
				if data.LayerType == global_const.LayerTypeTwo {
					var l2Tol1 worker.L2ToL1
					if unMarErr := json.Unmarshal([]byte(l1l2Data), &l2Tol1); unMarErr != nil {
						return unMarErr
					}
					l2Tol1.MessageHash = data.MsgHash
					l2Tol1.L1BlockNumber = data.LayerBlockNumber
					l2Tol1.L1FinalizeTxHash = data.LayerHash
					l2Tol1.Status = 1
					l2ToL1s = append(l2ToL1s, l2Tol1)
				}

			}

			if len(l1ToL2s) > 0 {
				saveErr := sp.db.L1ToL2.StoreL1ToL2Transactions(chainIdStr, l1ToL2s)
				if saveErr != nil {
					log.Println("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
					return saveErr
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
		if err := sp.db.MsgSentRelation.L1RelationClear(chainIdStr); err != nil {
			return err
		}
		if err := sp.db.MsgSentRelation.L2RelationClear(chainIdStr); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
