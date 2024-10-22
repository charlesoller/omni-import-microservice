-- name: UpsertCredit :exec
INSERT INTO credits (
  id
) VALUES (
  $1
)
ON CONFLICT (id) DO NOTHING;