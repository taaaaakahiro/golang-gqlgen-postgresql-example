package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/config"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/infrastractue/persistence"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/customhook"
	graph "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/generated"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/resolver"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/io"

	gqlCfg "github.com/99designs/gqlgen/codegen/config"
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

	// validation
	gqlCfg, err := gqlCfg.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		FieldHook: customhook.ValidationFieldHook,
	}

	err = api.Generate(gqlCfg, api.ReplacePlugin(&p))

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	// server
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
