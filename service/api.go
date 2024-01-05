package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/cache"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/service/routes"
	"github.com/cornerstone-labs/acorus/service/service"
)

const (
	HealthPath        = "/healthz"
	DepositsV1Path    = "/api/v1/deposits"
	WithdrawalsV1Path = "/api/v1/withdrawals"
)

type APIConfig struct {
	HTTPServer    config.Server
	MetricsServer config.Server
}

type API struct {
	log             log.Logger
	router          *chi.Mux
	metricsRegistry *prometheus.Registry
	apiServer       *httputil.HTTPServer
	metricsServer   *httputil.HTTPServer
	db              *database.DB
	stopped         atomic.Bool
}

func NewApi(ctx context.Context, log log.Logger, cfg *config.Config) (*API, error) {
	out := &API{log: log}
	if err := out.initFromConfig(ctx, cfg, log); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}

func (a *API) initFromConfig(ctx context.Context, cfg *config.Config, log log.Logger) error {
	if err := a.initDB(ctx, cfg, log); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	if err := a.startMetricsServer(cfg.Metrics); err != nil {
		return fmt.Errorf("failed to start metrics server: %w", err)
	}
	a.initRouter(cfg)
	if err := a.startServer(cfg.Server); err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	return nil
}

func (a *API) initRouter(conf *config.Config) {
	v := new(service.Validator)

	var lruCache = new(cache.LruCache)
	if conf.EnableApiCache {
		lruCache = cache.NewLruCache()
	}

	svc := service.New(v, a.db.L1ToL2, a.db.L2ToL1, a.log)
	apiRouter := chi.NewRouter()
	h := routes.NewRoutes(a.log, apiRouter, svc, conf.EnableApiCache, lruCache)

	apiRouter.Use(middleware.Timeout(time.Second * 12))
	apiRouter.Use(middleware.Recoverer)

	apiRouter.Use(middleware.Heartbeat(HealthPath))

	apiRouter.Get(fmt.Sprintf(DepositsV1Path), h.L1ToL2ListHandler)
	apiRouter.Get(fmt.Sprintf(WithdrawalsV1Path), h.L2ToL1ListHandler)

	a.router = apiRouter
}

func (a *API) initDB(ctx context.Context, cfg *config.Config, log log.Logger) error {
	var initDb *database.DB
	var err error
	if !cfg.SlaveDbEnable {
		initDb, err = database.NewDB(ctx, log, cfg.MasterDb)
		if err != nil {
			log.Error("failed to connect to master database", "err", err)
			return err
		}
	} else {
		initDb, err = database.NewDB(ctx, log, cfg.SlaveDb)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
	}
	a.db = initDb
	return nil
}

func (a *API) Start(ctx context.Context) error {
	return nil
}

func (a *API) Stop(ctx context.Context) error {
	var result error
	if a.apiServer != nil {
		if err := a.apiServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop API server: %w", err))
		}
	}
	if a.metricsServer != nil {
		if err := a.metricsServer.Stop(ctx); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to stop metrics server: %w", err))
		}
	}
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}
	a.stopped.Store(true)
	a.log.Info("API service shutdown complete")
	return result
}

func (a *API) startServer(serverConfig config.Server) error {
	a.log.Debug("API server listening...", "port", serverConfig.Port)
	addr := net.JoinHostPort(serverConfig.Host, strconv.Itoa(serverConfig.Port))
	srv, err := httputil.StartHTTPServer(addr, a.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}
	a.log.Info("API server started", "addr", srv.Addr().String())
	a.apiServer = srv
	return nil
}

func (a *API) startMetricsServer(metricsConfig config.Server) error {
	return nil
}

func (a *API) Stopped() bool {
	return a.stopped.Load()
}
