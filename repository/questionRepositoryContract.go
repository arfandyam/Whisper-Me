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
}