-- name: UpsertCreditCastMember :one
INSERT INTO credits_cast_member (
  credit_id, cast_id
) VALUES (
  $1, $2
)
ON CONFLICT (credit_id, cast_id) DO UPDATE SET
  credit_id = EXCLUDED.credit_id,  
  cast_id = EXCLUDED.cast_id
RETURNING *;
