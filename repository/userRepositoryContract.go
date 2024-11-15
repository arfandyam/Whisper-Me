package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(tx *gorm.DB, user *domain.User) (*domain.User, error)
	EditUser(tx *gorm.DB, user *domain.User) (*domain.User, error)
	FindUserById(db *gorm.DB, userId uuid.UUID) (*domain.User, error)
	ChangeUserPassword(tx *gorm.DB, userId uuid.UUID, password string) error
	GetUserPassword(db *gorm.DB, userId uuid.UUID) (*string, error)
	GetUserCredentials(db *gorm.DB, username string) (*domain.User, error)
}
