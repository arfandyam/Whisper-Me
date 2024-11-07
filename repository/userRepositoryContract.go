package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(tx *gorm.DB, user *domain.User) (*domain.User, error)
	EditUser(tx *gorm.DB, user *domain.User) (*domain.User, error)
	FindUserById(tx *gorm.DB, userId uuid.UUID) (*domain.User, error)
	// ChangePassword(tx *gorm.DB, user domain.User)
}