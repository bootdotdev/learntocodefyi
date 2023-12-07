-- name: CreateResponse :one
INSERT INTO responses (
    id,
    user_id,
    question_id,
    survey_id,
    answer,
    created_at
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    DATE('now')
) RETURNING *;
