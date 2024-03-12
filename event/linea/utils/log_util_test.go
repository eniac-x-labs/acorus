package utils

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
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
			log, _ := DecodeLog(abi.L1MessageEventABI, abi.L1MessageCallDataABI, "MessageSent", "distribute", v)
			fmt.Println(log)
		}
	}
}

func TestBase64(t *testing.T) {
	base64String := "zCmjBgAAAAAAAAAAAAAAADbFb6QycSL36jaDstUnwbJHbTLuAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC3H+eG8pAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKPcjZxtRawAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABldSHnAAAAAAAAAAAAAAAAcQvaMpsqYiTktEgz3jDzjn+B1WQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJ5X1jVAAA=="

	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Error("DecodeString ", "err", err)
	}

	decodedString := string(decodedBytes)

	fmt.Println(decodedString)

}

func TestAbi(t *testing.T) {

	// L1MessageSentCalldata
	l1MessageCallDataABI := abi.L1MessageCallDataABI
	l1abiEncodedCalldataHex := "CC29A30600000000000000000000000030FF41D76AA9D7E8042FD08528005AFA01AA161300000000000000000000000000000000000000000000000000108D38DD6D8E0000000000000000000000000000000000000000000000000000107809387B16400000000000000000000000000000000000000000000000000000018D5FA0C9AB000000000000000000000000710BDA329B2A6224E4B44833DE30F38E7F81D56400000000000000000000000000000000000000000000000000049E57D6354000"
	bytesL1, _ := hex.DecodeString(l1abiEncodedCalldataHex)
	bytesL1 = bytesL1[4:]
	datal1, _ := DecodeMessageCallData(l1MessageCallDataABI, abi.L1MessageCallDataFuncName, bytesL1)
	fmt.Println(datal1)

	// L2MessageSentCalldata
	l2MessageCallDataABI := abi.L2MessageCallDataABI

	l2abiEncodedCalldataHex := "E4D274510000000000000000000000002260FAC5E5542A773AA44FBCFEDF7C193BC2C5990000000000000000000000000000000000000000000000000000000017D7840000000000000000000000000048CCE57C4D2DBB31EAF79575ABF482BBB8DC071D000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000A00000000000000000000000000000000000000000000000000000000000000000"
	bytesL2, _ := hex.DecodeString(l2abiEncodedCalldataHex)
	bytesL2 = bytesL2[4:]
	datal2, _ := DecodeMessageCallData(l2MessageCallDataABI, abi.L2MessageCallDataFuncName, bytesL2)
	fmt.Println(datal2)
}

func TestAbiSign(t *testing.T) {
	methods := abi.L2MessageEventABI.Methods
	for method := range methods {
		id := abi.L2MessageEventABI.Methods[method].ID
		fmt.Println(hex.EncodeToString(id))
	}
}
