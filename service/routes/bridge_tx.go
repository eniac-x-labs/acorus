package routes

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/log"
)

// L1ToL2ListHandler ... Handles /api/v1/deposits GET requests
func (h Routes) L1ToL2ListHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")
	chainId := r.URL.Query().Get("chainId")

	params, err := h.svc.QueryDWParams(chainId, address, pageQuery, pageSizeQuery, order)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	cacheKey := fmt.Sprintf("l1ToL2List{chainId:%s,address:%s,page:%s,pageSize:%s,order:%s}", chainId, address, pageQuery, pageSizeQuery, order)
	if h.enableCache {
		response, _ := h.cache.GetL1ToL2List(cacheKey)
		if response != nil {
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	l1ToL2Txs, err := h.svc.GetDepositList(params)
	if h.enableCache {
		h.cache.AddL1ToL2List(cacheKey, l1ToL2Txs)
	}

	err = jsonResponse(w, l1ToL2Txs, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// L2ToL1ListHandler ... Handles /api/v1/withdrawals GET requests
func (h Routes) L2ToL1ListHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")
	chainId := r.URL.Query().Get("chainId")

	params, err := h.svc.QueryDWParams(chainId, address, pageQuery, pageSizeQuery, order)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	cacheKey := fmt.Sprintf("l1ToL2List{chainId:%s,address:%s,page:%s,pageSize:%s,order:%s}", chainId, address, pageQuery, pageSizeQuery, order)
	if h.enableCache {
		response, _ := h.cache.GetL2ToL1List(cacheKey)
		if response != nil {
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	l2ToL1Txs, err := h.svc.GetWithdrawalList(params)
	if err != nil {
		http.Error(w, "Internal server error reading l2tol1 list", http.StatusInternalServerError)
		log.Error("Unable to read l2tol1 list from DB", "err", err.Error())
		return
	}
	if h.enableCache {
		h.cache.AddL2ToL1List(cacheKey, l2ToL1Txs)
	}

	err = jsonResponse(w, l2ToL1Txs, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// StakingRecordsHandler ... Handles /api/v1/staking-records GET requests
func (h Routes) StakingRecordsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	txType := r.URL.Query().Get("txType")
	order := r.URL.Query().Get("order")

	params, err := h.svc.QuerySRParams(address, pageQuery, pageSizeQuery, order, txType)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	cacheKey := fmt.Sprintf("StakingRecords{address:%s,page:%s,pageSize:%s,order:%s}", address, pageQuery, pageSizeQuery, order)
	if h.enableCache {
		response, _ := h.cache.GetStakingRecords(cacheKey)
		if response != nil {
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	stakingRecords, err := h.svc.GetStakingRecords(params)
	if err != nil {
		http.Error(w, "Internal server error reading staking records", http.StatusInternalServerError)
		log.Error("Unable to read staking records from DB", "err", err.Error())
		return
	}
	if h.enableCache {
		h.cache.AddStakingRecords(cacheKey, stakingRecords)
	}

	err = jsonResponse(w, stakingRecords, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

// BridgeRecordsHandler ... Handles /api/v1/bridge-records GET requests
func (h Routes) BridgeRecordsHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")
	order := r.URL.Query().Get("order")

	params, err := h.svc.QueryBRParams(address, pageQuery, pageSizeQuery, order)
	if err != nil {
		http.Error(w, "invalid query params", http.StatusBadRequest)
		log.Error("error reading request params", "err", err.Error())
		return
	}

	cacheKey := fmt.Sprintf("BridgeRecords{address:%s,page:%s,pageSize:%s,order:%s}", address, pageQuery, pageSizeQuery, order)
	if h.enableCache {
		response, _ := h.cache.GetBridgeRecords(cacheKey)
		if response != nil {
			err = jsonResponse(w, response, http.StatusOK)
			if err != nil {
				log.Error("Error writing response", "err", err.Error())
			}
			return
		}
	}

	bridgeRecords, err := h.svc.GetBridgeRecords(params)
	if err != nil {
		http.Error(w, "Internal server error reading bridge records", http.StatusInternalServerError)
		log.Error("Unable to read bridge records from DB", "err", err.Error())
		return
	}
	if h.enableCache {
		h.cache.AddBridgeRecords(cacheKey, bridgeRecords)
	}

	err = jsonResponse(w, bridgeRecords, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

func (h Routes) BridgeValidHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	bridgeValid := h.svc.BridgeValid(address)
	err := jsonResponse(w, bridgeValid, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}

func (h Routes) StakingValidHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	stakingValid := h.svc.StakingValid(address)
	err := jsonResponse(w, stakingValid, http.StatusOK)
	if err != nil {
		log.Error("Error writing response", "err", err.Error())
	}
}
