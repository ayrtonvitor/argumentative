-- name: GetThesisById :many
SELECT
  id,
  creation_date,
  last_update_time,
  title,
  description
FROM thesis
WHERE id = ANY($1::UUID[]);
