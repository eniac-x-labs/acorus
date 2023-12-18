package processor

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type IEventUnpack interface {
	Start(ctx context.Context) error
	GetLatestHeader(chainId uint64) (*types.Header, *types.Header, error)
	Run() error
}
