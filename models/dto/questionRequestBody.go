package dto

type CreateQuestionRequest struct {
	Question string `validate:"required" json:"question"`
	Topic    string `validate:"required" json:"topic"`
}
