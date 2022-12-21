package graph

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/io"
)

type Resolver struct {
	DB *io.SQLdatabase
}
