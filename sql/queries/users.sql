-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $2, $3, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByAPI :one
SELECT * FROM users
WHERE api_key = $1;

