-- name: UpsertProductionCompany :one
INSERT INTO production_companies (
  id, name, logo_path, origin_country
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  logo_path = EXCLUDED.logo_path,
  origin_country = EXCLUDED.origin_country
RETURNING *;