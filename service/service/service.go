package service

import (
	"github.com/ethereum/go-ethereum/log"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/service/models"
)

type Service interface {
	GetDepositList(*models.QueryDWParams) (*models.DepositsResponse, error)
	GetWithdrawalList(params *models.QueryDWParams) (*models.WithdrawsResponse, error)

	QueryDWParams(chainId string, address string, page string, pageSize string, order string) (*models.QueryDWParams, error)
}

type HandlerSvc struct {
	v          *Validator
	l1ToL2View worker.L1ToL2View
	l2ToL1View worker.L2ToL1View
}

func New(v *Validator, l1l2v worker.L1ToL2View, l2l1v worker.L2ToL1View) Service {
	return &HandlerSvc{
		v:          v,
		l1ToL2View: l1l2v,
		l2ToL1View: l2l1v,
	}
}

func (h HandlerSvc) GetDepositList(params *models.QueryDWParams) (*models.DepositsResponse, error) {
	addressToLower := strings.ToLower(params.Address)
	l1L2List, total := h.l1ToL2View.L1ToL2List(params.ChainId, addressToLower, params.Page, params.PageSize, params.Order)
	return &models.DepositsResponse{
		Current: params.Page,
		Size:    params.PageSize,
		Total:   total,
		Records: l1L2List,
	}, nil
}

func (h HandlerSvc) GetWithdrawalList(params *models.QueryDWParams) (*models.WithdrawsResponse, error) {
	addressToLower := strings.ToLower(params.Address)
	l2L1List, total := h.l2ToL1View.L2ToL1List(params.ChainId, addressToLower, params.Page, params.PageSize, params.Order)
	return &models.WithdrawsResponse{
		Current: params.Page,
		Size:    params.PageSize,
		Total:   total,
		Records: l2L1List,
	}, nil
}

func (h HandlerSvc) QueryDWParams(chainId string, address string, page string, pageSize string, order string) (*models.QueryDWParams, error) {
	var paraAddress string
	if address == "0x00" {
		paraAddress = "0x00"
	} else {
		addr, err := h.v.ParseValidateAddress(address)
		if err != nil {
			log.Error("invalid address param", "address", address, "err", err)
			return nil, err
		}
		paraAddress = addr.String()
	}
	pageVal, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New("page must be an integer value")
	}
	err = h.v.ValidatePage(pageVal)
	if err != nil {
		log.Error("invalid page param", "page", page, "err", err)
		return nil, err
	}

	pageSizeVal, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.New("pageSize must be an integer value")
	}
	err = h.v.ValidatePageSize(pageSizeVal)
	if err != nil {
		log.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	err = h.v.ValidateOrder(order)
	if err != nil {
		log.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}
	return &models.QueryDWParams{
		ChainId:  chainId,
		Address:  paraAddress,
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    order,
	}, nil
}
