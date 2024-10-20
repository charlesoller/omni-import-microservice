-- name: UpsertMovieGenre :one
INSERT INTO movie_genres (
  movie_id, genre_id
) VALUES (
  $1, $2
)
ON CONFLICT (movie_id, genre_id) DO UPDATE SET
  movie_id = EXCLUDED.movie_id,  
  genre_id = EXCLUDED.genre_id
RETURNING *;
