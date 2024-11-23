package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionRepositoryInterface interface {
	CreateQuestion(tx *gorm.DB, userId uuid.UUID, question *domain.Question) (*domain.Question, error)
}