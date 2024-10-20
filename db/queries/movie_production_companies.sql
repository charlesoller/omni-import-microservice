-- name: UpsertMovieProductionCompany :one
INSERT INTO movie_production_companies (
  movie_id, company_id
) VALUES (
  $1, $2
)
ON CONFLICT (movie_id, company_id) DO UPDATE SET
  movie_id = EXCLUDED.movie_id,  
  company_id = EXCLUDED.company_id
RETURNING *;
