package utils

import (
	"log"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeLog(contactAbi *eth_abi.ABI, eventName string, rlpLog types.Log) (map[string]interface{}, error) {
	eventInfo := make(map[string]interface{})
	// unpack event's data
	unpackErr := contactAbi.UnpackIntoMap(eventInfo, eventName, rlpLog.Data)
	if unpackErr != nil {
		log.Println("Failed to unpack ClaimEvent", "err", unpackErr)
		return nil, unpackErr
	}
	// unpack topic's data
	chainTopics := rlpLog.Topics
	if !contactAbi.Events[eventName].Anonymous {
		chainTopics = rlpLog.Topics[1:]
	}
	var topicInputs eth_abi.Arguments
	inputs := contactAbi.Events[eventName].Inputs
	for _, input := range inputs {
		if input.Indexed {
			topicInputs = append(topicInputs, input)
		}
	}
	if len(topicInputs) > 0 {
		unpackErr := eth_abi.ParseTopicsIntoMap(eventInfo, topicInputs, chainTopics)
		if unpackErr != nil {
			log.Println(unpackErr)
			return nil, unpackErr
		}
	}
	return eventInfo, nil
}
