package domain

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Slug      string
	Topic     string
	Question  string
	UrlKey    string
	Rank      float64
	CreatedAt time.Time
}
