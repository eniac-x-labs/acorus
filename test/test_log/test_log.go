package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"
)

func main() {
	// 以太坊节点的URL
	nodeURL := "https://rpc.ankr.com/eth/7f3ae11204e03961b67c557c4996244f0a53222b23c31a7baf9ae91c6bd89702"

	// 以太坊区块号
	blockNumber := big.NewInt(19034517)

	// 日志的合约地址
	contractAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// 创建以太坊客户端
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal("连接以太坊节点错误:", err)
	}

	// 创建过滤器查询参数
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{contractAddress},
	}

	// 执行日志查询
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("执行日志查询错误:", err)
	}

	// 处理查询结果
	for _, log := range logs {
		marshal, _ := json.Marshal(log)
		//hash := v.TxHash.Hex()
		fmt.Println(string(marshal))
		//fmt.Println(log.BlockNumber)
		//logJson, _ := log.MarshalJSON()
		// 这里写了测试类，没用的。原生的log，就无法变成数组
		//bytes, _ := rlp.EncodeToBytes(marshal)
		//fmt.Println(err)
		hexStr := hexutil.Encode(marshal)
		//toString := hex.EncodeToString(bytes)
		fmt.Println(hexStr)
		//fmt.Println(toString)
		//decodeString, errDec := hex.DecodeString(toString)
		//fmt.Println(errDec)
		decode, _ := hexutil.Decode(hexStr)
		//fmt.Println(decode)
		newLog := new(types.Log)

		err := rlp.DecodeBytes(decode, &newLog)
		fmt.Println(err)
		fmt.Println(newLog.BlockNumber)
		marshalD, _ := json.Marshal(newLog)
		fmt.Println(string(marshalD))

		//var newLog types.Log
		//err := newLog.UnmarshalJSON(decode)
		//fmt.Println(err)
		//fmt.Println(newLog.BlockNumber)
		//fmt.Println("日志索引:", log.Index)
		//fmt.Println("区块号:", log.BlockNumber)
		//fmt.Println("交易哈希:", log.TxHash.Hex())
		//fmt.Println("日志数据:", log.Data)
		//fmt.Println("日志主题:", strings.Join(log.Topics, ", "))
		fmt.Println("----------------------------")
	}
}
