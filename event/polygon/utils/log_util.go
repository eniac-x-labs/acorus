package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"math/big"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeLog(contactAbi *eth_abi.ABI, eventName string, rlpLog types.Log) (map[string]interface{}, error) {
	eventInfo := make(map[string]interface{})
	// unpack event's data
	unpackErr := contactAbi.UnpackIntoMap(eventInfo, eventName, rlpLog.Data)
	if unpackErr != nil {
		log.Error("Failed to unpack ClaimEvent", "err", unpackErr)
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
			log.Error("ParseTopicsIntoMap unpack err ", "err", unpackErr)
			return nil, unpackErr
		}
	}
	return eventInfo, nil
}

func DecodeGlobalIndex(globalIndex *big.Int) (bool, *big.Int, *big.Int, error) {
	const lengthGlobalIndexInBytes = 32
	var buf [32]byte
	gIBytes := globalIndex.FillBytes(buf[:])
	if len(gIBytes) != lengthGlobalIndexInBytes {
		return false, big.NewInt(0), big.NewInt(0), fmt.Errorf("invalid globaIndex length. Should be 32. Current length: %d", len(gIBytes))
	}
	mainnetFlag := big.NewInt(0).SetBytes([]byte{gIBytes[23]}).Uint64() == 1
	rollupIndex := big.NewInt(0).SetBytes(gIBytes[24:28])
	localRootIndex := big.NewInt(0).SetBytes(gIBytes[29:32])
	return mainnetFlag, rollupIndex, localRootIndex, nil
}
