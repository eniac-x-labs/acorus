package acorus

import (
	"context"
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/database/create_table"
	"math/big"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/event/processor"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
)

type Acorus struct {
	log             log.Logger
	DB              *database.DB
	l1Client        node.EthClient
	l2Client        node.EthClient
	apiServer       *httputil.HTTPServer
	metricsServer   *httputil.HTTPServer
	metricsRegistry *prometheus.Registry
	L1Sync          *synchronizer.L1Sync
	L2Sync          *synchronizer.L2Sync
	EventProcessor  *processor.EventProcessor
	shutdown        context.CancelCauseFunc
	stopped         atomic.Bool
}

func NewAcorus(ctx context.Context, log log.Logger, cfg *config.Config, shutdown context.CancelCauseFunc) (*Acorus, error) {
	out := &Acorus{
		log:      log,
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}

func (as *Acorus) Start(ctx context.Context) error {
	if err := as.L1Sync.Start(); err != nil {
		return fmt.Errorf("failed to start L1 Sync: %w", err)
	}
	if err := as.L2Sync.Start(); err != nil {
		return fmt.Errorf("failed to start L2 Sync: %w", err)
	}
	if err := as.EventProcessor.Start(); err != nil {
		return fmt.Errorf("failed to start EventProcessor: %w", err)
	}
	return nil
}

func (as *Acorus) Stop(ctx context.Context) error {
	var result error
	if as.L1Sync != nil {
		if err := as.L1Sync.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close L1 Sync: %w", err))
		}
	}

	if as.L2Sync != nil {
		if err := as.L2Sync.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close L2 Sync: %w", err))
		}
	}

	if as.l1Client != nil {
		as.l1Client.Close()
	}
	if as.l2Client != nil {
		as.l2Client.Close()
	}

	if as.apiServer != nil {
		if err := as.apiServer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close acorus API server: %w", err))
		}
	}

	if as.DB != nil {
		if err := as.DB.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	if as.metricsServer != nil {
		if err := as.metricsServer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close metrics server: %w", err))
		}
	}

	as.stopped.Store(true)

	as.log.Info("acorus stopped")

	return result
}

func (as *Acorus) Stopped() bool {
	return as.stopped.Load()
}

func (as *Acorus) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := as.initRPCClients(ctx, cfg.RPCs); err != nil {
		return fmt.Errorf("failed to start RPC clients: %w", err)
	}
	if err := as.initDB(ctx, cfg.MasterDB); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	// init table
	as.initTable(cfg)

	if err := as.initL1Syncer(cfg.Chain, cfg.Chain.ChainId); err != nil {
		return fmt.Errorf("failed to init L1 Sync: %w", err)
	}
	if err := as.initL2ETL(cfg.Chain, cfg.Chain.ChainId); err != nil {
		return fmt.Errorf("failed to init L2 Sync: %w", err)
	}
	if err := as.initBridgeProcessor(cfg.Chain, cfg.RPCs); err != nil {
		return fmt.Errorf("failed to init Bridge Processor: %w", err)
	}
	if err := as.initBusinessProcessor(*cfg); err != nil {
		return fmt.Errorf("failed to init Business Processor: %w", err)
	}
	if err := as.startHttpServer(ctx, cfg.HTTPServer); err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}
	if err := as.startMetricsServer(ctx, cfg.MetricsServer); err != nil {
		return fmt.Errorf("failed to start Metrics server: %w", err)
	}
	return nil
}

func (as *Acorus) initRPCClients(ctx context.Context, rpcsConfig config.RPCsConfig) error {
	l1EthClient, err := node.DialEthClient(ctx, rpcsConfig.L1RPC)
	if err != nil {
		return fmt.Errorf("failed to dial L1 client: %w", err)
	}
	as.l1Client = l1EthClient

	l2EthClient, err := node.DialEthClient(ctx, rpcsConfig.L2RPC)
	if err != nil {
		return fmt.Errorf("failed to dial L2 client: %w", err)
	}
	as.l2Client = l2EthClient
	return nil
}

func (as *Acorus) initDB(ctx context.Context, cfg config.DBConfig) error {
	db, err := database.NewDB(ctx, as.log, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	as.DB = db
	return nil
}

func (as *Acorus) initTable(cfg *config.Config) {
	create_table.CreateTableService.CreateInit(int64(cfg.Chain.ChainId), as.DB)
}

func (as *Acorus) initL1Syncer(chainConfig config.ChainConfig, chainId uint) error {
	l1Cfg := synchronizer.Config{
		LoopIntervalMsec:  chainConfig.L1PollingInterval,
		HeaderBufferSize:  chainConfig.L1HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(chainConfig.L1ConfirmationDepth)),
		StartHeight:       big.NewInt(int64(chainConfig.L1StartingHeight)),
		ChainId:           chainId,
	}
	l1Sync, err := synchronizer.NewL1Sync(l1Cfg, as.log, as.DB, as.l1Client, chainConfig, as.shutdown, chainId)
	if err != nil {
		return err
	}
	as.L1Sync = l1Sync
	return nil
}

func (as *Acorus) initL2ETL(chainConfig config.ChainConfig, chainId uint) error {
	l2Cfg := synchronizer.Config{
		LoopIntervalMsec:  chainConfig.L2PollingInterval,
		HeaderBufferSize:  chainConfig.L2HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(chainConfig.L2ConfirmationDepth)),
		StartHeight:       big.NewInt(int64(chainConfig.L2StartingHeight)),
	}
	l2Sync, err := synchronizer.NewL2Sync(l2Cfg, as.log, as.DB, as.l2Client, chainConfig, as.shutdown, chainId)
	if err != nil {
		return err
	}
	as.L2Sync = l2Sync
	return nil
}

func (as *Acorus) initBridgeProcessor(chainConfig config.ChainConfig, rpcsConfig config.RPCsConfig) error {
	processor, err := processor.NewEventProcessor(as.log, as.DB, as.L1Sync, as.L2Sync, chainConfig, rpcsConfig, as.shutdown)
	if err != nil {
		return err
	}
	as.EventProcessor = processor
	return nil
}

func (as *Acorus) initBusinessProcessor(cfg config.Config) error {
	return nil
}

func (as *Acorus) startHttpServer(ctx context.Context, cfg config.ServerConfig) error {
	as.log.Debug("starting http server...", "port", cfg.Port)
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	as.apiServer = srv
	as.log.Info("http server started", "addr", srv.Addr())
	return nil
}

func (as *Acorus) startMetricsServer(ctx context.Context, cfg config.ServerConfig) error {
	return nil
}
