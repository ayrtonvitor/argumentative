-- name: CreateThesis :one
INSERT INTO thesis (
    id,
    creation_date,
    last_update_time,
    title,
	description
) VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
	$2
) RETURNING *;

