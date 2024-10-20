-- name: UpsertCollection :one
INSERT INTO collections (
  id, name, poster_path, backdrop_path
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  backdrop_path = EXCLUDED.backdrop_path,
  poster_path = EXCLUDED.poster_path
RETURNING *;