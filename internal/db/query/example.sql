-- name: ListExamples :many
SELECT * from examples
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetExample :one
SELECT * from examples
WHERE id = $1
LIMIT 1;

-- name: AddExample :one
INSERT INTO examples (name)
VALUES ($1)
RETURNING *;
