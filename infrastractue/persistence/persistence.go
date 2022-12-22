package persistence

import (
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/domain/repository"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/io"
)

type Repositories struct {
	db   *io.SQLdatabase
	User repository.IUserRepository
}

func NewRepositories(db *io.SQLdatabase) (*Repositories, error) {
	return &Repositories{
		db:   db,
		User: NewUserRepository(db),
	}, nil
}
