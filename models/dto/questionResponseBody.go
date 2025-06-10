package dto

import (
	"github.com/google/uuid"
	"time"
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

type QuestionSlug struct {
	Slug string `json:"slug"`
}

type CreateEditQuestionResponse struct {
	*Response
	Data QuestionDTO `json:"data"`
}

type FindQuestionResponse struct {
	*Response
	Data QuestionDTO `json:"data"`
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

type FindQuestionSlugByUrlKey struct {
	*Response
	Data QuestionSlug `json:"data"`
}
