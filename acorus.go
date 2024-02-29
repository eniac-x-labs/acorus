package acorus

import (
	"context"
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
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
	"github.com/cornerstone-labs/acorus/event/linea"
	"github.com/cornerstone-labs/acorus/event/op_stack"
	"github.com/cornerstone-labs/acorus/event/polygon"
	"github.com/cornerstone-labs/acorus/event/scroll"
	"github.com/cornerstone-labs/acorus/relayer"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	worker2 "github.com/cornerstone-labs/acorus/worker"
	op_stack2 "github.com/cornerstone-labs/acorus/worker/op-stack"
)

type Acorus struct {
	DB                    *database.DB
	ethClient             map[uint64]node.EthClient
	apiServer             *httputil.HTTPServer
	metricsServer         *httputil.HTTPServer
	metricsRegistry       *prometheus.Registry
	Synchronizer          map[uint64]*synchronizer.Synchronizer
	Processor             map[uint64]event2.IEventProcessor
	Worker                map[uint64]worker2.IWorkerProcessor
	Relayer               map[uint64]*relayer.RelayerListener
	shutdown              context.CancelCauseFunc
	stopped               atomic.Bool
	chainIdList           []uint64
	birdgeRpcService      bridge.BridgeRpcService
	relayerBridgeRelation *relayer.RelayerBridgeRelation
}

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

const MaxRecvMessageSize = 1024 * 1024 * 300

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
			if err := worker.WorkerStart(); err != nil {
				return fmt.Errorf("failed to start worker: %w", err)
			}
		}
		relayer := as.Relayer[realChainId]
		if relayer != nil {
			if err := relayer.Start(); err != nil {
				return fmt.Errorf("failed to start relayer: %w", err)
			}
		}
	}
	if err := as.relayerBridgeRelation.Start(); err != nil {
		return fmt.Errorf("failed to start relayerBridgeRelation: %w", err)
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
	if err := as.initBridgeRpcService(*cfg); err != nil {
		return fmt.Errorf("failed to init bridge rpc service: %w", err)
	}
	if err := as.initSynchronizer(cfg); err != nil {
		return fmt.Errorf("failed to init L1 Sync: %w", err)
	}
	if err := as.initEventProcessor(cfg); err != nil {
		return fmt.Errorf("failed to init Processor: %w", err)
	}
	if err := as.initRelayer(ctx, cfg); err != nil {
		return fmt.Errorf("failed to init relayer: %w", err)
	}
	if err := as.initRelayerRelation(); err != nil {
		fmt.Errorf("failed to init relayer relation: %w", err)
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
	for i := range cfg.RPCs {
		log.Println("init chain processor", "chainId", cfg.RPCs[i].ChainId)
		log.Println("Init processor success", "chainId", cfg.RPCs[i].ChainId)
		if as.Processor == nil {
			as.Processor = make(map[uint64]event2.IEventProcessor)
		}
		if as.Worker == nil {
			as.Worker = make(map[uint64]worker2.IWorkerProcessor)
		}
		rpcItem := cfg.RPCs[i]
		var processor event2.IEventProcessor
		var worker worker2.IWorkerProcessor
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
			processor, err = polygon.NewBridgeProcessor(as.DB, rpcItem, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		} else if rpcItem.ChainId == global_const.OpChinId {
			if worker, err = op_stack2.NewWorkerProcessor(as.DB, strconv.FormatUint(cfg.RPCs[i].ChainId, 10), as.shutdown); err != nil {
				return err
			}
			if processor, err = op_stack.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
			); err != nil {
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
		if worker != nil {
			as.Worker[rpcItem.ChainId] = worker
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
			HeaderBufferSize:  uint(rpcItem.HeaderBufferSize),
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

func (as *Acorus) initRelayer(ctx context.Context, cfg *config.Config) error {
	var loopInterval time.Duration = time.Second * 5
	var epoch uint64 = 10_000
	for i := range cfg.Relayers {
		relayerCfg := cfg.Relayers[i]
		chainIdStr := relayerCfg.ChainId
		chainId, uerr := strconv.ParseUint(chainIdStr, 10, 64)
		if uerr != nil {
			return uerr
		}
		var rlworker *relayer.RelayerListener
		var err error
		log.Println("Init Relayer success", "chainId", chainIdStr)
		if chainIdStr == "1" || chainIdStr == "11155111" {
			rlworker, err = relayer.NewRelayerListener(as.DB, chainIdStr, relayerCfg.Contracts, relayerCfg.Contracts,
				relayerCfg.EventStartBlock, relayerCfg.EventStartBlock, 1, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		} else {
			rlworker, err = relayer.NewRelayerListener(as.DB, chainIdStr, relayerCfg.Contracts, relayerCfg.Contracts,
				relayerCfg.EventStartBlock, relayerCfg.EventStartBlock, 2, as.shutdown, loopInterval, epoch)
			if err != nil {
				return err
			}
		}

		if as.Relayer == nil {
			as.Relayer = make(map[uint64]*relayer.RelayerListener)
		}

		as.Relayer[chainId] = rlworker
	}
	return nil
}

func (as *Acorus) initRelayerRelation() error {
	relayerBridgeRelation, err := relayer.NewRelayerBridgeRelation(as.DB, as.birdgeRpcService, as.shutdown)
	if err != nil {
		log.Println("initRelayerRelation failed", "err", err)
		return err
	} else {
		as.relayerBridgeRelation = relayerBridgeRelation
	}
	return nil
}

func (as *Acorus) initBridgeRpcService(cfg config.Config) error {
	service, err := bridge.NewBridgeRpcService(cfg.BridgeGrpcUrl)
	if err != nil {
		log.Println("init bridge rpc service failed", "err", err)
		return err
	} else {
		as.birdgeRpcService = service
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
