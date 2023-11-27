package bridge

import (
	l1_l2 "github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/abi"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/utils"
)

func L1DepositETH(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositETHSig := abi.L1DepositETHSig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositETHSig}
	depositEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositEvents))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(depositEvents))
	for i := range depositEvents {
		ethDepositTx := depositEvents[i]
		rlpLog := *depositEvents[i].RLPLog
		depositEvent := abi.DepositETH{}

		unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositEvent, "DepositETH", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}

		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           ethDepositTx.TransactionHash,
			InitiatedL1EventGUID: ethDepositTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: depositEvent.From,
				ToAddress:   depositEvent.To,
				ETHAmount:   depositEvent.Amount,
				ERC20Amount: nil,
				Data:        depositEvent.Data,
				Timestamp:   ethDepositTx.Timestamp,
			},
		}
		l1ToL2s[i] = worker.L1ToL2{
			L1BlockNumber:     rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositEvent.From,
			ToAddress:         depositEvent.To,
			L1TokenAddress:    common.Address{},
			L2TokenAddress:    common.Address{},
			ETHAmount:         depositEvent.Amount,
			ERC20Amount:       nil,
			GasLimit:          nil,
			Timestamp:         int64(ethDepositTx.Timestamp),
			AssetType:         int64(common2.ETH),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC20(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC20Sig := abi.L1DepositERC20Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC20Sig}
	depositErc20Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc20Events))
	transactionDeposits := make([]l1_l2.L1TransactionDeposit, len(depositErc20Events))

	for i := range depositErc20Events {
		depositTx := depositErc20Events[i]
		rlpLog := *depositErc20Events[i].RLPLog
		depositErc20Event := abi.ERC20MessageEvent{}
		transactionDeposits[i] = l1_l2.L1TransactionDeposit{
			SourceHash:           depositTx.TransactionHash,
			InitiatedL1EventGUID: depositTx.GUID,
			Tx: l1_l2.Transaction{
				FromAddress: depositErc20Event.From,
				ToAddress:   depositErc20Event.To,
				ETHAmount:   nil,
				ERC20Amount: depositErc20Event.Amount,
				Data:        depositErc20Event.Data,
				Timestamp:   depositTx.Timestamp,
			},
		}
		unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositErc20Event, "DepositERC20", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			L1BlockNumber:     rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc20Event.From,
			ToAddress:         depositErc20Event.To,
			L1TokenAddress:    depositErc20Event.L1Token,
			L2TokenAddress:    depositErc20Event.L2Token,
			ETHAmount:         nil,
			ERC20Amount:       depositErc20Event.Amount,
			GasLimit:          nil,
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC20),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.BridgeTransactions.StoreL1TransactionDeposits(transactionDeposits); err != nil {
			return err
		}
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC721(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC721Sig := abi.L1DepositERC721Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC721Sig}
	depositErc721Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc721Events))
	for i := range depositErc721Events {
		depositTx := depositErc721Events[i]
		rlpLog := *depositErc721Events[i].RLPLog
		depositErc721Event := abi.ERC721MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositErc721Event, "DepositERC721", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			L1BlockNumber:     rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc721Event.From,
			ToAddress:         depositErc721Event.To,
			L1TokenAddress:    depositErc721Event.L1Token,
			L2TokenAddress:    depositErc721Event.L2Token,
			ETHAmount:         depositErc721Event.Amount,
			ERC20Amount:       nil,
			GasLimit:          nil,
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC721),
			TokenIds:          depositErc721Event.TokenID.String(),
		}
	}
	if len(l1ToL2s) > 0 {

		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil
}

func L1DepositERC1155(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1DepositERC1155Sig := abi.L1DepositERC1155Sig
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1DepositERC1155Sig}
	depositErc1155Events, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1ToL2s := make([]worker.L1ToL2, len(depositErc1155Events))
	for i := range depositErc1155Events {
		depositTx := depositErc1155Events[i]
		rlpLog := *depositErc1155Events[i].RLPLog
		depositErc1155Event := abi.ERC1155MessageEvent{}

		unpackErr := utils.UnpackLog(abi.L1ETHGatewayABI, &depositErc1155Event, "DepositERC1155", rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		l1ToL2s[i] = worker.L1ToL2{
			L1BlockNumber:     rlpLog.BlockHash,
			QueueIndex:        nil,
			L1TransactionHash: rlpLog.TxHash,
			L2TransactionHash: common.Hash{},
			L1TxOrigin:        common.Hash{},
			Status:            0,
			FromAddress:       depositErc1155Event.From,
			ToAddress:         depositErc1155Event.To,
			L1TokenAddress:    depositErc1155Event.L1Token,
			L2TokenAddress:    depositErc1155Event.L2Token,
			ETHAmount:         depositErc1155Event.Amount,
			ERC20Amount:       nil,
			GasLimit:          nil,
			Timestamp:         int64(depositTx.Timestamp),
			AssetType:         int64(common2.ERC1155),
			TokenIds:          depositErc1155Event.TokenID.String(),
		}
	}
	if len(l1ToL2s) > 0 {
		if err := db.L1ToL2.StoreL1ToL2Transactions(l1ToL2s); err != nil {
			return err
		}
	}
	return nil

}

func L1SentMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1SentMessageEventSignature := abi.L1SentMessageEventSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1SentMessageEventSignature}
	sentMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	l1BridgeMsgs := make([]l1_l2.L1BridgeMessage, len(sentMessageEvents))
	for i := range sentMessageEvents {
		depositTx := sentMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		sentMessageEvent := abi.L1SentMessageEvent{}
		unpackErr := utils.UnpackLog(abi.L1ScrollMessengerABI, &sentMessageEvent, "SentMessage", rlpLog)
		if unpackErr != nil {
			log.Warn("Failed to unpack SentMessage event", "err", err)
			return unpackErr
		}
		msgHash := ComputeMessageHash(sentMessageEvent.Sender, sentMessageEvent.Target,
			sentMessageEvent.Value, sentMessageEvent.MessageNonce, sentMessageEvent.Message)
		l1BridgeMsgs[i] = l1_l2.L1BridgeMessage{
			BridgeMessage: l1_l2.BridgeMessage{
				MessageHash: msgHash,
				Nonce:       sentMessageEvent.MessageNonce,
				GasLimit:    sentMessageEvent.GasLimit,
				Tx: l1_l2.Transaction{
					FromAddress: sentMessageEvent.Sender,
					ToAddress:   sentMessageEvent.Target,
					ETHAmount:   sentMessageEvent.Value,
					ERC20Amount: nil,
					Data:        sentMessageEvent.Message,
					Timestamp:   depositTx.Timestamp,
				},
			},
			TransactionSourceHash: rlpLog.TxHash,
		}
	}
	if len(l1BridgeMsgs) > 0 {
		if err := db.BridgeMessages.StoreL1BridgeMessages(l1BridgeMsgs); err != nil {
			return err
		}
	}
	return nil
}

func L1BatchDepositERC721() {

}

func L1BatchDepositERC1155() {

}

func L1RelayedMessageEvent() {

}

func ComputeMessageHash(
	sender common.Address,
	target common.Address,
	value *big.Int,
	messageNonce *big.Int,
	message []byte,
) common.Hash {
	data, _ := abi.L2ScrollMessengerABI.Pack("relayMessage", sender, target, value, messageNonce, message)
	return common.BytesToHash(crypto.Keccak256(data))
}
