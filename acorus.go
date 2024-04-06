package acorus

import (
	"context"
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/appchain"
	"github.com/cornerstone-labs/acorus/event/mantle"
	"github.com/cornerstone-labs/acorus/event/okx"
	"github.com/cornerstone-labs/acorus/event/zkfair"
	"github.com/cornerstone-labs/acorus/rpc/airdrop"
	"github.com/cornerstone-labs/acorus/worker/clean_data_worker"
	"github.com/cornerstone-labs/acorus/worker/mantle_worker"
	"github.com/cornerstone-labs/acorus/worker/okx_worker"
	"github.com/cornerstone-labs/acorus/worker/point_worker"
	"github.com/cornerstone-labs/acorus/worker/polygon_worker"
	"github.com/cornerstone-labs/acorus/worker/zkfair_worker"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"net"
	"strconv"
	"sync/atomic"
	"time"

	chi "github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	common3 "github.com/cornerstone-labs/acorus/common"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	event2 "github.com/cornerstone-labs/acorus/event"
	"github.com/cornerstone-labs/acorus/event/base"
	"github.com/cornerstone-labs/acorus/event/linea"
	"github.com/cornerstone-labs/acorus/event/manta"
	"github.com/cornerstone-labs/acorus/event/op_stack"
	"github.com/cornerstone-labs/acorus/event/polygon"
	"github.com/cornerstone-labs/acorus/event/scroll"
	"github.com/cornerstone-labs/acorus/relayer"
	rpc_appchain "github.com/cornerstone-labs/acorus/rpc/appchain"
	"github.com/cornerstone-labs/acorus/rpc/bridge"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	worker2 "github.com/cornerstone-labs/acorus/worker"
	"github.com/cornerstone-labs/acorus/worker/base_worker"
	"github.com/cornerstone-labs/acorus/worker/manta_worker"
	op_stack2 "github.com/cornerstone-labs/acorus/worker/op-stack"
)

type Acorus struct {
	DB                     *database.DB
	ethClient              map[uint64]node.EthClient
	apiServer              *httputil.HTTPServer
	metricsServer          *httputil.HTTPServer
	metricsRegistry        *prometheus.Registry
	Synchronizer           map[uint64]*synchronizer.Synchronizer
	Processor              map[uint64]event2.IEventProcessor
	Worker                 map[uint64]worker2.IWorkerProcessor
	Relayer                map[uint64]*relayer.RelayerListener
	shutdown               context.CancelCauseFunc
	stopped                atomic.Bool
	chainIdList            []uint64
	birdgeRpcService       bridge.BridgeRpcService
	relayerBridgeRelation  *relayer.RelayerBridgeRelation
	relayerFundingPoolTask *relayer.RelayerFundingPool
	airDropWorker          *point_worker.PointWorker
	cleanDataWorker        *clean_data_worker.WorkerProcessor
	appChainL1             *appchain.L1AppChainListener
	appChainRpcServer      *rpc_appchain.RpcServer
}

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

const MaxRecvMessageSize = 1024 * 1024 * 300

func NewAcorus(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*Acorus, error) {
	log.Info("New acorus start️ 🕖")
	out := &Acorus{
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	log.Info("New acorus success🏅️")
	return out, nil
}

func (as *Acorus) Start(ctx context.Context) error {
	for i := range as.chainIdList {
		log.Info("starting Sync", "chainId", as.chainIdList[i])
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
		relayerRunner := as.Relayer[realChainId]
		if relayerRunner != nil {
			if err := relayerRunner.Start(); err != nil {
				return fmt.Errorf("failed to start relayer: %w", err)
			}
		}
	}
	if err := as.relayerBridgeRelation.Start(); err != nil {
		return fmt.Errorf("failed to start relayerBridgeRelation: %w", err)
	}
	if err := as.relayerFundingPoolTask.Start(); err != nil {
		return fmt.Errorf("failed to start relayerFundingPool: %w", err)
	}
	if err := as.airDropWorker.Start(); err != nil {
		log.Error("start airdrop worker failed", "err", err)
		return err
	}
	if err := as.cleanDataWorker.WorkerStart(); err != nil {
		log.Error("start clean data worker failed", "err", err)
		return err
	}
	if err := as.appChainL1.Start(); err != nil {
		log.Error("start appChainL1 failed", "err", err)
		return err
	}
	if err := as.appChainRpcServer.Start(ctx); err != nil {
		log.Error("start appChainRpcServer failed", "err", err)
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

	log.Info("acorus stopped")

	return result
}

func (as *Acorus) Stopped() bool {
	return as.stopped.Load()
}

func (as *Acorus) initFromConfig(ctx context.Context, cfg *config.Config) error {
	common3.ConfigInfo = cfg
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
	if err := as.initFundingPool(cfg); err != nil {
		fmt.Errorf("failed to init funding pool: %w", err)
	}
	if err := as.initAirDropWorker(cfg); err != nil {
		fmt.Errorf("failed to init tAirDropWorker: %w", err)
	}
	if err := as.initCleanDataWorker(cfg); err != nil {
		fmt.Errorf("failed to init clean data worker: %w", err)
	}
	if err := as.initAppChainL1(cfg); err != nil {
		fmt.Errorf("failed to initAppChainL1: %w", err)
	}
	if err := as.initAppChainRpcServer(cfg); err != nil {
		fmt.Errorf("failed to initAppChainRpcServer: %w", err)
	}
	return nil
}

func (as *Acorus) initAppChainL1(cfg *config.Config) error {
	chainId := cfg.AppChain.L1.ChainId
	contracts := cfg.AppChain.L1.Contracts
	startHeight := cfg.AppChain.L1.StartHeight
	var loopInterval time.Duration = time.Second * 5
	var epoch uint64 = 10_000
	processor, err := appchain.NewL1AppChainListener(chainId, contracts, as.shutdown, loopInterval, startHeight, epoch, as.birdgeRpcService, as.DB)
	if err != nil {
		log.Error("init initAppChainL1 failed", "err", err)
	}
	as.appChainL1 = processor
	return nil
}

func (as *Acorus) initCleanDataWorker(cfg *config.Config) error {
	processor, err := clean_data_worker.NewWorkerProcessor(as.DB, *cfg, as.shutdown)
	if err != nil {
		log.Error("init clean_data_worker failed", "err", err)
	}
	as.cleanDataWorker = processor
	return nil
}

func (as *Acorus) initAirDropWorker(cfg *config.Config) error {
	airDropRpcService, err := airdrop.NewAirDropRpcService(cfg.AirDropGRpcUrl)
	if err != nil {
		log.Error("init airdrop rpc service failed", "err", err)
		return err
	} else {
		airDropWorker, err := point_worker.NewPointWorker(as.DB, airDropRpcService)
		if err != nil {
			log.Error("init airdrop worker failed", "err", err)
			return err
		} else {
			as.airDropWorker = airDropWorker
		}
	}

	return nil
}

func (as *Acorus) initRPCClients(ctx context.Context, conf *config.Config) error {
	for i := range conf.RPCs {
		log.Info("Init rpc client", "ChainId", conf.RPCs[i].ChainId, "RpcUrl", conf.RPCs[i].RpcUrl)
		rpc := conf.RPCs[i]
		ethClient, err := node.DialEthClient(ctx, rpc.RpcUrl)
		if err != nil {
			log.Error("dial eth client fail", "err", err)
			return fmt.Errorf("Failed to dial L1 client: %w", err)
		}
		if as.ethClient == nil {
			as.ethClient = make(map[uint64]node.EthClient)
		}
		as.ethClient[rpc.ChainId] = ethClient
		as.chainIdList = append(as.chainIdList, rpc.ChainId)
	}
	log.Info("Init rpc client success")
	return nil
}

func (as *Acorus) initDB(ctx context.Context, cfg config.Database) error {

	db, err := database.NewDB(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	as.DB = db
	log.Info("Init database success")
	return nil
}

func (as *Acorus) initEventProcessor(cfg *config.Config) error {
	var loopInterval time.Duration = time.Second * 5
	var epoch uint64 = 10_000
	var l1StartBlockNumber *big.Int
	for i := range cfg.RPCs {
		log.Info("init chain processor", "chainId", cfg.RPCs[i].ChainId)
		log.Info("Init processor success", "chainId", cfg.RPCs[i].ChainId)
		if as.Processor == nil {
			as.Processor = make(map[uint64]event2.IEventProcessor)
		}
		if as.Worker == nil {
			as.Worker = make(map[uint64]worker2.IWorkerProcessor)
		}
		rpcItem := cfg.RPCs[i]
		chainId := rpcItem.ChainId
		isMainnet := rpcItem.IsMainnet
		l1ChainIdStr := strconv.FormatUint(rpcItem.L1ChainId, 10)
		l2ChainIdStr := strconv.FormatUint(chainId, 10)
		var processor event2.IEventProcessor
		var worker worker2.IWorkerProcessor
		var err error
		if chainId == global_const.EthereumChainId ||
			chainId == global_const.EthereumSepoliaChainId {
			l1StartBlockNumber = big.NewInt(int64(rpcItem.StartBlock))
		}
		if chainId == global_const.ScrollChainId {
			processor, err = scroll.NewBridgeProcessor(as.DB, rpcItem, as.shutdown,
				loopInterval, epoch, l1ChainIdStr, l2ChainIdStr)
			if err != nil {
				return err
			}
		} else if chainId == global_const.PolygonChainId ||
			chainId == global_const.PolygonSepoliaChainId {
			if worker, err = polygon_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			processor, err = polygon.NewBridgeProcessor(as.DB, rpcItem, as.shutdown,
				loopInterval, epoch, l1ChainIdStr, l2ChainIdStr)
			if err != nil {
				return err
			}
		} else if chainId == global_const.OpChinId ||
			chainId == global_const.OpTestChinId {
			if worker, err = op_stack2.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = op_stack.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
				l1ChainIdStr,
				l2ChainIdStr,
				isMainnet,
			); err != nil {
				return err
			}
		} else if chainId == global_const.LineaChainId {
			processor, err = linea.NewBridgeProcessor(as.DB, rpcItem, as.shutdown,
				loopInterval, epoch, l1ChainIdStr, l2ChainIdStr)
			if err != nil {
				return err
			}
		} else if chainId == global_const.BaseChainId ||
			chainId == global_const.BaseSepoliaChainId {
			if worker, err = base_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = base.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
				l1ChainIdStr,
				l2ChainIdStr,
				isMainnet,
			); err != nil {
				return err
			}
		} else if chainId == global_const.MantaChainId ||
			chainId == global_const.MantaSepoliaChainId {
			if worker, err = manta_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = manta.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
				l1ChainIdStr,
				l2ChainIdStr,
				isMainnet,
			); err != nil {
				return err
			}
		} else if chainId == global_const.MantleChainId ||
			chainId == global_const.MantleSepoliaChainId {
			if worker, err = mantle_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = mantle.NewBridgeProcessor(
				as.DB,
				l1StartBlockNumber,
				big.NewInt(int64(rpcItem.StartBlock)),
				as.shutdown,
				loopInterval,
				epoch,
				l1ChainIdStr,
				l2ChainIdStr,
				isMainnet,
			); err != nil {
				return err
			}
		} else if chainId == global_const.ZkFairSepoliaChainId ||
			chainId == global_const.ZkFairChainId {
			if worker, err = zkfair_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = zkfair.NewBridgeProcessor(as.DB, rpcItem, as.shutdown,
				loopInterval, epoch, l1ChainIdStr, l2ChainIdStr); err != nil {
				return err
			}
		} else if chainId == global_const.OkxSepoliaChainId ||
			chainId == global_const.OkxChainId {
			if worker, err = okx_worker.NewWorkerProcessor(as.DB, l2ChainIdStr, as.shutdown); err != nil {
				return err
			}
			if processor, err = okx.NewBridgeProcessor(as.DB, rpcItem, as.shutdown,
				loopInterval, epoch, l1ChainIdStr, l2ChainIdStr); err != nil {
				return err
			}
		}

		if processor != nil {
			as.Processor[chainId] = processor
		}
		if worker != nil {
			as.Worker[chainId] = worker
		}
	}
	return nil
}

func (as *Acorus) initSynchronizer(config *config.Config) error {
	for i := range config.RPCs {
		log.Info("Init synchronizer success", "chainId", config.RPCs[i].ChainId)
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
		log.Info("Init Relayer success", "chainId", chainIdStr)
		if chainIdStr == "1" || chainIdStr == "11155111" ||
			chainIdStr == "5" {
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
		log.Error("initRelayerRelation failed", "err", err)
		return err
	} else {
		as.relayerBridgeRelation = relayerBridgeRelation
	}
	return nil
}

func (as *Acorus) initAppChainRpcServer(cfg *config.Config) error {
	grpcCfg := &rpc_appchain.RpcServerConfig{
		GrpcHostname: cfg.Server.Host,
		GrpcPort:     strconv.Itoa(cfg.Server.GrpcPort),
	}
	server, err := rpc_appchain.NewRpcServer(grpcCfg, cfg.RPCs, as.shutdown)
	if err != nil {
		log.Error("initAppChainRpcServer failed", "err", err)
		return err
	}
	as.appChainRpcServer = server
	return nil
}

func (as *Acorus) initFundingPool(cfg *config.Config) error {
	fundingPool, err := relayer.NewRelayerFundingPool(as.DB, as.birdgeRpcService, cfg.RPCs, as.shutdown)
	if err != nil {
		log.Error("initFundingPool failed", "err", err)
		return err
	} else {
		as.relayerFundingPoolTask = fundingPool
	}
	return nil
}

func (as *Acorus) initBridgeRpcService(cfg config.Config) error {
	service, err := bridge.NewBridgeRpcService(cfg.BridgeGrpcUrl)
	if err != nil {
		log.Error("init bridge rpc service failed", "err", err)
		return err
	} else {
		as.birdgeRpcService = service
	}
	return nil
}

func (as *Acorus) startHttpServer(ctx context.Context, cfg config.Server) error {
	log.Info("starting http server...", "port", cfg.Port)
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	as.apiServer = srv
	log.Info("http server started", "addr", srv.Addr())
	return nil
}
