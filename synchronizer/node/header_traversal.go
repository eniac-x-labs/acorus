package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/cornerstone-labs/acorus/common/bigint"
)

var (
	ErrHeaderTraversalAheadOfProvider            = errors.New("the HeaderTraversal's internal state is ahead of the provider")
	ErrHeaderTraversalAndProviderMismatchedState = errors.New("the HeaderTraversal and provider have diverged in state")
)

type HeaderTraversal struct {
	ethClient EthClient
	chainId   uint

	latestHeader        *types.Header
	lastTraversedHeader *types.Header

	blockConfirmationDepth *big.Int
}

// NewHeaderTraversal instantiates a new instance of HeaderTraversal against the supplied rpc client.
// The HeaderTraversal will start fetching blocks starting from the supplied header unless nil, indicating genesis.
func NewHeaderTraversal(ethClient EthClient, fromHeader *types.Header, confDepth *big.Int, chainId uint) *HeaderTraversal {
	return &HeaderTraversal{
		ethClient:              ethClient,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
		chainId:                chainId,
	}
}

// LatestHeader returns the latest header reported by underlying eth client
// as headers are traversed via `NextHeaders`.
func (f *HeaderTraversal) LatestHeader() *types.Header {
	return f.latestHeader
}

// LastTraversedHeader returns the last header traversed.
//   - This is useful for testing the state of the HeaderTraversal
//   - LastTraversedHeader may be << LatestHeader depending on the number
//     headers traversed via `NextHeaders`.
func (f *HeaderTraversal) LastTraversedHeader() *types.Header {
	return f.lastTraversedHeader
}

// NextHeaders retrieves the next set of headers that have been
// marked as finalized by the connected client, bounded by the supplied size
func (f *HeaderTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := f.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestHeader = latestHeader
	}

	endHeight := new(big.Int).Sub(latestHeader.Number, f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		// No blocks with the provided confirmation depth available
		return nil, nil
	}

	if f.lastTraversedHeader != nil {
		cmp := f.lastTraversedHeader.Number.Cmp(endHeight)
		if cmp == 0 { // We're synced to head and there are no new headers
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrHeaderTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedHeader != nil {
		nextHeight = new(big.Int).Add(f.lastTraversedHeader.Number, bigint.One)
	}

	// endHeight = (nextHeight - endHeight) <= maxSize
	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.ethClient.BlockHeadersByRange(nextHeight, endHeight, f.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}

	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	}

	f.lastTraversedHeader = &headers[numHeaders-1]
	return headers, nil
}
