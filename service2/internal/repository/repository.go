package repository

import (
	"context"
	"database/sql"
	"fmt"
	"service2/config"
	"service2/models"
	"strings"

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

func (r *Repository) Save(ctx context.Context, hashes []string) (result []models.Hash, err error) {
	const op = "Repository.Save"

	tx, err := r.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	params := make([]string, 0, len(hashes))
	args := make([]any, 0, len(hashes))
	for i, h := range hashes {
		params = append(params, fmt.Sprintf("($%d)", i+1))
		args = append(args, h)
	}
	query := fmt.Sprintf("INSERT INTO hashes (hash) VALUES %s RETURNING id", strings.Join(params, ","))

	rows, err := r.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func() { _ = rows.Close() }()

	var ids []int64
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("%s: get id: %w", op, err)
		}
		ids = append(ids, id)
	}

	if len(ids) != len(hashes) {
		return nil, fmt.Errorf("%s: hashes length mismatch ids length", op)
	}

	for i, id := range ids {
		result = append(result, models.Hash{
			Hash: &hashes[i],
			ID:   &id,
		})
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
