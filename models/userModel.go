package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	// "time"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primaryKey;not null"`
	Username    string    `gorm:"unique;not null"`
	Firstname   string    `gorm:"not null"`
	Lastname    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string
	Oauth_id    string
	Is_oauth    bool `gorm:"not null"`
	Is_verified bool `gorm:"not null"`
	Questions   []Question
	Sessions    []Session
}

type Question struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;not null"`
	UserId    uuid.UUID
	Slug      string `gorm:"not null"`
	Question  string `gorm:"not null"`
	Responses []Response
}

type Response struct {
	gorm.Model
	Id         uuid.UUID `gorm:"primaryKey;not null"`
	QuestionID uuid.UUID
	Slug       string `gorm:"not null"`
	Response   string `gorm:"not null"`
}

type Session struct {
	UserID     uuid.UUID
	Token      string    `gorm:"unique;not null"`
	Issued_at  time.Time `gorm:"not null"`
	Expired_at time.Time `gorm:"not null"`
}
