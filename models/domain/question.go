package domain

import "github.com/google/uuid"

type Question struct {
	Id       uuid.UUID
	UserId   uuid.UUID
	Slug     string
	Topic    string
	Question string
	Rank     float64
}
