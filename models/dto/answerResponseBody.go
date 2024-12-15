package dto

import (
	"github.com/google/uuid"
	"time"
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
	Meta PageCursorInfo      `json:"meta"`
}

type SearchKeywordResponseByUserIdResponse struct {
	*Response
	Data []FullResponseDTO `json:"data"`
	Meta PageRankInfo      `json:"meta"`
}
