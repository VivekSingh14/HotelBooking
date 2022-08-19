package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Community/hotel-booking/config"
	"github.com/gorilla/mux"
)

type Server struct {
	config  *config.Config
	router  *mux.Router
	handler http.Handler
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) ListenAndServe(stopch <-chan struct{}) {
	s.registerHandlers()
	s.handler = s.router
	srv := s.startServer()

	<-stopch
	ctx, cancel := context.WithTimeout(context.Background(), s.config.Server.CtxDefaultTimeout)
	defer cancel()

	if srv != nil {
		if err := srv.Shutdown(ctx); err != http.ErrServerClosed {
			log.Fatalf("Http server graceful shutdown failed.")
		}
	}
}

func (s *Server) startServer() *http.Server {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", s.config.Server.Port),
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
		Handler:      s.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Http server crashed")
		}
	}()
	return srv
}
