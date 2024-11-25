package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"gorm.io/gorm"
)

type ResponseRepository struct{}

func NewResponseRepository() ResponseRepositoryInterface {
	return &ResponseRepository{}
}

func (repository *ResponseRepository) CreateResponse(tx *gorm.DB, response *domain.Response) (*domain.Response, error){
	sql := "INSERT INTO responses (id, question_id, response, created_at, updated_at, deleted_at) VALUES (?, ?, ?, NOW(), ?, ?) returning id, question_id, response"

	rows := tx.Raw(sql, response.Id, response.QuestionId, response.Response, nil, nil).Row()
	if err := rows.Scan(&response.Id, &response.QuestionId, &response.Response); err != nil {
		return nil, err
	}

	return response, nil
}