package routes

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"

	"github.com/cornerstone-labs/acorus/database/common"
	bridgeDb "github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
)

// Routes ... Route handler struct
type Routes struct {
	logger        log.Logger
	view          bridgeDb.BridgeTransfersView
	l1ToL2View    worker.L1ToL2View
	l2ToL1View    worker.L2ToL1View
	stateRootView worker.StateRootView
	blocksView    common.BlocksView
	router        *chi.Mux
	v             *Validator
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(logger log.Logger, bv bridgeDb.BridgeTransfersView, l1l2v worker.L1ToL2View, l2l1v worker.L2ToL1View, blv common.BlocksView, srv worker.StateRootView, r *chi.Mux) Routes {
	return Routes{
		logger:        logger,
		view:          bv,
		l1ToL2View:    l1l2v,
		l2ToL1View:    l2l1v,
		stateRootView: srv,
		blocksView:    blv,
		router:        r,
	}
}
