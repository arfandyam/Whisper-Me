package dto

type AuthRequestBody struct {
	Username string `validate:"required" json:"username"`
	Password string	`validate:"required,min=8" json:password`
}