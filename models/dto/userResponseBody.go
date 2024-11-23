package dto

import "github.com/google/uuid"

type UserEditResponse struct {
	*Response
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

type UserFindByIdResponse struct {
	*Response
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
}
