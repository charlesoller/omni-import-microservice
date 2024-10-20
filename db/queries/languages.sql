-- name: UpsertLanguage :one
INSERT INTO languages (
  english_name, iso_639_1, name
) VALUES (
  $1, $2, $3
)
ON CONFLICT (iso_639_1) DO UPDATE SET
  english_name = EXCLUDED.english_name,
  name = EXCLUDED.name
RETURNING *;