package acorus

import (
	"context"
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/event/linea"
	op_stack2 "github.com/cornerstone-labs/acorus/worker/op-stack"
	"log"
	"math/big"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	event2 "github.com/cornerstone-labs/acorus/event"
	"github.com/cornerstone-labs/acorus/event/op_stack"
	"github.com/cornerstone-labs/acorus/event/polygon"
	"github.com/cornerstone-labs/acorus/event/scroll"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
)

type Acorus struct {
	DB              *database.DB
	ethClient       map[uint64]node.EthClient
	apiServer       *httputil.HTTPServer
	metricsServer   *httputil.HTTPServer
	metricsRegistry *prometheus.Registry
	Synchronizer    map[uint64]*synchronizer.Synchronizer
	Processor       map[uint64]event2.IEventProcessor
	Worker          map[uint64]*op_stack2.WorkerProcessor
	shutdown        context.CancelCauseFunc
	stopped         atomic.Bool
	chainIdList     []uint64
}

func NewAcorus(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*Acorus, error) {
	log.Println("New acrous start️ 🕖")
	out := &Acorus{
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	log.Println("New acrous success🏅️")
	return out, nil
}

func (as *Acorus) Start(ctx context.Context) error {
	for i := range as.chainIdList {
		log.Println("starting Sync", "chainId", as.chainIdList[i])
		realChainId := as.chainIdList[i]
		if err := as.Synchronizer[realChainId].Start(); err != nil {
			return fmt.Errorf("failed to start L1 Sync: %w", err)
		}
		processor := as.Processor[realChainId]
		if processor != nil {
			if err := processor.StartUnpack(); err != nil {
				return fmt.Errorf("failed to start event: %w", err)
			}
		}
		worker := as.Worker[realChainId]
		if worker != nil {
			if err := worker.Start(context.Background()); err != nil {
				return fmt.Errorf("failed to start worker: %w", err)
			}
		}
	}
	return nil
}

func (as *Acorus) Stop(ctx context.Context) error {
	var result error
	for i := range as.chainIdList {
		if as.Synchronizer[as.chainIdList[i]] != nil {
			if err := as.Synchronizer[as.chainIdList[i]].Close(); err != nil {
				result = errors.Join(result, fmt.Errorf("failed to close L2 Synchronizer: %w", err))
			}
		}
		if as.ethClient[as.chainIdList[i]] != nil {
			as.ethClient[as.chainIdList[i]].Close()
		}
		if as.Processor[as.chainIdList[i]] != nil {
			as.Processor[as.chainIdList[i]].ClosetUnpack()
		}
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

	log.Println("acorus stopped")

	return result
}

func (as *Acorus) Stopped() bool {
	return as.stopped.Load()
}

func (as *Acorus) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := as.initRPCClients(ctx, cfg); err != nil {
		return fmt.Errorf("failed to start RPC clients: %w", err)
	}
	if err := as.initDB(ctx, cfg.MasterDb); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	if err := as.initSynchronizer(cfg); err != nil {
		return fmt.Errorf("failed to init L1 Sync: %w", err)
	}
	if err := as.initEventProcessor(cfg); err != nil {
		return fmt.Errorf("failed to init Processor: %w", err)
	}
	return nil
}

func (as *Acorus) initRPCClients(ctx context.Context, conf *config.Config) error {
	for i := range conf.RPCs {
		log.Println("Init rpc client", "ChainId", conf.RPCs[i].ChainId, "RpcUrl", conf.RPCs[i].RpcUrl)
		rpc := conf.RPCs[i]
		ethClient, err := node.DialEthClient(ctx, rpc.RpcUrl)
		if err != nil {
			log.Fatalf("dial eth client fail", "err", err)
			return fmt.Errorf("Failed to dial L1 client: %w", err)
		}
		if as.ethClient == nil {
			as.ethClient = make(map[uint64]node.EthClient)
		}
		as.ethClient[rpc.ChainId] = ethClient
		as.chainIdList = append(as.chainIdList, rpc.ChainId)
	}
	log.Println("Init rpc client success")
	return nil
}

func (as *Acorus) initDB(ctx context.Context, cfg config.Database) error {
	db, err := database.NewDB(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	as.DB = db
	log.Println("Init database success")
	return nil
}

func (as *Acorus) initEventProcessor(cfg *config.Config) error {
	var loopInterval time.Duration = time.Second * 5
	var epoch uint64 = 10_000
	var l1StartBlockNumber *big.Int
	var l1RPC *config.RPC
	for i := range cfg.RPCs {
		if cfg.RPCs[i].ChainId == global_const.EthereumChainId {
			l1RPC = cfg.RPCs[i]
		}
	}
	for i := range cfg.RPCs {
		log.Println("init chain processor", "chainId", cfg.RPCs[i].ChainId)
		if as.Processor == nil {
			as.Processor = make(map[uint64]event2.IEventProcessor)
		}
		rpcItem := cfg.RPCs[i]
		var processor event2.IEventProcessor
		var err error
		if rpcItem.ChainId == global_const.EthereumChainId {
			l1StartBlockNumber = big.NewInt(int64(rpcItem.StartBlock))
		}
		if rpcItem.ChainId == global_const.ScrollChainId {
			processor, err = scroll.NewBridgeProcessor(as.DB, rpcItem, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		} else if rpcItem.ChainId == global_const.PolygonChainId {
			processor, err = polygon.NewBridgeProcessor(as.DB, l1RPC, rpcItem, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		} else if rpcItem.ChainId == global_const.OpChinId {
			processor, err = op_stack.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
			)
			if err != nil {
				return err
			}
		} else if rpcItem.ChainId == global_const.LineaChainId {
			processor, err = linea.NewBridgeProcessor(as.DB, rpcItem, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		}
		if processor != nil {
			as.Processor[rpcItem.ChainId] = processor
		}
	}
	return nil
}

func (as *Acorus) initSynchronizer(config *config.Config) error {
	for i := range config.RPCs {
		log.Println("Init synchronizer success", "chainId", config.RPCs[i].ChainId)
		rpcItem := config.RPCs[i]
		cfg := synchronizer.Config{
			LoopIntervalMsec:  5,
			HeaderBufferSize:  500,
			ConfirmationDepth: big.NewInt(int64(1)),
			StartHeight:       big.NewInt(int64(rpcItem.StartBlock)),
			ChainId:           uint(rpcItem.ChainId),
		}
		synchronizerTemp, err := synchronizer.NewSynchronizer(&cfg, as.DB, as.ethClient[rpcItem.ChainId], as.shutdown)
		if err != nil {
			return err
		}
		if as.Synchronizer == nil {
			as.Synchronizer = make(map[uint64]*synchronizer.Synchronizer)
		}
		as.Synchronizer[rpcItem.ChainId] = synchronizerTemp
	}
	return nil
}

func (as *Acorus) initBusinessProcessor(cfg config.Config) error {
	for i := range cfg.RPCs {
		log.Println("Init processor success", "chainId", cfg.RPCs[i].ChainId)
		rpcItem := cfg.RPCs[i]
		worker, err := op_stack2.NewWorkerProcessor(as.DB, strconv.FormatUint(cfg.RPCs[i].ChainId, 10), as.shutdown)
		if err != nil {
			return err
		}
		if as.Worker == nil {
			as.Worker = make(map[uint64]*op_stack2.WorkerProcessor)
		}
		as.Worker[rpcItem.ChainId] = worker
	}
	return nil
}

func (as *Acorus) startHttpServer(ctx context.Context, cfg config.Server) error {
	log.Println("starting http server...", "port", cfg.Port)
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	as.apiServer = srv
	log.Println("http server started", "addr", srv.Addr())
	return nil
}

func (as *Acorus) startMetricsServer(ctx context.Context, cfg config.Server) error {
	return nil
}
