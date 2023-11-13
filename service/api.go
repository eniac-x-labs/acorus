package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database/business"
	common2 "github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/service/routes"
)

const ethereumAddressRegex = `^0x[a-fA-F0-9]{40}$`

type API struct {
	log             log.Logger
	router          *chi.Mux
	serverConfig    config.ServerConfig
	metricsConfig   config.ServerConfig
	metricsRegistry *prometheus.Registry
}

const (
	MetricsNamespace = "acorus_api"

	HealthPath        = "/healthz"
	DepositsV1Path    = "/service/v1/deposits"
	WithdrawalsV1Path = "/service/v1/withdrawals"
)

func NewApi(logger log.Logger, bv l1_l2.BridgeTransfersView, l1l2v business.L1ToL2View, l2l1v business.L2ToL1View, blv common2.BlocksView, srv business.StateRootView, serverConfig config.ServerConfig, metricsConfig config.ServerConfig) *API {
	apiRouter := chi.NewRouter()
	h := routes.NewRoutes(logger, bv, l1l2v, l2l1v, blv, srv, apiRouter)

	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.Heartbeat(HealthPath))

	apiRouter.Get(fmt.Sprintf(DepositsV1Path), h.L1ToL2ListHandler)
	apiRouter.Get(fmt.Sprintf(WithdrawalsV1Path), h.L2ToL1ListHandler)

	return &API{log: logger, router: apiRouter, serverConfig: serverConfig, metricsConfig: metricsConfig}
}

func (a *API) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	processCtx, processCancel := context.WithCancel(ctx)
	runProcess := func(start func(ctx context.Context) error) {
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					a.log.Error("halting service on panic", "err", err)
					debug.PrintStack()
					errCh <- fmt.Errorf("panic: %v", err)
				}
				processCancel()
				wg.Done()
			}()

			errCh <- start(processCtx)
		}()
	}

	runProcess(a.startServer)
	runProcess(a.startMetricsServer)

	wg.Wait()
	err := <-errCh
	if err != nil {
		a.log.Error("service stopped", "err", err)
	} else {
		a.log.Info("service stopped")
	}

	return err
}

func (a *API) Port() int {
	return a.serverConfig.Port
}

func (a *API) startServer(ctx context.Context) error {
	a.log.Debug("service server listening...", "port", a.serverConfig.Port)
	addr := net.JoinHostPort(a.serverConfig.Host, strconv.Itoa(a.serverConfig.Port))
	srv, err := httputil.StartHTTPServer(addr, a.router)
	if err != nil {
		return fmt.Errorf("failed to start API server: %w", err)
	}

	host, portStr, err := net.SplitHostPort(srv.Addr().String())
	if err != nil {
		return errors.Join(err, srv.Close())
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return errors.Join(err, srv.Close())
	}

	a.serverConfig.Host = host
	a.serverConfig.Port = port

	<-ctx.Done()
	if err := srv.Stop(context.Background()); err != nil {
		return fmt.Errorf("failed to shutdown service server: %w", err)
	}
	return nil
}

func (a *API) startMetricsServer(ctx context.Context) error {
	return nil
}
