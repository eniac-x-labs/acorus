package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/cache"
	"github.com/cornerstone-labs/acorus/service/service"
)

type Routes struct {
	logger      log.Logger
	router      *chi.Mux
	svc         service.Service
	enableCache bool
	cache       *cache.LruCache
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(l log.Logger, r *chi.Mux, svc service.Service, enableCache bool, cache *cache.LruCache) Routes {
	return Routes{
		logger:      l,
		router:      r,
		svc:         svc,
		enableCache: enableCache,
		cache:       cache,
	}
}
