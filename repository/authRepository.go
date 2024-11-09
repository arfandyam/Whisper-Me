package repository

import (
	"gorm.io/gorm"
)

type AuthRepository struct{}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

func (repository *AuthRepository) LoginUser(tx *gorm.DB, token *string) error {
	sql := "INSERT INTO authentications (token) VALUES(?) RETURNING token"

	rows := tx.Raw(sql, token).Row()
	if err := rows.Scan(&token); err != nil {
		return err
	}

	return nil
}

// func (repository *AuthRepository) VerifyRefreshToken(db *gorm.DB, token *string) error {
// 	sql := "SELECT token FROM authentications WHERE token = ?"

// 	if err := db.Raw(sql, token).Scan(&token).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
