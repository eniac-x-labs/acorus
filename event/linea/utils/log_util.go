package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"strings"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeLog(contactAbi, callDataAbi *eth_abi.ABI, eventName, calldataFuncName string, rlpLog types.Log) (map[string]interface{}, error) {
	eventInfo := make(map[string]interface{})
	// unpack event's data
	unpackErr := contactAbi.UnpackIntoMap(eventInfo, eventName, rlpLog.Data)
	if unpackErr != nil {
		log.Error("Failed to unpack SentMessage event", "err", unpackErr)
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
			log.Error("ParseTopicsIntoMap linea", "err", unpackErr)
			return nil, unpackErr
		}
	}

	if calldataBytes, ok := eventInfo["_calldata"].([]byte); ok {
		if len(calldataBytes) > 0 {
			bytes := calldataBytes[:4]
			fmt.Println(hex.EncodeToString(bytes))
			calldataBytes = calldataBytes[4:]
			calldata, calldataUnpackErr := DecodeMessageCallData(callDataAbi, calldataFuncName, calldataBytes)
			if calldataUnpackErr != nil {
				log.Error("Failed to unpack DecodeMessageCallData function", "err", calldataUnpackErr)
				return nil, calldataUnpackErr
			}
			delete(eventInfo, "_calldata")
			eventInfo["calldata"] = calldata
		}
	}
	return eventInfo, nil
}

func DecodeMessageCallData(callDataAbi *eth_abi.ABI, calldataFuncName string, data []byte) (map[string]interface{}, error) {
	calldataInfo := make(map[string]interface{})
	method := callDataAbi.Methods[calldataFuncName]
	args, calldataUnpackErr := method.Inputs.UnpackValues(data)
	if calldataUnpackErr != nil {
		return nil, calldataUnpackErr
	}
	for i, arg := range method.Inputs {
		calldataInfo[arg.Name] = args[i]
	}
	if _tokenMetadata, ok := calldataInfo["_tokenMetadata"].([]byte); ok {
		if len(_tokenMetadata) > 0 {
			tokenMetadata, tokenMetadataUnpackErr := DecodeTokenMetadata(_tokenMetadata)
			if tokenMetadataUnpackErr != nil {
				log.Error("Failed to unpack DecodeMessageCallData function", "err", tokenMetadataUnpackErr)
				return nil, tokenMetadataUnpackErr
			}
			calldataInfo["tokenMetadata"] = tokenMetadata
			delete(calldataInfo, "_tokenMetadata")
		}
	}
	return calldataInfo, nil
}

func DecodeTokenMetadata(data []byte) (map[string]interface{}, error) {
	callDataAbi := "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"safeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"safeSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"safeDecimals\",\"type\":\"uint8\"}],\"name\":\"tokenMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	thisAbi, _ := eth_abi.JSON(strings.NewReader(callDataAbi))
	info := make(map[string]interface{})
	method := thisAbi.Methods["tokenMetadata"]
	args, tokenMetadataUnpackErr := method.Inputs.UnpackValues(data)
	if tokenMetadataUnpackErr != nil {
		return nil, tokenMetadataUnpackErr
	}
	for i, arg := range method.Inputs {
		info[arg.Name] = args[i]
	}
	return info, nil
}
