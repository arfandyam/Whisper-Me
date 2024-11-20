package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (repository *AuthRepository) InsertRefreshToken(tx *gorm.DB, userId uuid.UUID, token string, iat time.Time, exp time.Time) error {
	sql := "INSERT INTO sessions (user_id, token, issued_at, expired_at) VALUES(?, ?, ?, ?) RETURNING token"

	rows := tx.Raw(sql, userId, token, iat, exp).Row()
	if err := rows.Scan(&token); err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepository) VerifyRefreshToken(db *gorm.DB, token *string) error {
	sql := "SELECT token FROM sessions WHERE token = ?"

	if err := db.Raw(sql, token).Scan(&token).Error; err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepository) DeleteRefreshToken(tx *gorm.DB, token string) error {
	fmt.Println("token dari repo", token)
	sql := "DELETE FROM sessions WHERE token = ?"

	if err := tx.Exec(sql, token).Error; err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
