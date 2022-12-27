package command

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/config"
	graph "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/generated"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/resolver"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/infrastractue/persistence"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/io"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	exitOK   = 0
	existErr = 1
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return existErr
	}
	defer logger.Sync()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return existErr
	}

	// init listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		logger.Error("failed to listen port", zap.String("port", cfg.Port), zap.Error(err))
		return existErr

	}
	logger.Info("server start listening", zap.String("port", cfg.Port))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	db, err := io.NewSQLdatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repository, err := persistence.NewRepositories(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &resolver.Resolver{
					Repo: repository,
				},
			},
		),
	)

	httpServer := server.NewServer(srv, logger, cfg)
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return httpServer.Serve(listener)
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
	case <-ctx.Done():
	}

	if err := httpServer.GracefulShutdown(ctx); err != nil {
		return existErr
	}

	cancel()
	if err := wg.Wait(); err != nil {
		return existErr
	}

	return exitOK
}
