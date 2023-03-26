package http

import (
	"context"
	"net/http"
	"time"

	config "github.com/lefes/discord-message-scheduler/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:              cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler:           handler,
		ReadHeaderTimeout: time.Duration(cfg.HTTP.ReadHeaderTimeout) * time.Second,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
