-- +goose Up 
CREATE EXTENSION pgcrypto;
ALTER TABLE users
ADD COLUMN password TEXT NOT NULL;

-- +goose Down
DROP EXTENSION pgcrypto;
ALTER TABLE users
DROP COLUMN password;

