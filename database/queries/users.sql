-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING id, name, email;

-- name: UpdateUserInformation :one
UPDATE users SET
name = $2,
email = $3
WHERE id = $1
RETURNING id, name, email;

-- name: UpdateUserPassword :exec
UPDATE users SET
password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

