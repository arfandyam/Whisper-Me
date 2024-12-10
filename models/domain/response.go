package domain

import (
	"github.com/google/uuid"
	"time"
)

type Response struct {
	Id         uuid.UUID
	QuestionId uuid.UUID
	Response   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
