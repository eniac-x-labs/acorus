package utils

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/arb/abi/bridgegen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"strings"
)

var (
	WithdrawalInitiatedHash       = "0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73"
	L2ToL1TxHash                  = "0x3e7aafa77dbf186b7fd488006beff893744caa3c4f6f299e8a709fa2087374fc"
	OutBoxTransactionExecutedHash = "0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964"
	MessageDeliveredHash          = "0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1"
	IBridge                       = new(bridgegen.IBridge)
)

func init() {
	logger := log.New("IBridge", IBridge)
	if bridge, err := bridgegen.NewIBridge(common.HexToAddress("0x8315177aB297bA92A06054cE80a67Ed4DBd7ed3a"), nil); err != nil {
		logger.Info("failed to create IBridge contract instance", "error", err)
		panic(err)
	} else {
		logger.Info("successfully created IBridge contract instance")

		IBridge = bridge
	}
}

func ParseEthDepositData(messageData string) (common.Address, *big.Int, error) {
	// https://github.com/OffchainLabs/nitro/blob/aa84e899cbc902bf6da753b1d66668a1def2c106/contracts/src/bridge/Inbox.sol#L242
	// ethers.defaultAbiCoder doesnt decode packed args, so we do a hardcoded parsing
	const addressEnd = 2 + 20*2
	address := "0x" + messageData[2:addressEnd]
	value := messageData[addressEnd:]
	if setString, ok := new(big.Int).SetString(value, 16); !ok {
		return common.Address{}, nil, fmt.Errorf("format big.int error %v", ok)
	} else {
		return common.HexToAddress(address), setString, nil
	}
	//bytes := common.Hex2Bytes(messageData)
	//rlp.DecodeBytes(bytes,)
}

// l2ChainId: number,
// messageNumber: BigNumber,
// fromAddress: string,
// toAddress: string,
// value: BigNumber

type ArbitrumDepositTx struct {
	ChainId     *big.Int
	L1RequestId common.Hash
	From        common.Address
	To          common.Address
	Value       *big.Int
}

func CalculateDepositTxId(fromAddress, toAddress string, l2ChainId, messageNumber, value *big.Int) {
	// https://github.com/OffchainLabs/go-ethereum/blob/07e017aa73e32be92aadb52fa327c552e1b7b118/core/types/arb_types.go#L302-L308
	//rlp.EncodeToBytes()

	//fields := string struct{
	//	l2ChainId,
	//	common.BigToHash(messageNumber),
	//	common.HexToAddress(fromAddress),
	//	common.HexToAddress(toAddress),
	//	value,
	//}

	bytes, err := rlp.EncodeToBytes(l2ChainId.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	//rlpEnc := "0x64" + string(bytes)
	//fmt.Println("RLP Encoded:", rlpEnc)
	h := crypto.Keccak256Hash(bytes)
	fmt.Println("l2 txHash ", h)
}
func hexConcat(fields []string) string {
	// 处理十六进制字符串的拼接逻辑，根据具体需求实现
	return strings.Join(fields, "")
}

//func rlpEncode(fields []string) string {
//	// 处理 RLP 编码的逻辑，根据具体需求实现
//	encodedFields := []string{}
//	for _, field := range fields {
//		encodedField := encodeField(field)
//		encodedFields = append(encodedFields, encodedField)
//	}
//	return hexConcat(encodedFields)
//}
