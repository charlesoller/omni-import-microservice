-- name: UpsertCastMember :one
INSERT INTO cast_members (
    id, cast_id, credit_id, gender, adult, known_for_department, 
    name, original_name, popularity, profile_path
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
ON CONFLICT (id) 
DO UPDATE SET
    cast_id = EXCLUDED.cast_id,
    credit_id = EXCLUDED.credit_id,
    gender = EXCLUDED.gender,
    adult = EXCLUDED.adult,
    known_for_department = EXCLUDED.known_for_department,
    name = EXCLUDED.name,
    original_name = EXCLUDED.original_name,
    popularity = EXCLUDED.popularity,
    profile_path = EXCLUDED.profile_path
RETURNING *;
