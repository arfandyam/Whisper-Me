package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Username    string
	Firstname   string
	Lastname    string
	Email       string
	Password    *string
	Oauth_id    *string
	Is_oauth    bool
	Is_verified bool
}
