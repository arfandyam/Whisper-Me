package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseRepositoryInterface interface {
	CreateResponse(tx *gorm.DB, response *domain.Response) (*domain.Response, error)
	FindResponseByQuestionId(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, cursor *uuid.UUID) ([]domain.Response, error)
	FindPrevCursorResponse(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, cursor *uuid.UUID) (*uuid.UUID, error)
	SearchResponsesByKeyword(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, keyword string, rank *float64) ([]domain.Response, error)
	FindPrevRankResponse(tx *gorm.DB, questionId uuid.UUID, fetchPerPage int, keyword string, rank *float64) (*float64, error)
}