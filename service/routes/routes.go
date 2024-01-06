package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/cornerstone-labs/acorus/cache"
	"github.com/cornerstone-labs/acorus/service/service"
)

type Routes struct {
	router      *chi.Mux
	svc         service.Service
	enableCache bool
	cache       *cache.LruCache
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(r *chi.Mux, svc service.Service, enableCache bool, cache *cache.LruCache) Routes {
	return Routes{
		router:      r,
		svc:         svc,
		enableCache: enableCache,
		cache:       cache,
	}
}
