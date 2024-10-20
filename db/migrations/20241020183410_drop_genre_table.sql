-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS genres
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT ''
);
-- +goose StatementEnd
