package utils

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	eth_abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/cornerstone-labs/acorus/event/linea/abi"
)

func TestGetAbi_LogUnpack(t *testing.T) {
	query := ethereum.FilterQuery{}
	query.FromBlock = big.NewInt(18702886)
	query.ToBlock = big.NewInt(18702886)
	//query.Addresses = []common.Address{common.HexToAddress("0x98d8def8aba6e5f3c9ef32363a68428c2834b499")}
	query.Topics = [][]common.Hash{{abi.L1SentMessageSignature}}
	client, _ := ethclient.Dial("https://rpc.ankr.com/eth/7f3ae11204e03961b67c557c4996244f0a53222b23c31a7baf9ae91c6bd89702")

	logs, _ := client.FilterLogs(context.Background(), query)
	for _, v := range logs {

		marshal, _ := json.Marshal(v)
		hash := v.TxHash.Hex()
		if hash == "0x21f956ad759a003e5dba63a4923175ae334db3c4cc3fc2a0f7ffc43e0d321332" {
			fmt.Println(hash)
			fmt.Println(string(marshal))
			log, _ := DecodeLog(abi.L1MessageEventABI, "MessageSent", v)
			fmt.Println(log)
			//bytes := log["_calldata"].([]byte)
			//fmt.Println(string(bytes))
		}
	}
}

func TestBase64(t *testing.T) {
	// 要解码的Base64字符串
	base64String := "zCmjBgAAAAAAAAAAAAAAADbFb6QycSL36jaDstUnwbJHbTLuAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC3H+eG8pAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKPcjZxtRawAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABldSHnAAAAAAAAAAAAAAAAcQvaMpsqYiTktEgz3jDzjn+B1WQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJ5X1jVAAA=="

	// 对Base64字符串进行解码
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal(err)
	}

	// 将解码后的字节转换为字符串
	decodedString := string(decodedBytes)

	// 输出解码后的字符串
	fmt.Println(decodedString)

}

func TestAbi(t *testing.T) {
	// Solidity 合约中的函数签名
	//functionSignature := "completeBridging(address,uint256,address,uint256,bytes)"

	// 要解析的 ABI 编码的 calldata
	//cc29a306
	abiEncodedCalldataHex := "cc29a30600000000000000000000000036c56fa4327122f7ea3683b2d527c1b2476d32ee000000000000000000000000000000000000000000000000002dc7f9e1bca4000000000000000000000000000000000000000000000000000028f723671b516b00000000000000000000000000000000000000000000000000000000657521e7000000000000000000000000710bda329b2a6224e4b44833de30f38e7f81d56400000000000000000000000000000000000000000000000000049e57d6354000"
	method, _ := eth_abi.JSON(strings.NewReader(fmt.Sprintf("[{\"constant\":false,\"inputs\":[{\"name\":\"_nativeToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_tokenMetadata\",\"type\":\"bytes\"}],\"name\":\"completeBridging\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]")))
	// 将 ABI 编码的 calldata 从十六进制字符串解码为字节数组
	decodedCalldata, _ := hex.DecodeString(abiEncodedCalldataHex)
	fmt.Println(decodedCalldata)
	s := decodedCalldata[0:4]
	fmt.Println(s)
	decodedCalldata = decodedCalldata[4:]
	//decodedCalldataL := len(decodedCalldata) / 32
	inputs := method.Methods["completeBridging"].Inputs
	fmt.Println(inputs[0].Type)
	resultInfo := make(map[string]interface{})
	for i := 0; i < len(inputs); i++ {
		startIndex := i * 32
		data := decodedCalldata[startIndex : startIndex+32]
		fmt.Println("11111", inputs[i].Name, inputs[i].Type.String())
		goType := GoType(inputs[i].Type.String(), data)
		resultInfo[inputs[i].Name] = goType
		fmt.Println(goType)
	}
	fmt.Println(resultInfo)

	// 解析函数签名
	//method, err := eth_abi.JSON(strings.NewReader(fmt.Sprintf("[{\"constant\":false,\"inputs\":[{\"name\":\"_nativeToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_tokenMetadata\",\"type\":\"bytes\"}],\"name\":\"completeBridging\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]")))
	//if err != nil {
	//	log.Fatal(err)
	//}
	////// 将 ABI 编码的 calldata 从十六进制字符串解码为字节数组
	//decodedCalldata, err := hex.DecodeString(abiEncodedCalldataHex)
	//fmt.Println(len(decodedCalldata))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//unpack, err := method.Methods["completeBridging"].Inputs.UnpackValues(decodedCalldata[4:])
	//fmt.Println(unpack)

	// 解析 calldata 数据
	//var (
	//	_nativeToken   string
	//	_amount        uint64
	//	_recipient     string
	//	_chainId       uint64
	//	_tokenMetadata []byte
	//)

	//method.Methods["completeBridging"].Inputs.UnpackIntoMap(map[string]interface{}{
	//	"_nativeToken":   &_nativeToken,
	//	"_amount":        &_amount,
	//	"_recipient":     &_recipient,
	//	"_chainId":       &_chainId,
	//	"_tokenMetadata": &_tokenMetadata,
	//}, decodedCalldata)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 输出解析后的参数值
	//fmt.Printf("NativeToken: %s\n", _nativeToken)
	//fmt.Printf("Amount: %d\n", _amount)
	//fmt.Printf("Recipient: %s\n", _recipient)
	//fmt.Printf("ChainId: %d\n", _chainId)
	//fmt.Printf("TokenMetadata: %v\n", _tokenMetadata)

	//unpack, _ := method.Methods["completeBridging"].Inputs.Unpack(decodedCalldata)
	//fmt.Println(unpack)
	// 这是 Solidity 中执行的 abi.encode 的数据（十六进制格式）

}

func GoType(thisType string, data []byte) interface{} {
	switch thisType {
	case "address":
		return common.HexToAddress(hex.EncodeToString(data))
	case "uint256":
		return new(big.Int).SetBytes(data)
	case "bytes":
		return hex.EncodeToString(data)
	case "string":
		return string(data)
	}
	return nil
}

type CallData struct {
	Method string        // Method is the function signature
	Params []interface{} // Params is a list of function parameters
}
