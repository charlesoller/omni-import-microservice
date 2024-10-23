-- name: UpsertCreditCastMember :one
INSERT INTO credits_cast_member (
  credit_id, member_id, character, "order"
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (credit_id, member_id) DO UPDATE SET
  credit_id = EXCLUDED.credit_id,  
  member_id = EXCLUDED.member_id,
  character = EXCLUDED.character,
  "order" = EXCLUDED."order"
RETURNING *;
