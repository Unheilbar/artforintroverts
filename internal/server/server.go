package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router chi.Router

	httpServer *http.Server

	addr string
}

func NewServer(router chi.Router, addr string) *Server {
	return &Server{
		router: router,
		addr:   addr,
	}
}

func (s *Server) Start() chan error {
	serverErr := make(chan error)

	s.httpServer = &http.Server{
		Addr:    s.addr,
		Handler: s.router,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	return serverErr
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}
