package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (repository *AuthRepository) InsertRefreshToken(tx *gorm.DB, token *string) error {
	sql := "INSERT INTO authentications (token) VALUES(?) RETURNING token"

	rows := tx.Raw(sql, token).Row()
	if err := rows.Scan(&token); err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepository) VerifyRefreshToken(db *gorm.DB, token *string) error {
	sql := "SELECT token FROM authentications WHERE token = ?"

	if err := db.Raw(sql, token).Scan(&token).Error; err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepository) DeleteRefreshToken(tx *gorm.DB, token string) error {
	fmt.Println("token dari repo", token)
	sql := "DELETE FROM authentications WHERE token = ?"

	if err := tx.Exec(sql, token).Error; err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
