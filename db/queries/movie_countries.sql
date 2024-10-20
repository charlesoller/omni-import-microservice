-- name: UpsertMovieCountry :one
INSERT INTO movie_countries (
  movie_id, country_id
) VALUES (
  $1, $2
)
ON CONFLICT (movie_id, country_id) DO UPDATE SET
  movie_id = EXCLUDED.movie_id,  
  country_id = EXCLUDED.country_id
RETURNING *;
