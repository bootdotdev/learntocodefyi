// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package sqlcdb

import (
	"time"
)

type Response struct {
	ID         string
	UserID     string
	QuestionID string
	SurveyID   string
	Answer     string
	CreatedAt  time.Time
}

type User struct {
	ID        string
	HashedIp  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
