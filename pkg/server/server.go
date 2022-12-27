package server

import (
	"context"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	Router  *mux.Router
	server  *http.Server
	handler *handler.Server
	log     *zap.Logger
}

func NewServer(handler *handler.Server, logger *zap.Logger, cfg *config.Config) *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		handler: handler,
		log:     logger,
	}

	s.registerHandler(cfg)

	return s
}

func (s *Server) Serve(listener net.Listener) error {
	server := &http.Server{
		Handler: s.Router,
	}
	s.server = server
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandler(env *config.Config) {

	// GraphQL PlayGround
	s.Router.Handle("/query", s.handler)
	s.Router.HandleFunc("/", playground.Handler("GraphQL playground", "/query")).Methods(http.MethodGet, http.MethodOptions)
	s.Router.HandleFunc("/healthz", s.healthCheckHandler).Methods(http.MethodGet)
}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
