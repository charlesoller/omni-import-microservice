-- name: UpsertCrewMember :one
INSERT INTO crew_members (
    id, credit_id, gender, adult, known_for_department,
    name, original_name, popularity, profile_path
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (id) DO UPDATE SET
    gender = EXCLUDED.gender,
    adult = EXCLUDED.adult,
    known_for_department = EXCLUDED.known_for_department,
    name = EXCLUDED.name,
    original_name = EXCLUDED.original_name,
    popularity = EXCLUDED.popularity,
    profile_path = EXCLUDED.profile_path
RETURNING id;
