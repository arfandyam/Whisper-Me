package repository

import (
	"fmt"

	"github.com/arfandyam/Whisper-Me/models/domain"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) CreateUser(tx *gorm.DB, user *domain.User) (*domain.User, error) {
	fmt.Println("user dari service", user)
	fmt.Println("user.Id dari repo", user.Id)
	fmt.Println("user.Password dari repo", *user.Password)
	sql := "INSERT INTO users (id, username, firstname, lastname, email, password, oauth_id, is_oauth, is_verified, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?) RETURNING id"
	
	rows := tx.Raw(sql, user.Id, user.Username, user.Firstname, user.Lastname, user.Email, *user.Password, user.Oauth_id, user.Is_oauth, user.Is_verified, nil, nil).Row()

	fmt.Println("rows", &rows)
	if err := rows.Scan(&user.Id); err != nil {
		tx.Rollback()
		return nil, err
	} else {
		tx.Commit()
		return user, nil
	}
}
