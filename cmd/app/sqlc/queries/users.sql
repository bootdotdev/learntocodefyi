-- name: InsertUser :one
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
) 
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;
