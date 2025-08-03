-- name: InsertTheisJoinArgument :exec
INSERT INTO thesis_argument (
  thesis_id,
  argument_id
)
VALUES (
  $1,
  $2
);
