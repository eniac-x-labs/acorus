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
	UpdateFundingPoolBalance(sourceChainId, destChainId, amount,
		receiveAddress, tokenAddress, hash string) (*pb.UpdateFundingPoolBalanceResponse, error)
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
	crossChainTransfer, err := r.bRpcService.CrossChainTransfer(ctx, crossChainReq)
	return crossChainTransfer, err
}

func (r *bridgeRpcService) UpdateFundingPoolBalance(sourceChainId, destChainId, amount,
	receiveAddress, tokenAddress, hash string) (*pb.UpdateFundingPoolBalanceResponse, error) {
	ctx := context.Background()
	updateFundingPoolReq := &pb.UpdateFundingPoolBalanceRequest{
		SourceChainId:  sourceChainId,
		DestChainId:    destChainId,
		Amount:         amount,
		ReceiveAddress: receiveAddress,
		TokenAddress:   tokenAddress,
		SourceHash:     hash,
	}
	poolBalanceResponse, err := r.bRpcService.UpdateFundingPoolBalance(ctx, updateFundingPoolReq)
	return poolBalanceResponse, err
}
