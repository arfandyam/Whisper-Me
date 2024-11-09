package dto

type AuthResponseBody struct {
	*Response
	AccessToken  string `json:"accesstoken"`
	RefreshToken string	`json:"refreshtoken"`
}
