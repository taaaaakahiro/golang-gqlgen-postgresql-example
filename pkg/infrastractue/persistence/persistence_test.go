package persistence

import (
	"context"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/config"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/io"
)

var (
	userRepo *UserRepository
)

func TestMain(m *testing.M) {
	cfg, _ := config.LoadConfig(context.Background())
	db, _ := io.NewSQLdatabase(cfg)
	userRepo = NewUserRepository(db)

	res := m.Run()

	os.Exit(res)

}
