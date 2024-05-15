-- name: CreateFeedFollow :one
INSERT INTO feed_follows(id, created_at, updated_at, feed_id, user_id)
VALUES($1, $2, $2, $3, $4)
RETURNING *;
--

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE $1 = id AND user_id = $2;
--

-- name: GetUsersFeedFollows :many
SELECT * FROM feed_follows
WHERE user_id = $1;

