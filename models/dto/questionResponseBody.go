package dto

import "github.com/google/uuid"

type CreateEditQuestionResponse struct {
	*Response
	Id       uuid.UUID `json:"id"`
	Slug     string    `json:"slug"`
	Topic    string    `json:"topic"`
	Question string    `json:"question"`
}
