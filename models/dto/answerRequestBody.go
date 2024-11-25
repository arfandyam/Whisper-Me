package dto

import "github.com/google/uuid"

type CreateEditAnswerRequestBody struct {
	QuestionId uuid.UUID `json:"question_id"`
	Response   string    `json:"response"`
}
