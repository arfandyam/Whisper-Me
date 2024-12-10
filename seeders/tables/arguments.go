package tables

import (
	"github.com/google/uuid"
)

type Options struct {
	Amount  int
	UserId uuid.UUID
	QuestionId uuid.UUID
}