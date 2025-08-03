-- name: CreateArgument :one
INSERT INTO argument (
  id,
  creation_date,
  last_update_time,
  brief,
  description
)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
) RETURNING *;
