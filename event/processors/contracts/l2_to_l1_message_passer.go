package contracts

import (
	"github.com/cornerstone-labs/acorus/event/op-bindings/bindings"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/event/l1-l2"
)

type L2ToL1MessagePasserMessagePassed struct {
	Event          *event.ContractEvent
	WithdrawalHash common.Hash
	GasLimit       *big.Int
	Nonce          *big.Int
	Tx             l1_l2.Transaction
}

func L2ToL1MessagePasserMessagePassedEvents(contractAddress common.Address, db *database.DB, fromHeight, toHeight *big.Int) ([]L2ToL1MessagePasserMessagePassed, error) {
	l2ToL1MessagePasserAbi, err := bindings.L2ToL1MessagePasserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	messagePassedAbi := l2ToL1MessagePasserAbi.Events["MessagePassed"]
	contractEventFilter := event.ContractEvent{ContractAddress: contractAddress, EventSignature: messagePassedAbi.ID}
	messagePassedEvents, err := db.ContractEvents.L2ContractEventsWithFilter(contractEventFilter, fromHeight, toHeight)
	if err != nil {
		return nil, err
	}

	messagesPassed := make([]L2ToL1MessagePasserMessagePassed, len(messagePassedEvents))
	for i := range messagePassedEvents {
		messagePassed := bindings.L2ToL1MessagePasserMessagePassed{Raw: *messagePassedEvents[i].RLPLog}
		err := UnpackLog(&messagePassed, messagePassedEvents[i].RLPLog, messagePassedAbi.Name, l2ToL1MessagePasserAbi)
		if err != nil {
			return nil, err
		}

		messagesPassed[i] = L2ToL1MessagePasserMessagePassed{
			Event:          &messagePassedEvents[i].ContractEvent,
			WithdrawalHash: messagePassed.WithdrawalHash,
			Nonce:          messagePassed.Nonce,
			GasLimit:       messagePassed.GasLimit,
			Tx: l1_l2.Transaction{
				FromAddress: messagePassed.Sender,
				ToAddress:   messagePassed.Target,
				ETHAmount:   messagePassed.EthValue,
				Data:        messagePassed.Data,
				Timestamp:   messagePassedEvents[i].Timestamp,
				ERC20Amount: messagePassed.MntValue,
			},
		}
	}

	return messagesPassed, nil
}
