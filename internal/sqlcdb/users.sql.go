// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: users.sql

package sqlcdb

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, email FROM users
WHERE id = ?
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email FROM users
WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
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
RETURNING id, created_at, updated_at, email
`

type InsertUserParams struct {
	ID    string
	Email string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}
