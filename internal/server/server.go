package server

import (
	"AuthApp/internal/config"
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func Init(config config.Http, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + config.Port,
			Handler:        handler,
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
			MaxHeaderBytes: config.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
