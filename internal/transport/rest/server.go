package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	readTimeout, writeTimeout = 10 * time.Second, 10 * time.Second
	maxHeaderMegabytes        = 1
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", port),
			Handler:        handler,
			MaxHeaderBytes: maxHeaderMegabytes << 20, // 1 MB
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
