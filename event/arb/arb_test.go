package arb

import (
	"context"
	"fmt"

	"math/big"
	"testing"

	"github.com/cornerstone-labs/acorus/event/arb/abi/bridgegen"
	"github.com/cornerstone-labs/acorus/event/arb/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestAbi_contract(t *testing.T) {
	abi, _ := bridgegen.BridgeMetaData.GetAbi()
	MessageDeliveredId := abi.Events["MessageDelivered"].ID
	fmt.Println(MessageDeliveredId)
	fmt.Println(utils.MessageDeliveredHash)
}

func Test_UnpackMessageDelivered(t *testing.T) {
	query := ethereum.FilterQuery{}
	query.FromBlock = big.NewInt(19416649)
	query.ToBlock = big.NewInt(19416649)
	query.Addresses = []common.Address{common.HexToAddress("0x8315177aB297bA92A06054cE80a67Ed4DBd7ed3a")}
	query.Topics = [][]common.Hash{{common.HexToHash("0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1")}}
	client, _ := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/-EK96JwUb8C_l_EfZqaLumlZG6PV8SDq")

	logs, _ := client.FilterLogs(context.Background(), query)

	for _, v := range logs {
		delivered, err := utils.IBridge.ParseMessageDelivered(v)
		if err != nil {
			fmt.Println(err)
		}
		hash := delivered.Inbox
		fmt.Println(hash.String())
	}
}

func Test_ParseEthDepositData(t *testing.T) {
	messageData := "A6E514E8FCF50C0DA68B45B89616A9FF674B28D6FA83011D03FC6BE3BDA385E1"
	fmt.Println(messageData)
	data, b, _ := utils.ParseEthDepositData(messageData)
	fmt.Println(data.Hex(), common.BigToHash(b).Hex())
}

func Test_CalculateDepositTxId(t *testing.T) {
	messageData := "4F6EB9618EA6D2131B1344A07EDBB612523E33185E72EF0445446C67EA31CBFE"
	data, b, _ := utils.ParseEthDepositData(messageData)
	l2ChainId, _ := big.NewInt(0).SetString("42161", 10)
	fmt.Println(common.BigToHash(l2ChainId).TerminalString())
	//rlp.s
	messageNumber := big.NewInt(1444514)
	utils.CalculateDepositTxId("0xeE488A11475d588908001e0E99e4fD89aBDa6545", data.Hex(), l2ChainId, messageNumber, b)
}
