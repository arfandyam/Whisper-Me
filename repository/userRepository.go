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

func (repository *UserRepository) FindUserById(db *gorm.DB, userId uuid.UUID) (*domain.User, error){
	sql := "SELECT * FROM users WHERE id = ?"
	user := &domain.User{}

	if err := db.Raw(sql, userId).Scan(&user).Error; err != nil {
		fmt.Println("oiii")
		return nil, err
	}

	fmt.Println("oiii2")
	return user, nil
}

func (repository *UserRepository) ChangeUserPassword(tx *gorm.DB, userId uuid.UUID, password string) error {
	sql := "UPDATE users SET password = ? WHERE id = ?"

	if err := tx.Exec(sql, password, userId).Error; err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) GetUserPassword(db *gorm.DB, userId uuid.UUID) (*string, error) {
	var password string
	sql := "SELECT password FROM users WHERE id = ?"
	if err := db.Raw(sql, userId).Scan(&password).Error; err != nil {
		return nil, err
	}

	return &password, nil
}

func (repository *UserRepository) GetUserCredentials(db *gorm.DB, username string) (*domain.User, error) {
	fmt.Println("username", username)
	user := domain.User{}
	sql := "SELECT id, username, password FROM users WHERE username = ?"
	rows := db.Raw(sql, username).Row()
	if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
