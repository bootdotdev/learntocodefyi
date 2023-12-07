-- name: CreateUser :one
INSERT INTO users (
    id,
    hashed_ip,
    created_at,
    updated_at
) VALUES (
    ?,
    ?,
    DATE('now'),
    DATE('now')
) RETURNING *;
