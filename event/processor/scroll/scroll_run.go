package scroll

import (
	"context"
	"github.com/cornerstone-labs/acorus/common/bigint"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processor"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processor/scroll/bridge"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type ScrollProcessor struct {
	processor.IEventUnpack
	log         log.Logger
	db          *database.DB
	l1Etl       *synchronizer.L1Sync
	chainConfig config.ChainConfig
}

// GetChainUnpackService get service implement
func (s *ScrollProcessor) GetChainUnpackService(eventParam processor.GetEventFactoryParams) processor.IEventUnpack {
	return &ScrollProcessor{
		log:         eventParam.Log,
		db:          eventParam.Db,
		l1Etl:       eventParam.L1Syncer,
		chainConfig: eventParam.ChainConfig,
	}
}

// todo should is common func
func (s *ScrollProcessor) GetLatestHeader(chainId uint64) (*types.Header, *types.Header, error) {
	latestL1Header, err := s.db.ContractEvents.L1LatestBlockHeader(chainId)
	if err != nil {
		return nil, nil, err
	}
	latestL2Header, err := s.db.ContractEvents.L2LatestBlockHeader(chainId)
	if err != nil {
		return nil, nil, err
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
	return l1Header, l2Header, nil
}

func (s *ScrollProcessor) Start(ctx context.Context) error {
	done := ctx.Done()
	l1EtlUpdates := s.l1Etl.Notify()
	startup := make(chan interface{}, 1)
	startup <- nil
	s.log.Info("starting scroll bridge processor...")
	for {
		select {
		case <-done:
			s.log.Info("stopping scroll bridge processor")
			return nil

		// Tickers
		case <-startup:
		case <-l1EtlUpdates:
		}
		s.Run()
	}
}

func (s *ScrollProcessor) Run() error {
	maxEpochRange := uint64(10_000)
	l1Header, l2Header, err := s.GetLatestHeader(s.chainConfig.ChainId)
	if err != nil {
		return err
	}

	genesisL1Height := big.NewInt(int64(s.chainConfig.L2StartHeight))
	maxEpochRangeBigInt := big.NewInt(int64(maxEpochRange))

	fromL1Height, fromL2Height := genesisL1Height, bigint.Zero
	if l1Header != nil {
		fromL1Height = new(big.Int).Add(l1Header.Number, bigint.One)
	}
	if l2Header != nil {
		fromL2Height = new(big.Int).Add(l2Header.Number, bigint.One)

	}
	toL1Height, toL2Height := new(big.Int).Add(fromL1Height, maxEpochRangeBigInt), new(big.Int).Add(fromL2Height, maxEpochRangeBigInt)

	if err := s.db.Transaction(func(tx *database.DB) error {
		l1EventsFetchErr := s.l1EventsFetch(fromL1Height, toL1Height)
		if l1EventsFetchErr != nil {
			return l1EventsFetchErr
		}
		l2EventsFetchErr := s.l2EventsFetch(fromL2Height, toL2Height)
		if l2EventsFetchErr != nil {
			return l2EventsFetchErr
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *ScrollProcessor) l1EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l1Contracts := s.chainConfig.L1Contracts
	batchLog := s.log.New("start_number", fromL1Height, "end_number", toL1Height)
	for _, l1contract := range l1Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: l1contract}
		events, err := s.db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			batchLog.Error("failed to index L1ContractEventsWithFilter ", "err", err)
			return err
		}
		l1ToL2s := make([]*worker.L1ToL2, len(events))
		for _, contractEvent := range events {
			unpackEvent, unpackErr := s.l1EventUnpack(contractEvent)
			if unpackErr != nil {
				batchLog.Error("failed to index L1 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l1ToL2s = append(l1ToL2s, unpackEvent)
			}
		}
		if len(l1ToL2s) > 0 {
			saveErr := s.db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s)
			if saveErr != nil {
				batchLog.Error("failed to StoreL1ToL2Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (s *ScrollProcessor) l2EventsFetch(fromL1Height, toL1Height *big.Int) error {
	l2Contracts := s.chainConfig.L2Contracts
	batchLog := s.log.New("start_number", fromL1Height, "end_number", toL1Height)
	for _, l2contract := range l2Contracts {
		contractEventFilter := event.ContractEvent{ContractAddress: l2contract}
		events, err := s.db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromL1Height, toL1Height)
		if err != nil {
			batchLog.Error("failed to index L2ContractEventsWithFilter ", "err", err)
			return err
		}
		l2ToL1s := make([]*worker.L2ToL1, len(events))
		for _, contractEvent := range events {
			unpackEvent, unpackErr := s.l2EventUnpack(contractEvent)
			if unpackErr != nil {
				batchLog.Error("failed to index L2 bridge events", "unpackErr", unpackErr)
				return unpackErr
			}
			if unpackEvent != nil {
				l2ToL1s = append(l2ToL1s, unpackEvent)
			}
		}
		if len(l2ToL1s) > 0 {
			saveErr := s.db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s)
			if saveErr != nil {
				batchLog.Error("failed to StoreL2ToL1Transactions", "saveErr", saveErr)
				return saveErr
			}
		}
	}
	return nil
}

func (s *ScrollProcessor) l1EventUnpack(event event.L1ContractEvent) (*worker.L1ToL2, error) {

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
		_, err := bridge.L1SentMessageEvent(event, s.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case abi.L1RelayedMessageEventSignature.String():
		_, err := bridge.L1RelayedMessageEvent(event, s.db)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (s *ScrollProcessor) l2EventUnpack(event event.L2ContractEvent) (*worker.L2ToL1, error) {
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
		_, err := bridge.L2SentMessageEvent(event, s.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	case abi.L2RelayedMessageEventSignature.String():
		_, err := bridge.L2RelayedMessageEvent(event, s.db)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	return nil, nil
}
