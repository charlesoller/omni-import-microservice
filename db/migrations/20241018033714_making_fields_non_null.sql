-- +goose Up
-- +goose StatementBegin

ALTER TABLE production_companies
    ALTER COLUMN logo_path SET NOT NULL,
    ALTER COLUMN origin_country SET NOT NULL;

ALTER TABLE collections
    ALTER COLUMN poster_path SET NOT NULL,
    ALTER COLUMN backdrop_path SET NOT NULL;

ALTER TABLE movies
    ALTER COLUMN overview SET NOT NULL,
    ALTER COLUMN release_date SET NOT NULL,
    ALTER COLUMN runtime SET NOT NULL,
    ALTER COLUMN budget SET NOT NULL,
    ALTER COLUMN revenue SET NOT NULL,
    ALTER COLUMN popularity SET NOT NULL,
    ALTER COLUMN vote_average SET NOT NULL,
    ALTER COLUMN vote_count SET NOT NULL,
    ALTER COLUMN status SET NOT NULL,
    ALTER COLUMN tagline SET NOT NULL,
    ALTER COLUMN homepage SET NOT NULL,
    ALTER COLUMN original_language SET NOT NULL,
    ALTER COLUMN adult SET NOT NULL,
    ALTER COLUMN backdrop_path SET NOT NULL,
    ALTER COLUMN poster_path SET NOT NULL,
    ALTER COLUMN collection_id SET NOT NULL;

ALTER TABLE movie_genres
    ALTER COLUMN movie_id SET NOT NULL,
    ALTER COLUMN genre_id SET NOT NULL;

ALTER TABLE movie_production_companies
    ALTER COLUMN movie_id SET NOT NULL,
    ALTER COLUMN company_id SET NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE production_companies
    ALTER COLUMN logo_path DROP NOT NULL,
    ALTER COLUMN origin_country DROP NOT NULL;

ALTER TABLE collections
    ALTER COLUMN poster_path DROP NOT NULL,
    ALTER COLUMN backdrop_path DROP NOT NULL;

ALTER TABLE movies
    ALTER COLUMN overview DROP NOT NULL,
    ALTER COLUMN release_date DROP NOT NULL,
    ALTER COLUMN runtime DROP NOT NULL,
    ALTER COLUMN budget DROP NOT NULL,
    ALTER COLUMN revenue DROP NOT NULL,
    ALTER COLUMN popularity DROP NOT NULL,
    ALTER COLUMN vote_average DROP NOT NULL,
    ALTER COLUMN vote_count DROP NOT NULL,
    ALTER COLUMN status DROP NOT NULL,
    ALTER COLUMN tagline DROP NOT NULL,
    ALTER COLUMN homepage DROP NOT NULL,
    ALTER COLUMN original_language DROP NOT NULL,
    ALTER COLUMN adult DROP NOT NULL,
    ALTER COLUMN backdrop_path DROP NOT NULL,
    ALTER COLUMN poster_path DROP NOT NULL,
    ALTER COLUMN collection_id DROP NOT NULL;

ALTER TABLE movie_genres
    ALTER COLUMN movie_id DROP NOT NULL,
    ALTER COLUMN genre_id DROP NOT NULL;

ALTER TABLE movie_production_companies
    ALTER COLUMN movie_id DROP NOT NULL,
    ALTER COLUMN company_id DROP NOT NULL;

-- +goose StatementEnd
