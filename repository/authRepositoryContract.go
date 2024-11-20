package repository

import (
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	InsertRefreshToken(tx *gorm.DB, token *string) error
	VerifyRefreshToken(db *gorm.DB, token *string) error
	DeleteRefreshToken(tx *gorm.DB, token string) error
}