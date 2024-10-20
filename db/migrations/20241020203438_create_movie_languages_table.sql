-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS movie_languages (
    movie_id INT REFERENCES movies(id),
    language_id VARCHAR(2) REFERENCES languages(iso_639_1),
    PRIMARY KEY (movie_id, language_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_languages;
-- +goose StatementEnd
