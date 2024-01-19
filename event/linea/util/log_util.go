package util

import (
	"encoding/json"
	"fmt"
	"strings"

	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func DecodeLog(contactAbi *eth_abi.ABI, eventName string, rlpLog types.Log) (map[string]interface{}, error) {
	eventInfo := make(map[string]interface{})
	// unpack event's data
	unpackErr := contactAbi.UnpackIntoMap(eventInfo, eventName, rlpLog.Data)
	if unpackErr != nil {
		log.Warn("Failed to unpack SentMessage event", "err", unpackErr)
		return nil, unpackErr
	}
	// unpack topic's data
	chainTopics := rlpLog.Topics
	if !contactAbi.Events[eventName].Anonymous {
		chainTopics = rlpLog.Topics
	}
	var topicInputs eth_abi.Arguments
	inputs := contactAbi.Events[eventName].Inputs
	for _, input := range inputs {
		if input.Indexed {
			topicInputs = append(topicInputs, input)
		}
	}
	if len(topicInputs) > 0 {
		eth_abi.ParseTopicsIntoMap(eventInfo, topicInputs, chainTopics)
	}
	fmt.Println(eventInfo)

	//marshal, _ := json.Marshal(eventInfo)
	//fmt.Println(string(marshal))
	calldataBytes := eventInfo["_calldata"].([]byte)

	//calldataStr := "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006eb4c000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c4cc29a30600000000000000000000000036c56fa4327122f7ea3683b2d527c1b2476d32ee000000000000000000000000000000000000000000000000002dc7f9e1bca4000000000000000000000000000000000000000000000000000028f723671b516b00000000000000000000000000000000000000000000000000000000657521e7000000000000000000000000710bda329b2a6224e4b44833de30f38e7f81d56400000000000000000000000000000000000000000000000000049e57d635400000000000000000000000000000000000000000000000000000000000"
	//decodedCalldata, _ := hex.DecodeString(calldataStr)

	if len(calldataBytes) > 0 {
		calldataBytes = calldataBytes[4:]
		calldata, calldataUnpackErr := DecodeMessageCallData(calldataBytes)
		if calldataUnpackErr != nil {
			log.Warn("Failed to unpack DecodeMessageCallData function", "err", calldataUnpackErr)
			return nil, calldataUnpackErr
		}
		eventInfo["calldata"] = calldata
	}
	return eventInfo, nil
}

func DecodeMessageCallData(data []byte) (map[string]interface{}, error) {
	//functionSignature := "completeBridging(address,uint256,address,uint256,bytes)"
	callDataAbi := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_tokenMetadata\",\"type\":\"bytes\"}],\"name\":\"completeBridging\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	thisAbi, _ := eth_abi.JSON(strings.NewReader(callDataAbi))
	calldataInfo := make(map[string]interface{})
	// 解析 calldata 数据

	unpackErr := thisAbi.UnpackIntoMap(calldataInfo, "completeBridging", data)
	//unpackErr := thisAbi.Methods[functionSignature].Inputs.UnpackIntoMap(calldataInfo, data)
	//_, unpackErr := thisAbi.Unpack("completeBridging", data)
	if unpackErr != nil {
		fmt.Println(unpackErr)
		log.Warn("Failed to unpack DecodeMessageCallData function", "err", unpackErr)
		return nil, unpackErr
	}
	fmt.Println(calldataInfo)
	_tokenMetadata := calldataInfo["_tokenMetadata"].([]byte)
	if len(_tokenMetadata) > 0 {
		//tokenMetadata, tokenMetadataUnpackErr := DecodeTokenMetadata(_tokenMetadata)
		//if tokenMetadataUnpackErr != nil {
		//	log.Warn("Failed to unpack DecodeMessageCallData function", "err", tokenMetadataUnpackErr)
		//	return nil, tokenMetadataUnpackErr
		//}
		//calldataInfo["tokenMetadata"] = tokenMetadata
	}
	return calldataInfo, nil
}

func DecodeTokenMetadata(data []byte) (map[string]interface{}, error) {
	callDataAbi := "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"safeName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"safeSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"safeDecimals\",\"type\":\"uint8\"}],\"name\":\"tokenMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	thisAbi, _ := eth_abi.JSON(strings.NewReader(callDataAbi))
	info := make(map[string]interface{})
	unpackErr := thisAbi.UnpackIntoMap(info, "tokenMetadata", data)
	if unpackErr != nil {
		log.Warn("Failed to unpack DecodeTokenMetadata function", "err", unpackErr)
		return nil, unpackErr
	}
	marshal, _ := json.Marshal(info)
	fmt.Println(string(marshal))
	return info, nil
}
