package main

import (
	"encoding/json"
	"fmt"
	"github.com/cornerstone-labs/acorus/appchain/unpack"
	"github.com/cornerstone-labs/acorus/rpc/bridge/protobuf/pb"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func main() {
	sharesMap := make(map[uint64]*pb.ShareMap)
	for i := 0; i < 5; i++ {
		strategyAmountMap := make(map[string]string)
		for j := 0; j < 5; j++ {
			strategyAmountMap[string(j)] = "1"
			unpack.AddSharesToMap(sharesMap, uint64(i), common.BigToAddress(big.NewInt(int64(1))), big.NewInt(1))
		}

	}
	totalShares := unpack.GetTotalShares(sharesMap)
	fmt.Println(totalShares)
	marshal, _ := json.Marshal(sharesMap)
	fmt.Println(string(marshal))
}
