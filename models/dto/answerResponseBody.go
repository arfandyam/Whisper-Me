package dto

import (
	"github.com/google/uuid"
	"time"
)

type ResponseDTO struct {
	Id         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Response   string    `json:"response"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAnswerResponse struct {
	*Response
	Data ResponseDTO `json:"data"`
}

type FindAnswerResponse struct {
	*Response
	Data []ResponseDTO  `json:"data"`
	Meta PageCursorInfo `json:"meta"`
}

type SearchKeywordResponseByUserIdResponse struct {
	*Response
	Data []ResponseDTO `json:"data"`
	Meta PageRankInfo  `json:"meta"`
}
