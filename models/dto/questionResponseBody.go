package dto

import (
	"time"
	"github.com/google/uuid"
)

type QuestionDTO struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Slug      string    `json:"slug"`
	Topic     string    `json:"topic"`
	Question  string    `json:"question"`
	UrlKey    string    `json:"url_key"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateEditQuestionResponse struct {
	*Response
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Slug     string    `json:"slug"`
	Topic    string    `json:"topic"`
	Question string    `json:"question"`
}

type FindQuestionResponse struct {
	*Response
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Slug     string    `json:"slug"`
	Topic    string    `json:"topic"`
	Question string    `json:"question"`
}

type FindQuestionsByUserIdResponse struct {
	*Response
	Data []QuestionDTO  `json:"data"`
	Meta PageCursorInfo `json:"meta"`
}

type SearchKeywordQuestionsByUserIdResponse struct {
	*Response
	Data []QuestionDTO `json:"data"`
	Meta PageRankInfo  `json:"meta"`
}
