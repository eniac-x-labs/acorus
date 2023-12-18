package event

import (
	"fmt"

	common2 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/event/processor"
	"github.com/cornerstone-labs/acorus/event/processor/linea"
	"github.com/cornerstone-labs/acorus/event/processor/scroll"
)

func GetEventFactory(params processor.GetEventFactoryParams) (processor.IEventFactory, error) {
	if params.ChainBridge == common2.Scroll {
		return &scroll.ScrollProcessor{}, nil
	}
	if params.ChainBridge == common2.Linea {
		return &linea.LineaProcessor{}, nil
	}
	return nil, fmt.Errorf("Wrong event type passed")
}
