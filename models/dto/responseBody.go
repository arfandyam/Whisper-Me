package dto

import "github.com/google/uuid"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CreateResponse struct {
	*Response
	Id uuid.UUID `json:"id"`
}
