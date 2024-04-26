-- +goose Up 
CREATE TABLE feed_follows(
	id TEXT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	feed_id TEXT NOT NULL,
	FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
	user_id TEXT NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
	UNIQUE(user_id, feed_id) 
);

-- +goose Down 
DROP TABLE feed_follows;

