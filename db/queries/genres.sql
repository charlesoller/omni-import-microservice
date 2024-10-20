-- name: UpsertGenre :one
INSERT INTO genres (
  id, name
) VALUES (
  $1, $2
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name
RETURNING *;