-- name: GetSourceFromArgumentId :many
SELECT
  id,
  creation_date,
  last_update_time,
  content,
  argument_id
FROM argumentSources
WHERE argument_id = ANY(@ids::UUID[]);
