package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"gorm.io/gorm"
)

type ResponseRepositoryInterface interface {
	CreateResponse(tx *gorm.DB, response *domain.Response) (*domain.Response, error)
}