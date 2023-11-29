package bridge

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/processors/scroll/abi"
	"testing"
)

func TestGetAbi_Test(t *testing.T) {
	l1DepositETHSig := abi.L1DepositETHSig
	fmt.Println(l1DepositETHSig)
	fmt.Println(abi.L1ERC721GatewayABI.Events)
	fmt.Println(abi.L1ERC1155GatewayABI.Events)

	fmt.Println(abi.L2ERC721GatewayABI.Events)
	fmt.Println(abi.L2ERC1155GatewayABI.Events)

}
