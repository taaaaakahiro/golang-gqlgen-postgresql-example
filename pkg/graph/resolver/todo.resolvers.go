package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"log"

	graph "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/generated"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/model"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/validation"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	if err := validation.ValidateInputModel(input); err != nil {
		log.Print(err)
		return nil, err
	}

	return nil, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
