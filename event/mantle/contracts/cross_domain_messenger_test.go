package contracts

import (
	"fmt"
	"github.com/cornerstone-labs/acorus/event/mantle/op-bindings/bindings"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestCrossDomainMessengerSentMessageEvents(t *testing.T) {
	crossDomainMessengerAbi, err := bindings.CrossDomainMessengerMetaData.GetAbi()
	require.NoError(t, err)
	l1CrossDomainMessengerAbi, err := bindings.L1CrossDomainMessengerMetaData.GetAbi()
	require.NoError(t, err)
	l2CrossDomainMessengerAbi, err := bindings.L2CrossDomainMessengerMetaData.GetAbi()
	require.NoError(t, err)
	baseSentMessageEventAbi := crossDomainMessengerAbi.Events["SentMessage"]
	l1SentMessageEventAbi := l1CrossDomainMessengerAbi.Events["SentMessage"]
	l2SentMessageEventAbi := l2CrossDomainMessengerAbi.Events["SentMessage"]
	require.Equal(t, baseSentMessageEventAbi.ID.String(), l1SentMessageEventAbi.ID.String(), l2SentMessageEventAbi.ID.String())
	baseSentMessageExtensionEventAbi := crossDomainMessengerAbi.Events["SentMessageExtension1"]
	l1SentMessageExtensionEventAbi := l1CrossDomainMessengerAbi.Events["SentMessageExtension1"]
	l2SentMessageExtensionEventAbi := l2CrossDomainMessengerAbi.Events["SentMessageExtension1"]
	require.Equal(t, baseSentMessageExtensionEventAbi.ID.String(), l1SentMessageExtensionEventAbi.ID.String(), l2SentMessageExtensionEventAbi.ID.String())

}

// 0x5a94ba494dd912a96ec9cd02b044f9f5876a41c67454760fab8eabd234dc19f0
func TestCrossDomainMessengerV0MessageHash(t *testing.T) {
	sentMessage := bindings.CrossDomainMessengerSentMessage{
		Target:       common.HexToAddress("0x21F308067241B2028503c07bd7cB3751FFab0Fb2"),
		Sender:       common.HexToAddress("0x4200000000000000000000000000000000000010"),
		Message:      []byte("0xa9f9e6750000000000000000000000000e7f22a690ffd2b70c21ef2ea2e5bb5de5ae0ee6000000000000000000000000d1236015605d816ce8fa66b10797c667dfab1b6a00000000000000000000000074a5ca71176cbcedcf3cff3b9ec5095be55f3900000000000000000000000000b6e9efce878cf6cfb56232881211e005b61cf3fd000000000000000000000000000000000000000000000000000000003b9aca0000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000000"),
		MessageNonce: big.NewInt(100125),
		// GasLimit:     big.NewInt(300000),
	}
	//mntValue := big.NewInt(0)
	//ethValue := big.NewInt(300000)
	messageCalldata, _ := CrossDomainMessageCalldataV0(&sentMessage)
	fmt.Println("message call data hash===", crypto.Keccak256Hash(messageCalldata))
}
