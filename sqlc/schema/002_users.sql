-- name: CreateTableUsers :exec
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    hashed_ip TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
