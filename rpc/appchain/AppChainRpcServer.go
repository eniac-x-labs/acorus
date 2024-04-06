package appchain

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"strings"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/cornerstone-labs/acorus/appchain/bindings"
	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/rpc/appchain/protobuf/pb"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     string
}

type RpcServer struct {
	*RpcServerConfig
	l1RewardManagers map[uint64]*bindings.L1RewardManager
	l2RewardManagers map[uint64]*bindings.L2RewardManager
	tasks            tasks.Group
	stopped          atomic.Bool
}

func NewRpcServer(grpcCfg *RpcServerConfig, chainRpcCfg []*config.RPC, shutdown context.CancelCauseFunc) (*RpcServer, error) {
	l1RewardManagers := make(map[uint64]*bindings.L1RewardManager)
	l2RewardManagers := make(map[uint64]*bindings.L2RewardManager)
	for i := 0; i < len(chainRpcCfg); i++ {
		if chainRpcCfg[i].ChainId == global_const.EthereumSepoliaChainId ||
			chainRpcCfg[i].ChainId == global_const.EthereumChainId {
			continue
		}
		client, _ := ethclient.Dial(chainRpcCfg[i].RpcUrl)

		l1RewardManagerAddr := chainRpcCfg[i].AppChainAddr.L2.L1RewardManager
		l2RewardManagerAddr := chainRpcCfg[i].AppChainAddr.L2.L2RewardManager
		if l1RewardManagerAddr == "" || l2RewardManagerAddr == "" {
			continue
		}
		L1RewardManager, _ := bindings.NewL1RewardManager(common.HexToAddress(l1RewardManagerAddr), client)
		L2RewardManager, _ := bindings.NewL2RewardManager(common.HexToAddress(l2RewardManagerAddr), client)
		l1RewardManagers[chainRpcCfg[i].ChainId] = L1RewardManager
		l2RewardManagers[chainRpcCfg[i].ChainId] = L2RewardManager
	}

	return &RpcServer{
		RpcServerConfig:  grpcCfg,
		l1RewardManagers: l1RewardManagers,
		l2RewardManagers: l2RewardManagers,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in acorus processor: %w", err))
		}},
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%s", s.GrpcHostname, s.GrpcPort)
		log.Info("start rpc server", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(),
		)
		reflection.Register(gs)
		pb.RegisterAppChainServiceServer(gs, s)

		log.Info("grpc info", "port", s.GrpcPort, "address", listener.Addr().String())
		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC server")
		}
	}(s)
	return nil
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func (s *RpcServer) L1StakerRewardsAmount(ctx context.Context, request *pb.L1StakerRewardsAmountRequest) (*pb.L1StakerRewardsAmountResponse, error) {
	chainId := request.ChainId
	chainIdBig, _ := big.NewInt(0).SetString(chainId, 10)
	chainIdUint := chainIdBig.Uint64()
	// get call instance
	l1RewardManager := s.l1RewardManagers[chainIdUint]
	// create req params
	stakerAddress := request.StakerAddress
	strategies := request.Strategies
	strategiesSplit := strings.Split(strategies, ",")
	strategiesParam := make([]common.Address, 0)
	for _, strategy := range strategiesSplit {
		strategiesParam = append(strategiesParam, common.HexToAddress(strategy))
	}
	l1RewardManagerCaller := l1RewardManager.L1RewardManagerCaller
	rawCaller := &bindings.L1RewardManagerCallerRaw{Contract: &l1RewardManagerCaller}
	var incomeResult []interface{}
	err := rawCaller.Call(&bind.CallOpts{Context: ctx,
		From: common.HexToAddress(stakerAddress),
	}, &incomeResult, "stakerRewardsAmount", strategiesParam)
	if err != nil {
		log.Error("appchain rpc server ", "method", "L1StakerRewardsAmount", "err", err)
		return &pb.L1StakerRewardsAmountResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	log.Info("appchain rpc server ", "method", "L1StakerRewardsAmount", "incomeResult", incomeResult)
	return &pb.L1StakerRewardsAmountResponse{
		Success: true,
		Message: "success",
		Income:  incomeResult[0].(*big.Int).String(),
	}, nil
}

func (s *RpcServer) L2StakerRewardsAmount(ctx context.Context, request *pb.L2StakerRewardsAmountRequest) (*pb.L2StakerRewardsAmountResponse, error) {
	chainId := request.ChainId
	chainIdBig, _ := big.NewInt(0).SetString(chainId, 10)
	chainIdUint := chainIdBig.Uint64()
	// get call instance
	l2RewardManager := s.l2RewardManagers[chainIdUint]
	// create req params
	stakerAddress := request.StakerAddress
	strategy := request.Strategy
	paramAddress := make([]common.Address, 0)
	paramAddress = append(paramAddress, common.HexToAddress(stakerAddress))
	l2RewardManagerCaller := l2RewardManager.L2RewardManagerCaller
	rawCaller := &bindings.L2RewardManagerCallerRaw{Contract: &l2RewardManagerCaller}
	var incomeResult []interface{}
	err := rawCaller.Call(&bind.CallOpts{Context: ctx,
		From: common.HexToAddress(stakerAddress),
	}, &incomeResult, "stakerRewardsAmount",
		common.HexToAddress(strategy),
	)
	if err != nil {
		log.Error("appchain rpc server ", "method", "L2StakerRewardsAmount", "err", err)
		return &pb.L2StakerRewardsAmountResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	log.Info("appchain rpc server ", "method", "L2StakerRewardsAmount", "incomeResult", incomeResult)
	return &pb.L2StakerRewardsAmountResponse{
		Success: true,
		Message: "success",
		Income:  incomeResult[0].(*big.Int).String(),
	}, nil
}
