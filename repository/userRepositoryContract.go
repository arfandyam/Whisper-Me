package repository

import (
	"github.com/arfandyam/Whisper-Me/models/domain"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(tx *gorm.DB, user *domain.User) (*domain.User, error)
	// EditUser(tx *gorm.DB, user domain.User) domain.User
	// FindUserById(tx *gorm.DB, user domain.User) domain.User
	// ChangePassword(tx *gorm.DB, user domain.User)
}