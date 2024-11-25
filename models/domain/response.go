package domain

import "github.com/google/uuid"

type Response struct {
	Id         uuid.UUID
	QuestionId uuid.UUID
	Response   string
}
