-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
VALUES($1, $2, $2, $3, $4)
RETURNING *;

