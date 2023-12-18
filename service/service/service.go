package service

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/cornerstone-labs/acorus/database/common"
	"github.com/cornerstone-labs/acorus/database/worker"
	"github.com/cornerstone-labs/acorus/service/models"
	"github.com/ethereum/go-ethereum/log"
)

type Service interface {
	GetDepositList(*models.QueryDWParams) (*gplus.Page[worker.L1ToL2], error)
	GetWithdrawalList(params *models.QueryDWParams) (*gplus.Page[worker.L2ToL1], error)
	GetStateRootList(*models.QueryPageListParams) (*gplus.Page[worker.StateRoot], error)
	GetStateRootByIndex(*models.QueryByIndexParams) (*worker.StateRoot, error)

	QueryDWParams(address string, page string, pageSize string, order string) (*models.QueryDWParams, error)
	QueryPageListParams(page string, pageSize string, order string) (*models.QueryPageListParams, error)
	QueryByIdParams(id string) (*models.QueryByIdParams, error)
	QueryByIndexParams(index string) (*models.QueryByIndexParams, error)
}

type HandlerSvc struct {
	logger        log.Logger
	v             *Validator
	l1ToL2View    worker.L1ToL2View
	l2ToL1View    worker.L2ToL1View
	stateRootView worker.StateRootView
	blocksView    common.BlocksView
}

func New(v *Validator, l1l2v worker.L1ToL2View, l2l1v worker.L2ToL1View, blv common.BlocksView, srv worker.StateRootView, l log.Logger) Service {
	return &HandlerSvc{
		logger:        l,
		v:             v,
		l1ToL2View:    l1l2v,
		l2ToL1View:    l2l1v,
		stateRootView: srv,
		blocksView:    blv,
	}
}

func (h HandlerSvc) GetDepositList(params *models.QueryDWParams) (*gplus.Page[worker.L1ToL2], error) {
	addressToLower := strings.ToLower(params.Address)
	return h.l1ToL2View.L1ToL2List(addressToLower, params.Page, params.PageSize, params.Order)
}

func (h HandlerSvc) GetWithdrawalList(params *models.QueryDWParams) (*gplus.Page[worker.L2ToL1], error) {
	addressToLower := strings.ToLower(params.Address)
	return h.l2ToL1View.L2ToL1List(addressToLower, params.Page, params.PageSize, params.Order)
}

func (h HandlerSvc) GetStateRootList(params *models.QueryPageListParams) (*gplus.Page[worker.StateRoot], error) {
	return h.stateRootView.StateRootList(params.Page, params.PageSize, params.Order)
}

func (h HandlerSvc) GetStateRootByIndex(params *models.QueryByIndexParams) (*worker.StateRoot, error) {
	return h.stateRootView.StateRootByIndex(big.NewInt(int64(params.Index)))
}

func (h HandlerSvc) QueryDWParams(address string, page string, pageSize string, order string) (*models.QueryDWParams, error) {
	var paraAddress string
	if address == "0x00" {
		paraAddress = "0x00"
	} else {
		addr, err := h.v.ParseValidateAddress(address)
		if err != nil {
			h.logger.Error("invalid address param", "address", address, "err", err)
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
		h.logger.Error("invalid page param", "page", page, "err", err)
		return nil, err
	}

	pageSizeVal, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.New("pageSize must be an integer value")
	}
	err = h.v.ValidatePageSize(pageSizeVal)
	if err != nil {
		h.logger.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	err = h.v.ValidateOrder(order)
	if err != nil {
		h.logger.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	return &models.QueryDWParams{
		Address:  paraAddress,
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    order,
	}, nil
}

func (h HandlerSvc) QueryPageListParams(page string, pageSize string, order string) (*models.QueryPageListParams, error) {
	pageVal, err := strconv.Atoi(page)
	if err != nil {
		return nil, errors.New("page must be an integer value")
	}
	err = h.v.ValidatePage(pageVal)
	if err != nil {
		h.logger.Error("invalid page param", "page", page, "err", err)
		return nil, err
	}

	pageSizeVal, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.New("pageSize must be an integer value")
	}
	err = h.v.ValidatePageSize(pageSizeVal)
	if err != nil {
		h.logger.Error("invalid query param", "pageSize", pageSize, "err", err)
		return nil, err
	}

	err = h.v.ValidateOrder(order)
	if err != nil {
		h.logger.Error("invalid query param", "order", order, "err", err)
		return nil, err
	}

	return &models.QueryPageListParams{
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    order,
	}, nil
}

func (h HandlerSvc) QueryByIdParams(id string) (*models.QueryByIdParams, error) {
	idValue, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("id must be an integer value")
	}
	err = h.v.ValidateIdOrIndex(idValue)
	if err != nil {
		h.logger.Error("invalid query param", "id", id, "err", err)
		return nil, err
	}
	return &models.QueryByIdParams{
		Id: idValue,
	}, nil
}

func (h HandlerSvc) QueryByIndexParams(index string) (*models.QueryByIndexParams, error) {
	indexValue, err := strconv.ParseUint(index, 10, 64)
	if err != nil {
		return nil, errors.New("index must be an integer value")
	}
	err = h.v.ValidateIdOrIndex(indexValue)
	if err != nil {
		h.logger.Error("invalid query param", "index", index, "err", err)
		return nil, err
	}
	return &models.QueryByIndexParams{
		Index: indexValue,
	}, nil
}
