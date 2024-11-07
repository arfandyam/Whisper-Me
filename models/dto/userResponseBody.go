package dto

import "github.com/google/uuid"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserCreateResponse struct {
	*Response
	Id uuid.UUID `json:"id"`
}

type UserEditResponse struct {
	*Response
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
