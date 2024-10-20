-- +goose Up
-- +goose StatementBegin

ALTER TABLE movie_production_countries
RENAME TO movie_countries;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE movie_countries
RENAME TO movie_production_countries;

-- +goose StatementEnd