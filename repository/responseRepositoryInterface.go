package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseRepositoryInterface interface {
	CreateResponse(tx *gorm.DB, response *domain.Response) (*domain.Response, error)
	FindResponseByQuestionId(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, offset int) ([]domain.Response, error)
}