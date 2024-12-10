package dto

import (
	"time"

	"github.com/google/uuid"
)

type ResponseDTO struct {
	Id         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Response   string    `json:"response"`
}

type ResponseTimestampDTO struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FullResponseDTO struct {
	ResponseDTO
	ResponseTimestampDTO
}

type CreateEditAnswerResponse struct {
	*Response
	Data ResponseDTO
}

type FindAnswerResponse struct {
	*Response
	Data []FullResponseDTO `json:"data"`
}
