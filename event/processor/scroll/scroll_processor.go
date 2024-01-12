package scroll

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/bridge"
	"math/big"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/event/processor"
)

type ScrollProcessor struct {
	processor.IProcessor
}

func NewBridgeProcessor(log log.Logger, db *database.DB,
	cfg *config.Config, shutdown context.CancelCauseFunc) (*processor.IProcessor, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &processor.IProcessor{
		Log:            log,
		Db:             db,
		Cfg:            cfg,
		ResourceCtx:    resCtx,
		ResourceCancel: resCancel,
		Tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (sp *ScrollProcessor) Start() error {
	sp.Log.Info("starting bridge processor...")
	sp.Tasks.Go(func() error {
		sp.onL1Data()
		sp.Log.Info("no more l1 etl updates. shutting down l1 task")
		return nil
	})
	// start L2 worker
	sp.Tasks.Go(func() error {

		sp.onL2Data()

		sp.Log.Info("no more l2 etl updates. shutting down l2 task")
		return nil
	})
	return nil
}

func (sp *ScrollProcessor) Close() error {
	sp.ResourceCancel()
	return sp.Tasks.Wait()
}

func (sp *ScrollProcessor) onL1Data() error {
	//fromL1Height := new(big.Int).Add(sp.LastFinalizedL1Header.Number, bigint.One)
	//toL1Height := new(big.Int).Add(fromL1Height, big.NewInt(10000))
	//if err := sp.Db.Transaction(func(tx *database.DB) error {
	//	l1EventsFetchErr := sp.l1EventsFetch(fromL1Height, toL1Height)
	//	if l1EventsFetchErr != nil {
	//		return l1EventsFetchErr
	//	}
	//	return nil
	//}); err != nil {
	//	return err
	//}
	return nil
}

func (sp *ScrollProcessor) onL2Data() error {
	//fromL2Height := new(big.Int).Add(sp.LastFinalizedL2Header.Number, bigint.One)
	//toL2Height := new(big.Int).Add(fromL2Height, big.NewInt(10000))
	//if err := sp.Db.Transaction(func(tx *database.DB) error {
	//	l2EventsFetchErr := sp.l2EventsFetch(fromL2Height, toL2Height)
	//	if l2EventsFetchErr != nil {
	//		return l2EventsFetchErr
	//	}
	//	return nil
	//}); err != nil {
	//	return err
	//}
	return nil
}

func (sp *ScrollProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	//l1Contracts := sp.ChainConfig.L1Contracts
	//batchLog := sp.Log.New("start_number", fromL1Height, "end_number", toL1Height)
	//for _, l1contract := range l1Contracts {
	//	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l1contract)}
	//	events, err := sp.Db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
	//	if err != nil {
	//		batchLog.Error("failed to index L1ContractEventsWithFilter ", "err", err)
	//		return err
	//	}
	//	l1ToL2s := make([]worker.L1ToL2, len(events))
	//	for _, contractEvent := range events {
	//		unpackEvent, unpackErr := sp.l1EventUnpack(contractEvent)
	//		if unpackErr != nil {
	//			batchLog.Error("failed to index L1 bridge events", "unpackErr", unpackErr)
	//			return unpackErr
	//		}
	//		if unpackEvent != nil {
	//			l1ToL2s = append(l1ToL2s, *unpackEvent)
	//		}
	//	}
	//	if len(l1ToL2s) > 0 {
	//		saveErr := sp.Db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s)
	//		if saveErr != nil {
	//			batchLog.Error("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
	//			return saveErr
	//		}
	//	}
	//}
	return nil
}

func (sp *ScrollProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	//l2Contracts := sp.ChainConfig.L2Contracts
	//batchLog := sp.Log.New("start_number", fromL1Height, "end_number", toL1Height)
	//for _, l2contract := range l2Contracts {
	//	contractEventFilter := event.ContractEvent{ContractAddress: common.HexToAddress(l2contract)}
	//	events, err := sp.Db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
	//	if err != nil {
	//		batchLog.Error("failed to index L2ContractEventsWithFilter ", "err", err)
	//		return err
	//	}
	//	l2ToL1s := make([]worker.L2ToL1, len(events))
	//	for _, contractEvent := range events {
	//		unpackEvent, unpackErr := sp.l2EventUnpack(contractEvent)
	//		if unpackErr != nil {
	//			batchLog.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
	//			return unpackErr
	//		}
	//		if unpackEvent != nil {
	//			l2ToL1s = append(l2ToL1s, *unpackEvent)
	//		}
	//	}
	//	if len(l2ToL1s) > 0 {
	//		saveErr := sp.Db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s)
	//		if saveErr != nil {
	//			batchLog.Error("failed to StoreL2ToL1Transactions", "saveErr", saveErr)
	//			return saveErr
	//		}
	//	}
	//}
	return nil
}

func (sp *ScrollProcessor) l1EventUnpack(event event.ContractEvent) (*worker.L1ToL2, error) {

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
		//case abi.L1SentMessageEventSignature.String():
		//	_, err := bridge.L1SentMessageEvent(event, sp.Db)
		//	if err != nil {
		//		return nil, err
		//	}
		//	return nil, nil
		//case abi.L1RelayedMessageEventSignature.String():
		//	_, err := bridge.L1RelayedMessageEvent(event, sp.Db)
		//	if err != nil {
		//		return nil, err
		//	}
	}
	return nil, nil
}

func (sp *ScrollProcessor) l2EventUnpack(event event.ContractEvent) (*worker.L2ToL1, error) {
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
		//case abi.L2SentMessageEventSignature.String():
		//	_, err := bridge.L2SentMessageEvent(event, sp.Db)
		//	if err != nil {
		//		return nil, err
		//	}
		//	return nil, nil
		//case abi.L2RelayedMessageEventSignature.String():
		//	_, err := bridge.L2RelayedMessageEvent(event, sp.Db)
		//	if err != nil {
		//		return nil, err
		//	}
		//	return nil, nil
	}
	return nil, nil
}
