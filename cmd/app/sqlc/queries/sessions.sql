-- name: CreateSession :one
INSERT INTO sessions (
    id,
    created_at,
    user_id
) VALUES (
    ?,
    DATE('now'),
    ?
) RETURNING *;

-- name: GetSessionCreatedAfter :one
SELECT * FROM sessions
WHERE user_id = ?
    AND created_at > ?;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = ?;
