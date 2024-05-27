-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key, password)
VALUES ($1, $2, $2, $3, encode(sha256(random()::text::bytea), 'hex'), crypt($4, gen_salt('bf')))
RETURNING id, created_at, updated_at, name, api_key;

-- name: GetUserByAPI :one
SELECT * FROM users
WHERE api_key = $1;

-- name: LogInAsUser :one
SELECT id, created_at, updated_at, name, api_key FROM users
WHERE name = $1
AND password = crypt($2, password);

