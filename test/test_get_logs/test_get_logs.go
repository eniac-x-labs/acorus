package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

func main() {
	query := ethereum.FilterQuery{}
	query.FromBlock = big.NewInt(41417845)
	query.ToBlock = big.NewInt(41417845)
	//query.Addresses = []common.Address{common.HexToAddress("0x98d8def8aba6e5f3c9ef32363a68428c2834b499")}
	query.Topics = [][]common.Hash{{common.HexToHash("0x7afd75e2560c96bc7dca1d0b63913ecf478fd9feccc33fa450b984d7b8de540a")}}
	client, _ := ethclient.Dial("wss://hidden-nameless-crater.matic-testnet.discover.quiknode.pro/1fce4264658e4fe1714327ee9df7c63741472ae5/")

	logs, _ := client.FilterLogs(context.Background(), query)

	for _, v := range logs {
		marshal, _ := json.Marshal(v)
		//hash := v.TxHash.Hex()
		fmt.Println(string(marshal))
		bytes, _ := rlp.EncodeToBytes(v)
		fmt.Println(bytes)
		//data := v.Data
		//result := log_util.DecodeLog(abiInfo, data, topics[0])
		//fmt.Println(result)
		//fundAddress := v.Address.Hex()
		//fmt.Println(fundAddress)
		//whitelist.WhiteListService.BatchAddWhitelistDecode(result, "80001", v)
		//amountIn := result["amountIn"].(*big.Int).String()
		//amountOut := result["amountOut"].(*big.Int).String()
		//fmt.Println(amountIn)
		//fmt.Println(amountOut)
		//if hash == "0x2d6d1f31705db16af7fd8ed2825d9ee39de6990a8f4ed9fb534a0211538e6f47" {
		//	fmt.Println(string(marshal))
		//}
	}
}
