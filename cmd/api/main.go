package main

import (
	"context"
	"log"
	"net/http"

	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/config"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/infrastractue/persistence"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/generated"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/resolver"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/io"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

	// server
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}