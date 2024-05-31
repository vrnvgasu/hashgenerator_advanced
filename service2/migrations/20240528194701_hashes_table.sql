-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hashes (
    id SERIAL PRIMARY KEY,
    hash TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hashesl
-- +goose StatementEnd
