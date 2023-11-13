package routes

import (
	"net/http"
)

func (h Routes) L1ToL2ListHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")

	page, err := h.v.ParseValidatePage(pageQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error("Invalid query params", "err", err)
		return
	}
	pageSize, err := h.v.ParseValidateLimit(pageSizeQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error("Invalid query params", "err", err)
		return
	}

	l1ToL2Txs, err := h.l1ToL2View.L1ToL2List(address, page, pageSize)
	if err != nil {
		http.Error(w, "Internal server error reading l1tol2 list", http.StatusInternalServerError)
		h.logger.Error("Unable to read l1tol2 list from DB", "err", err.Error())
		return
	}

	err = jsonResponse(w, l1ToL2Txs, http.StatusOK)
	if err != nil {
		h.logger.Error("Error writing response", "err", err.Error())
	}
}

func (h Routes) L2ToL1ListHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	pageQuery := r.URL.Query().Get("page")
	pageSizeQuery := r.URL.Query().Get("pageSize")

	page, err := h.v.ParseValidatePage(pageQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error("Invalid query params", "err", err)
		return
	}
	pageSize, err := h.v.ParseValidateLimit(pageSizeQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Error("Invalid query params", "err", err)
		return
	}

	l2ToL1Txs, err := h.l2ToL1View.L2ToL1List(address, page, pageSize)
	if err != nil {
		http.Error(w, "Internal server error reading l2tol1 list", http.StatusInternalServerError)
		h.logger.Error("Unable to read l2tol1 list from DB", "err", err.Error())
		return
	}

	err = jsonResponse(w, l2ToL1Txs, http.StatusOK)
	if err != nil {
		h.logger.Error("Error writing response", "err", err.Error())
	}
}
