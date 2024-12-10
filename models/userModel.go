package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	Username    string    `gorm:"type:varchar(255);unique;not null"`
	Firstname   string    `gorm:"type:varchar(255);not null"`
	Lastname    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string
	Oauth_id    string
	Is_oauth    bool 	`gorm:"not null"`
	Is_verified bool `gorm:"not null"`
	Questions   []Question
	Sessions    []Session
}

type Question struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	UserId    uuid.UUID `gorm:"type:uuid"`
	Slug      string `gorm:"unique;not null"`
	Topic     string `gorm:"not null"`
	Question  string `gorm:"not null"`
	Responses []Response
}

type Response struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	QuestionID uuid.UUID `gorm:"type:uuid"`
	Response   string `gorm:"not null"`
}

type Session struct {
	UserID     uuid.UUID `gorm:"type:uuid"`
	Token      string    `gorm:"unique;not null"`
	Issued_at  time.Time `gorm:"not null"`
	Expired_at time.Time `gorm:"not null"`
}
