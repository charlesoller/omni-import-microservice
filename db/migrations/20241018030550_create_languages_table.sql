-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS languages (
    id SERIAL PRIMARY KEY,
    english_name VARCHAR(255) NOT NULL DEFAULT '',
    iso_639_1 VARCHAR(2) UNIQUE NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL DEFAULT ''
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS languages;

-- +goose StatementEnd
