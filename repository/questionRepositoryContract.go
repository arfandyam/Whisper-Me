package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepositoryInterface interface {
	CreateQuestion(tx *gorm.DB, question *domain.Question) (*domain.Question, error)
	EditQuestion(tx *gorm.DB, question *domain.Question) (*domain.Question, error)
	FindQuestionById(tx *gorm.DB, questionId uuid.UUID) (*domain.Question, error)
	FindQuestionsByUserId(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, offset int) ([]domain.Question, error)
	FindQuestionOwner(tx *gorm.DB, questionId uuid.UUID) (*uuid.UUID, error)
	SearchQuestionsByKeyword(tx *gorm.DB, userId uuid.UUID, keyword string, fetchPerPage int, offset int) ([]domain.Question, error)
}