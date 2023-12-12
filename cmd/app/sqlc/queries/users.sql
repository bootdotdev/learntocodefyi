-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    created_at,
    updated_at
) VALUES (
    ?,
    ?,
    DATE('now'),
    DATE('now')
) RETURNING *;
