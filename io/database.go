package io

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/config"
)

type SQLdatabase struct {
	Database *sqlx.DB
}

func NewSQLdatabase(cfg *config.Config) (*SQLdatabase, error) {
	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s  host=localhost port=5432 sslmode=disable",
		cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &SQLdatabase{Database: db}, nil
}