-- name: UpsertResponse :one
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
) ON CONFLICT (user_id, question_id) DO UPDATE SET
    answer = EXCLUDED.answer,
    updated_at = DATE('now')
RETURNING *;

-- name: GetAnsweredQuestions :many
SELECT
    *
FROM
    responses
WHERE
    user_id = ?
    AND survey_id = ?
    ;
