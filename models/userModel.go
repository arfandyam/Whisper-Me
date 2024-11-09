package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	// "time"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `gorm:primaryKey;notNull`
	Username    string    `gorm:unique;notNull`
	Firstname   string    `gorm:notNull`
	Lastname    string    `gorm:notNull`
	Email       string    `gorm:unique;notNull`
	Password    string
	Oauth_id    string
	Is_oauth    bool `gorm:notNull`
	Is_verified bool `gorm:notNull`
}

type Question struct {
	gorm.Model
	Id       uuid.UUID `gorm:primaryKey;notNull`
	User     User      `gorm:"foreignKey:id";notNull`
	Slug     string    `gorm:notNull`
	Question string    `gorm:notNull`
}

type Response struct {
	gorm.Model
	Id       uuid.UUID `gorm:primaryKey;notNull`
	Question Question  `gorm:"foreignKey:id";notNull`
	Slug     string    `gorm:notNull`
	Response string    `gorm:notNull`
}

type Authentication struct {
	Token string `gorm:notNull`
}
