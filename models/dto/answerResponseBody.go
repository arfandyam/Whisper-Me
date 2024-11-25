package dto

import "github.com/google/uuid"

type ResponseDTO struct {
	Id         uuid.UUID `json:"id"`
	QuestionId uuid.UUID `json:"question_id"`
	Response   string    `json:"response"`
}

type CreateEditAnswerResponse struct {
	*Response
	Data ResponseDTO
}
