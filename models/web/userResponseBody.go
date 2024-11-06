package web

import "github.com/google/uuid"

type UserCreateResponse struct {
	Status  string    `json: "status"`
	Message string    `json: "message"`
	Id      *uuid.UUID `json: "id"`
}
