package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/relayer/bindings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func TestAbi_contract(t *testing.T) {
	msgAbi, _ := bindings.IMessageManagerMetaData.GetAbi()
	for s := range msgAbi.Events {
		fmt.Println(s + "-----" + msgAbi.Events[s].ID.String())
	}

	l1Abi, _ := bindings.L1PoolManagerMetaData.GetAbi()
	for s := range l1Abi.Events {
		fmt.Println(s + "-----" + l1Abi.Events[s].ID.String())
	}

	l2Abi, _ := bindings.L2PoolManagerMetaData.GetAbi()
	for s := range l2Abi.Events {
		fmt.Println(s + "-----" + l2Abi.Events[s].ID.String())
	}

}

func TestUnpack(t *testing.T) {
	filterer, _ := bindings.NewL1PoolManagerFilterer(common.Address{}, nil)
	// 以太坊节点的URL
	nodeURL := "https://eth-sepolia.g.alchemy.com/v2/o0zYyGdlny_0bvOGS8qtNdMG4REyM0I_"

	// 以太坊区块号
	blockNumber := big.NewInt(5218875)

	// 日志的合约地址
	contractAddress := common.HexToAddress("0xAC13457009a2aE49820Bd7517818436dE2202549")

	// 创建以太坊客户端
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal("连接以太坊节点错误:", err)
	}
	hash := common.HexToHash("0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735")

	// 创建过滤器查询参数
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{hash}},
	}

	// 执行日志查询
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal("执行日志查询错误:", err)
	}

	// 处理查询结果
	for _, log := range logs {
		marshal, _ := json.Marshal(log)
		fmt.Println(string(marshal))

		event, errr := filterer.ParseStarkingERC20Event(log)
		fmt.Println(errr)
		fmt.Println(event)
	}
}
