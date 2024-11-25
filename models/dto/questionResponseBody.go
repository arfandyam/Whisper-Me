package dto

import "github.com/google/uuid"

type QuestionDTO struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	Slug     string    `json:"slug"`
	Topic    string    `json:"topic"`
	Question string    `json:"question"`
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
	Data []QuestionDTO `json:"data"`
}
