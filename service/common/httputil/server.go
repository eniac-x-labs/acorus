package httputil

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
)

type HTTPServer struct {
	listener net.Listener
	srv      *http.Server
	closed   atomic.Bool
}

type HTTPOption func(srv *HTTPServer) error

func StartHTTPServer(addr string, handler http.Handler, opts ...HTTPOption) (*HTTPServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to address %q: %w", addr, err)
	}

	srvCtx, srvCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Handler:           handler,
		ReadTimeout:       DefaultTimeouts.ReadTimeout,
		ReadHeaderTimeout: DefaultTimeouts.ReadHeaderTimeout,
		WriteTimeout:      DefaultTimeouts.WriteTimeout,
		IdleTimeout:       DefaultTimeouts.IdleTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			return srvCtx
		},
	}
	out := &HTTPServer{listener: listener, srv: srv}
	for _, opt := range opts {
		if err := opt(out); err != nil {
			srvCancel()
			return nil, errors.Join(fmt.Errorf("failed to apply HTTP option: %w", err), listener.Close())
		}
	}
	go func() {
		err := out.srv.Serve(listener)
		srvCancel()
		if errors.Is(err, http.ErrServerClosed) {
			out.closed.Store(true)
		} else {
			panic(fmt.Errorf("unexpected serve error: %w", err))
		}
	}()
	return out, nil
}

func (s *HTTPServer) Closed() bool {
	return s.closed.Load()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	if err := s.Shutdown(ctx); err != nil {
		if errors.Is(err, ctx.Err()) {
			return s.Close()
		}
		return err
	}
	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *HTTPServer) Close() error {
	return s.srv.Close()
}

func (s *HTTPServer) Addr() net.Addr {
	return s.listener.Addr()
}

func WithMaxHeaderBytes(max int) HTTPOption {
	return func(srv *HTTPServer) error {
		srv.srv.MaxHeaderBytes = max
		return nil
	}
}
