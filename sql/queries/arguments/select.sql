-- name: GetArgumentFromThesisId :many
SELECT
  argument.id,
  argument.creation_date,
  argument.last_update_time,
  argument.brief,
  argument.description
FROM argument
JOIN thesis_argument
  ON argument.id = thesis_argument.argument_id
WHERE thesis_argument.thesis_id = $1;
