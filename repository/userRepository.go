package repository

import (
	"fmt"

	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (repository *UserRepository) CreateUser(tx *gorm.DB, user *domain.User) (*domain.User, error) {
	sql := "INSERT INTO users (id, username, firstname, lastname, email, password, oauth_id, is_oauth, is_verified, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?) RETURNING id"
	
	rows := tx.Raw(sql, user.Id, user.Username, user.Firstname, user.Lastname, user.Email, &user.Password, user.Oauth_id, user.Is_oauth, user.Is_verified, nil, nil).Row()

	if err := rows.Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) EditUser(tx *gorm.DB, user *domain.User) (*domain.User, error) {
	fmt.Println("kocakedituser")
	sql := "UPDATE users SET firstname = ?, lastname = ?, updated_at = NOW() WHERE id = ? RETURNING firstname, lastname"

	rows := tx.Raw(sql, user.Firstname, user.Lastname, user.Id).Row()

	if err := rows.Scan(&user.Firstname, &user.Lastname); err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) FindUserById(tx *gorm.DB, userId uuid.UUID) (*domain.User, error){
	sql := "SELECT * FROM users WHERE id = ?"
	user := &domain.User{}

	if err := tx.Raw(sql, userId).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
