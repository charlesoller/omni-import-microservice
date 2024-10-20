-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS origin_countries (
    id SERIAL PRIMARY KEY,
    iso_3166_1 VARCHAR(2) UNIQUE NOT NULL DEFAULT '',
    name VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS production_companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    logo_path VARCHAR(255) DEFAULT '',
    origin_country VARCHAR(2) REFERENCES origin_countries(iso_3166_1) DEFAULT ''
);

CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS collections (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT '',
    poster_path VARCHAR(255) DEFAULT '',
    backdrop_path VARCHAR(255) DEFAULT ''
);

CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL DEFAULT '',
    original_title VARCHAR(255) NOT NULL DEFAULT '',
    overview TEXT DEFAULT '',
    release_date DATE DEFAULT NULL,
    runtime INT DEFAULT 0,
    budget BIGINT DEFAULT 0,
    revenue BIGINT DEFAULT 0,
    popularity FLOAT DEFAULT 0,
    vote_average FLOAT DEFAULT 0,
    vote_count INT DEFAULT 0,
    status VARCHAR(50) DEFAULT '',
    tagline VARCHAR(255) DEFAULT '',
    homepage VARCHAR(255) DEFAULT '',
    original_language VARCHAR(10) DEFAULT '',
    adult BOOLEAN DEFAULT false,
    backdrop_path VARCHAR(255) DEFAULT '',
    poster_path VARCHAR(255) DEFAULT '',
    collection_id INT REFERENCES collections(id) DEFAULT 0
);

CREATE TABLE IF NOT EXISTS movie_genres (
    movie_id INT REFERENCES movies(id),
    genre_id INT REFERENCES genres(id),
    PRIMARY KEY (movie_id, genre_id)
);

CREATE TABLE IF NOT EXISTS movie_production_companies (
    movie_id INT REFERENCES movies(id),
    company_id INT REFERENCES production_companies(id),
    PRIMARY KEY (movie_id, company_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS movie_production_companies;
DROP TABLE IF EXISTS movie_genres;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS collections;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS production_companies;
DROP TABLE IF EXISTS origin_countries;

-- +goose StatementEnd
