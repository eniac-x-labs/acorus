package bridge

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/processors/linea/util"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/event/processors/linea/abi"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/utils"
)

func L1SentMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1SentMessageSignature := abi.L1SentMessageSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1SentMessageSignature}
	sentMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	fmt.Println(len(sentMessageEvents))
	for i := range sentMessageEvents {
		eventTx := sentMessageEvents[i]
		rlpLog := *eventTx.RLPLog
		decodeLog, decodeErr := util.DecodeLog(abi.L1MessageEventABI, "MessageSent", rlpLog)
		if decodeErr != nil {
			log.Warn("Failed to unpack SentMessage event", "err", err)
		}
		fmt.Println(decodeLog)

	}
	//l1BridgeMsgs := make([]l1_l2.L1BridgeMessage, len(sentMessageEvents))
	//for
	return nil
}

func L1ClaimedMessageEvent(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) error {
	l1ClaimedMessageSignature := abi.L1ClaimedMessageSignature
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: l1ClaimedMessageSignature}
	claimedMessageEvents, err := db.ContractEvents.L1ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return err
	}
	for i := range claimedMessageEvents {
		depositTx := claimedMessageEvents[i]
		rlpLog := *depositTx.RLPLog
		claimedMessageEvent := abi.L1MessageClaimed{}
		unpackErr := utils.UnpackLog(abi.L1MessageEventABI, &claimedMessageEvent, "MessageClaimed", rlpLog)
		if unpackErr != nil {
			log.Warn("Failed to unpack SentMessage event", "err", err)
			return unpackErr
		}
		//claimedMessageEvent
	}
	return nil
}
