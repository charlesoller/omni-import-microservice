-- +goose Up
-- +goose StatementBegin

CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    english_name VARCHAR(255) NOT NULL DEFAULT '',
    iso_639_1 VARCHAR(2) UNIQUE NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL DEFAULT ''
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE languages;

-- +goose StatementEnd
