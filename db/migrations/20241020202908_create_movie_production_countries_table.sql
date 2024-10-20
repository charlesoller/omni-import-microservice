-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_production_countries (
    movie_id INT REFERENCES movies(id),
    country_id VARCHAR(2) REFERENCES countries(iso_3166_1),
    PRIMARY KEY (movie_id, country_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_production_countries;
-- +goose StatementEnd
