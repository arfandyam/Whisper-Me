package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
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

func (repository *ResponseRepository) FindResponseByQuestionId(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, offset int) ([]domain.Response, error){
	responses := []domain.Response{}
	sql := "SELECT id, question_id, response, created_at, updated_at FROM responses WHERE question_id = ? LIMIT ? OFFSET ?"

	rows := tx.Raw(sql, questionId, fetchPerPage, offset)
	if err := rows.Scan(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil
}