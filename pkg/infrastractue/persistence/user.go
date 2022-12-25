package persistence

import (
	"context"
	"database/sql"

	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/io"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/domain/entity"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/domain/repository"
)

type UserRepository struct {
	database *io.SQLdatabase
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *io.SQLdatabase) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	err := r.database.SQLX.Get(&user, "SELECT id, name FROM users WHERE id=$1", userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return &user, nil
}
