package linea

import (
	"context"
	"fmt"

	"github.com/cornerstone-labs/acorus/event/processor"
)

type LineaProcessor struct {
	processor.IEventUnpack
}

// GetChainUnpackService get service implement
func (l *LineaProcessor) GetChainUnpackService(eventParam processor.GetEventFactoryParams) processor.IEventUnpack {
	return &LineaProcessor{}
}

func (l *LineaProcessor) Start(ctx context.Context) error {
	fmt.Println("LineaProcessor Start")
	return nil
}

func (l *LineaProcessor) Run() error {
	return nil
}
