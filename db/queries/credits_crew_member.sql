-- name: UpsertCreditCrewMember :one
INSERT INTO credits_crew_member (
  credit_id, crew_id
) VALUES (
  $1, $2
)
ON CONFLICT (credit_id, crew_id) DO UPDATE SET
  credit_id = EXCLUDED.credit_id,  
  crew_id = EXCLUDED.crew_id
RETURNING *;
