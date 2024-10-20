-- name: UpsertCountry :one
INSERT INTO countries (
  iso_3166_1, name
) VALUES (
  $1, $2
)
ON CONFLICT (iso_3166_1) DO UPDATE SET
  name = EXCLUDED.name
RETURNING *;

-- name: UpsertCountryISO :one
INSERT INTO countries (
  iso_3166_1
) VALUES (
  $1
)
ON CONFLICT (iso_3166_1) DO NOTHING
RETURNING *;