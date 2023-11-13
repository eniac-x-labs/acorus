package routes

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"

	"github.com/cornerstone-labs/acorus/database/business"
	"github.com/cornerstone-labs/acorus/database/common"
	bridgeDb "github.com/cornerstone-labs/acorus/database/event/l1-l2"
)

// Routes ... Route handler struct
type Routes struct {
	logger        log.Logger
	view          bridgeDb.BridgeTransfersView
	l1ToL2View    business.L1ToL2View
	l2ToL1View    business.L2ToL1View
	stateRootView business.StateRootView
	blocksView    common.BlocksView
	router        *chi.Mux
	v             *Validator
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(logger log.Logger, bv bridgeDb.BridgeTransfersView, l1l2v business.L1ToL2View, l2l1v business.L2ToL1View, blv common.BlocksView, srv business.StateRootView, r *chi.Mux) Routes {
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
