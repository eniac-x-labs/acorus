package httputil

import (
	"net/http"

	"github.com/ethereum/go-ethereum/rpc"
)

var timeouts = rpc.DefaultHTTPTimeouts

func NewHttpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
	}
}
