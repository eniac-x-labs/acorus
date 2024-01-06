package routes

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"net/http"
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
