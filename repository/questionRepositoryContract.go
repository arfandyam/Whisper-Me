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
	FindQuestionBySlug(tx *gorm.DB, questionSlug string) (*domain.Question, error)
	FindQuestionsByUserId(tx *gorm.DB, userId uuid.UUID, cursor *uuid.UUID, fetchPerPage int) ([]domain.Question, error)
	FindPrevCursorQuestion(tx *gorm.DB, userId uuid.UUID, cursor *uuid.UUID, fetchPerPage int) (*uuid.UUID, error)
	SearchQuestionsByKeyword(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, keyword string, rank *float64) ([]domain.Question, error)
	FindPrevRankQuestion(tx *gorm.DB, userId uuid.UUID, fetchPerPage int, keyword string, rank *float64) (*float64, error)
	FindQuestionOwner(tx *gorm.DB, questionId uuid.UUID) (*uuid.UUID, error)
}