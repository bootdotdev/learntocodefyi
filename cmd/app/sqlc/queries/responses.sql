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

-- name: GetAnsweredQuestions :many
SELECT
    *
FROM
    responses
WHERE
    user_id = ?
    AND survey_id = ?
    AND question_id IN (?);
