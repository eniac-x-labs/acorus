package contracts

import (
	"github.com/cornerstone-labs/acorus/event/processors/op-stack/mantle/op-bindings/bindings"
	"github.com/stretchr/testify/require"
	"testing"
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