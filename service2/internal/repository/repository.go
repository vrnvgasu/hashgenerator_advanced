package repository

import (
	"database/sql"
	"service2/config"

	_ "github.com/lib/pq"
)

var Repo *Repository

type Repository struct {
	*sql.DB
}

func NewRepository(cfg *config.Config) *Repository {
	mdb, _ := sql.Open("postgres", cfg.PostgresCon())
	err := mdb.Ping()
	if err != nil {
		panic(err)
	}

	return &Repository{mdb}
}
