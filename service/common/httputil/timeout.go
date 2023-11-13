package httputil

import "time"

var DefaultTimeouts = HTTPTimeouts{
	ReadTimeout:       30 * time.Second,
	ReadHeaderTimeout: 30 * time.Second,
	WriteTimeout:      30 * time.Second,
	IdleTimeout:       120 * time.Second,
}

type HTTPTimeouts struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func WithTimeouts(timeouts HTTPTimeouts) HTTPOption {
	return func(s *HTTPServer) error {
		s.srv.ReadTimeout = timeouts.ReadTimeout
		s.srv.ReadHeaderTimeout = timeouts.ReadHeaderTimeout
		s.srv.WriteTimeout = timeouts.WriteTimeout
		s.srv.IdleTimeout = timeouts.IdleTimeout
		return nil
	}
}
