package indexer

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processors"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	node2 "github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/cornerstone-labs/acorus/worker"
)

type Acorus struct {
	log               log.Logger
	db                *database.DB
	redis             *redis.Client
	httpConfig        config.ServerConfig
	metricsConfig     config.ServerConfig
	metricsRegistry   *prometheus.Registry
	L1Sync            *synchronizer.L1Sync
	L2Sync            *synchronizer.L2Sync
	BridgeProcessor   *processors.BridgeProcessor
	BusinessProcessor *worker.BusinessProcessor
}

func NewAcorus(
	log log.Logger,
	db *database.DB,
	redis *redis.Client,
	chainConfig config.ChainConfig,
	rpcsConfig config.RPCsConfig,
	httpConfig config.ServerConfig,
	metricsConfig config.ServerConfig,
) (*Acorus, error) {
	l1EthClient, err := node2.DialEthClient(rpcsConfig.L1RPC)
	if err != nil {
		return nil, err
	}
	l1Cfg := synchronizer.Config{
		LoopIntervalMsec:  chainConfig.L1PollingInterval,
		HeaderBufferSize:  chainConfig.L1HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(chainConfig.L1ConfirmationDepth)),
		StartHeight:       big.NewInt(int64(chainConfig.L1StartingHeight)),
	}
	l1Syncer, err := synchronizer.NewL1Sync(l1Cfg, log, db, l1EthClient, chainConfig.L1Contracts, chainConfig.Preset)
	if err != nil {
		return nil, err
	}

	l2EthClient, err := node2.DialEthClient(rpcsConfig.L2RPC)
	if err != nil {
		return nil, err
	}
	l2Cfg := synchronizer.Config{
		LoopIntervalMsec:  chainConfig.L2PollingInterval,
		HeaderBufferSize:  chainConfig.L2HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(chainConfig.L2ConfirmationDepth)),
	}
	l2Syncer, err := synchronizer.NewL2Sync(l2Cfg, log, db, l2EthClient, chainConfig.L2Contracts)
	if err != nil {
		return nil, err
	}

	bridgeProcessor, err := processors.NewBridgeProcessor(log, db, l1Syncer, chainConfig)
	if err != nil {
		return nil, err
	}

	businessProcessor := worker.NewBusinessProcessor(log, db, redis)

	acorus := &Acorus{
		log:               log,
		db:                db,
		redis:             redis,
		httpConfig:        httpConfig,
		metricsConfig:     metricsConfig,
		L1Sync:            l1Syncer,
		L2Sync:            l2Syncer,
		BridgeProcessor:   bridgeProcessor,
		BusinessProcessor: businessProcessor,
	}
	return acorus, nil
}

func (i *Acorus) startHttpServer(ctx context.Context) error {
	i.log.Debug("starting http server...", "port", i.httpConfig.Host)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))

	addr := net.JoinHostPort(i.httpConfig.Host, strconv.Itoa(i.httpConfig.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	i.log.Info("http server started", "addr", srv.Addr())
	<-ctx.Done()
	defer i.log.Info("http server stopped")
	return srv.Stop(context.Background())
}

func (i *Acorus) startMetricsServer(ctx context.Context) error {
	return nil
}

func (i *Acorus) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 5)

	// if any goroutine halts, we stop the entire Acorus
	processCtx, processCancel := context.WithCancel(ctx)
	runProcess := func(start func(ctx context.Context) error) {
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					i.log.Error("halting acorus on panic", "err", err)
					debug.PrintStack()
					errCh <- fmt.Errorf("panic: %v", err)
				}
				processCancel()
				wg.Done()
			}()
			errCh <- start(processCtx)
		}()
	}

	// synchronizer engine
	runProcess(i.L1Sync.Start)
	runProcess(i.L2Sync.Start)

	// event engine
	runProcess(i.BridgeProcessor.Start)

	// worker engine
	runProcess(i.BusinessProcessor.Start)

	// metrics server
	runProcess(i.startMetricsServer)

	// service server
	runProcess(i.startHttpServer)

	wg.Wait()

	err := <-errCh
	if err != nil {
		i.log.Error("acorus stopped", "err", err)
	} else {
		i.log.Info("acorus stopped")
	}
	return err
}
