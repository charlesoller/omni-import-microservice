-- name: UpsertCreditCrewMember :one
INSERT INTO credits_crew_member (
  credit_id, member_id, job, department
) VALUES (
  $1, $2, $3, $4
)
ON CONFLICT (credit_id, member_id) DO UPDATE SET
  credit_id = EXCLUDED.credit_id,  
  member_id = EXCLUDED.member_id,
  job = EXCLUDED.job,
  department = EXCLUDED.department
RETURNING *;
