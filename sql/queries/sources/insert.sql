-- name: CreateArgumentSources :one
INSERT INTO argumentSources (
  id,
  creation_date,
  last_update_time,
  content,
  argument_id
)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;

