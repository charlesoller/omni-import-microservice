-- name: UpsertMovie :one
INSERT INTO movies (
  id, title, original_title, overview, release_date, runtime, budget, revenue,
  popularity, vote_average, vote_count, status, tagline, homepage,
  original_language, adult, backdrop_path, poster_path, collection_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8,
  $9, $10, $11, $12, $13, $14,
  $15, $16, $17, $18, $19
)
ON CONFLICT (id) DO UPDATE SET
  title = EXCLUDED.title,
  original_title = EXCLUDED.original_title,
  overview = EXCLUDED.overview,
  release_date = EXCLUDED.release_date,
  runtime = EXCLUDED.runtime,
  budget = EXCLUDED.budget,
  revenue = EXCLUDED.revenue,
  popularity = EXCLUDED.popularity,
  vote_average = EXCLUDED.vote_average,
  vote_count = EXCLUDED.vote_count,
  status = EXCLUDED.status,
  tagline = EXCLUDED.tagline,
  homepage = EXCLUDED.homepage,
  original_language = EXCLUDED.original_language,
  adult = EXCLUDED.adult,
  backdrop_path = EXCLUDED.backdrop_path,
  poster_path = EXCLUDED.poster_path,
  collection_id = EXCLUDED.collection_id
RETURNING *;

-- name: UpdateMovieEmbedding :exec
UPDATE movies 
SET embedding = $2
WHERE id = $1;
