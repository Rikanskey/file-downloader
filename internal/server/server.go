package server

import (
	"file-downloader/internal/config"
	"net/http"
)

type Server struct {
	standardServer *http.Server
}

func New(cfg *config.HttpConfig, handler http.Handler) *Server {
	return &Server{
		standardServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error { return s.standardServer.ListenAndServe() }
