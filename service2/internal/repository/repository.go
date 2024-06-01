package repository

import (
	"context"
	"database/sql"
	"fmt"
	"service2/config"
	"service2/internal/logwrapper"
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
		logwrapper.Logger.Payload(logwrapper.NewPayload().Op("NewRepository").Package("repository")).Fatal(err)
		panic(err)
	}

	return &Repository{mdb}
}

func (r *Repository) FindHashesByIds(ctx context.Context, ids []int64) (m []*models.Hash, err error) {
	const op = "FindHashesByIds"

	defer func() {
		if err != nil {
			em := fmt.Errorf("%s: %w", op, err)
			logwrapper.Logger.Payload(logwrapper.NewPayload().Op(op).Package("repository")).Error(em)
		}
	}()

	params := make([]string, 0, len(ids))
	args := make([]any, 0, len(ids))
	for i, id := range ids {
		params = append(params, fmt.Sprintf("$%d", i+1))
		args = append(args, id)
	}
	query := fmt.Sprintf("SELECT id, hash FROM hashes WHERE id IN (%s)", strings.Join(params, ","))

	rows, err := r.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func() { _ = rows.Close() }()

	result := make([]*models.Hash, 0)
	for rows.Next() {
		var hash models.Hash
		err = rows.Scan(&hash.ID, &hash.Hash)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		result = append(result, &hash)
	}

	return result, nil
}

func (r *Repository) FindHashes(ctx context.Context, hashes []string) (m []*models.Hash, err error) {
	const op = "FindHashes"

	defer func() {
		if err != nil {
			em := fmt.Errorf("%s: %w", op, err)
			logwrapper.Logger.Payload(logwrapper.NewPayload().Op(op).Package("repository")).Error(em)
		}
	}()

	params := make([]string, 0, len(hashes))
	args := make([]any, 0, len(hashes))
	for i, h := range hashes {
		params = append(params, fmt.Sprintf("$%d", i+1))
		args = append(args, h)
	}
	query := fmt.Sprintf("SELECT id, hash FROM hashes WHERE hash IN (%s)", strings.Join(params, ","))

	rows, err := r.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer func() { _ = rows.Close() }()

	result := make([]*models.Hash, 0)
	for rows.Next() {
		var hash models.Hash
		err = rows.Scan(&hash.ID, &hash.Hash)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		result = append(result, &hash)
	}

	return result, nil
}

func (r *Repository) Save(ctx context.Context, hashes []string) (result []*models.Hash, err error) {
	const op = "Repository.Save"

	defer func() {
		if err != nil {
			em := fmt.Errorf("%s: %w", op, err)
			logwrapper.Logger.Payload(logwrapper.NewPayload().Op(op).Package("repository")).Error(em)
		}
	}()

	if len(hashes) == 0 {
		return make([]*models.Hash, 0), nil
	}

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
		result = append(result, &models.Hash{
			Hash: &hashes[i],
			ID:   &id,
		})
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
