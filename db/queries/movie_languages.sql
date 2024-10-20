-- name: UpsertMovieLanguage :one
INSERT INTO movie_languages (
  movie_id, language_id
) VALUES (
  $1, $2
)
ON CONFLICT (movie_id, language_id) DO UPDATE SET
  movie_id = EXCLUDED.movie_id,  
  language_id = EXCLUDED.language_id
RETURNING *;
