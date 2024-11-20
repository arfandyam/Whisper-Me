package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	InsertRefreshToken(tx *gorm.DB, userId uuid.UUID, token string, iat time.Time, exp time.Time) error
	VerifyRefreshToken(db *gorm.DB, token *string) error
	DeleteRefreshToken(tx *gorm.DB, token string) error
}