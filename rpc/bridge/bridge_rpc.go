package bridge

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cornerstone-labs/acorus/rpc/bridge/protobuf/pb"
)

type BridgeRpcService interface {
	ChangeTransferStatus(sourceChainId, destChainId, txHash string) (*pb.CrossChainTransferStatusResponse, error)
	CrossChainTransfer(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, fee, nonce, sourceHash string) (*pb.CrossChainTransferResponse, error)
	UpdateWithdrawFundingPoolBalance(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, hash string) (*pb.UpdateWithdrawFundingPoolBalanceResponse, error)
	UpdateDepositFundingPoolBalance(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, hash string) (*pb.UpdateDepositFundingPoolBalanceResponse, error)
	UnstakeBatch(sourceHash, sourceChainId, destChainId string, strategyMap map[string]uint64) (*pb.UnstakeBatchResponse, error)
	MigrateL1Shares(sourceHash, chainId, strategy, staker, shares string, nonce uint64) (*pb.MigrateL1SharesResponse, error)
	BatchMint(batchId uint64, batchMint map[string]string) (*pb.BatchMintResponse, error)
	TransferToL2DappLinkBridge(batchId uint64, ChainId, StrategyAddress string) (*pb.TransferToL2DappLinkBridgeResponse, error)
	TransferL2Share(request *pb.TransferL2ShareRequest) (*pb.TransferL2ShareResponse, error)
}

type bridgeRpcService struct {
	bRpcService pb.BridgeServiceClient
}

func NewBridgeRpcService(rpcUrl string) (BridgeRpcService, error) {
	conn, err := grpc.Dial(rpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("bridge rpc did not connect: ", err)
		return nil, err
	}
	bridgeServiceClient := pb.NewBridgeServiceClient(conn)
	brService := &bridgeRpcService{bridgeServiceClient}
	return brService, nil
}

func (r *bridgeRpcService) ChangeTransferStatus(sourceChainId, destChainId, txHash string) (*pb.CrossChainTransferStatusResponse, error) {
	ctx := context.Background()
	transferRequest := &pb.CrossChainTransferStatusRequest{SourceChainId: sourceChainId,
		DestChainId: destChainId, TxHash: txHash,
	}
	log.Info("ChangeTransferStatusRpc", "transferRequest", transferRequest)
	transferStatus, err := r.bRpcService.ChangeTransferStatus(ctx, transferRequest)
	return transferStatus, err
}

func (r *bridgeRpcService) CrossChainTransfer(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, fee, nonce, sourceHash string) (*pb.CrossChainTransferResponse, error) {
	ctx := context.Background()
	crossChainReq := &pb.CrossChainTransferRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		Fee:            fee,
		Nonce:          nonce,
		SourceHash:     sourceHash,
	}
	log.Info("CrossChainTransferRpc", "crossChainReq", crossChainReq)
	crossChainTransfer, err := r.bRpcService.CrossChainTransfer(ctx, crossChainReq)
	return crossChainTransfer, err
}

func (r *bridgeRpcService) UpdateDepositFundingPoolBalance(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, hash string) (*pb.UpdateDepositFundingPoolBalanceResponse, error) {
	ctx := context.Background()
	updateFundingPoolReq := &pb.UpdateDepositFundingPoolBalanceRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		SourceHash:     hash,
	}
	log.Info("UpdateDepositFundingPoolBalanceRpc", "updateFundingPoolReq", updateFundingPoolReq)
	poolBalanceResponse, err := r.bRpcService.UpdateDepositFundingPoolBalance(ctx, updateFundingPoolReq)
	return poolBalanceResponse, err
}

func (r *bridgeRpcService) UpdateWithdrawFundingPoolBalance(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, hash string) (*pb.UpdateWithdrawFundingPoolBalanceResponse, error) {
	ctx := context.Background()
	updateFundingPoolReq := &pb.UpdateWithdrawFundingPoolBalanceRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		SourceHash:     hash,
	}
	log.Info("UpdateWithdrawFundingPoolBalanceRpc", "updateFundingPoolReq", updateFundingPoolReq)
	poolBalanceResponse, err := r.bRpcService.UpdateWithdrawFundingPoolBalance(ctx, updateFundingPoolReq)
	return poolBalanceResponse, err
}

func (r *bridgeRpcService) UnstakeBatch(sourceHash, sourceChainId, destChainId string, strategyMap map[string]uint64) (*pb.UnstakeBatchResponse, error) {
	ctx := context.Background()
	upstakeBatchReq := &pb.UnstakeBatchRequest{
		SourceChainId: sourceChainId,
		DestChainId:   destChainId,
		SourceHash:    sourceHash,
		Strategy:      strategyMap,
		GasLimit:      "200000",
	}
	log.Info("UnstakeBatchRpc", "upstakeBatchReq", upstakeBatchReq)
	return r.bRpcService.UnstakeBatch(ctx, upstakeBatchReq)
}

func (r *bridgeRpcService) BatchMint(batchId uint64, batchMint map[string]string) (*pb.BatchMintResponse, error) {
	ctx := context.Background()
	batchMintReq := &pb.BatchMintRequest{
		Batch: batchId,
		Mint:  batchMint,
	}
	log.Info("BatchMintRpc", "batchId", batchId, "batchMintReq", batchMintReq)
	return r.bRpcService.BatchMint(ctx, batchMintReq)
}

func (r *bridgeRpcService) TransferToL2DappLinkBridge(batchId uint64, ChainId, StrategyAddress string) (*pb.TransferToL2DappLinkBridgeResponse, error) {
	ctx := context.Background()
	transferToL2DappLinkBridgeReq := &pb.TransferToL2DappLinkBridgeRequest{
		Batch:           batchId,
		ChainId:         ChainId,
		StrategyAddress: StrategyAddress,
	}
	log.Info("TransferToL2DappLinkBridge", "ChainId", ChainId, "batchId", batchId, "transferToL2DappLinkBridgeReq", transferToL2DappLinkBridgeReq)
	return r.bRpcService.TransferToL2DappLinkBridge(ctx, transferToL2DappLinkBridgeReq)
}

func (r *bridgeRpcService) MigrateL1Shares(sourceHash, chainId, strategy, staker, shares string, nonce uint64) (*pb.MigrateL1SharesResponse, error) {
	ctx := context.Background()
	migrateL1SharesReq := &pb.MigrateL1SharesRequest{
		SourceHash:            sourceHash,
		ChainId:               chainId,
		Strategies:            strategy,
		Withdrawer:            staker,
		Shares:                shares,
		L1UnStakeMessageNonce: nonce,
	}
	log.Info("MigrateL1SharesRpc", "migrateL1SharesReq", migrateL1SharesReq)
	return r.bRpcService.MigrateL1Shares(ctx, migrateL1SharesReq)
}

func (r *bridgeRpcService) TransferL2Share(request *pb.TransferL2ShareRequest) (*pb.TransferL2ShareResponse, error) {
	log.Info("TransferL2ShareRpc", "transferL2ShareReq", request)
	ctx := context.Background()
	return r.bRpcService.TransferL2Share(ctx, request)
}
