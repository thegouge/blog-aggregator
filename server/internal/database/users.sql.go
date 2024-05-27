// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key, password)
VALUES ($1, $2, $2, $3, encode(sha256(random()::text::bytea), 'hex'), crypt($4, gen_salt('bf')))
RETURNING id, created_at, updated_at, name, api_key
`

type CreateUserParams struct {
	ID        string
	CreatedAt time.Time
	Name      string
	Crypt     string
}

type CreateUserRow struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.Name,
		arg.Crypt,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

const getUserByAPI = `-- name: GetUserByAPI :one
SELECT id, created_at, updated_at, name, api_key, password FROM users
WHERE api_key = $1
`

func (q *Queries) GetUserByAPI(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAPI, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
		&i.Password,
	)
	return i, err
}

const logInAsUser = `-- name: LogInAsUser :one
SELECT id, created_at, updated_at, name, api_key FROM users
WHERE name = $1
AND password = crypt($2, password)
`

type LogInAsUserParams struct {
	Name  string
	Crypt string
}

type LogInAsUserRow struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

func (q *Queries) LogInAsUser(ctx context.Context, arg LogInAsUserParams) (LogInAsUserRow, error) {
	row := q.db.QueryRowContext(ctx, logInAsUser, arg.Name, arg.Crypt)
	var i LogInAsUserRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}
