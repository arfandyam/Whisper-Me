package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthUserInfo struct {
	Id             uuid.UUID       `json:"id"`
	Username       string          `json:"username"`
	Firstname      string          `json:"first_name"`
	Lastname       string          `json:"last_name"`
	Email          string          `json:"email"`
	Is_oauth       bool            `json:"is_oauth"`
	Is_verified    bool            `json:"is_verified"`
	AccessToken    string          `json:"access_token"`
	AccessTokenIat jwt.NumericDate `json:"access_token_iat"`
	AccessTokenExp jwt.NumericDate `json:"access_token_exp"`
	RefreshToken   string          `json:"refresh_token"`
}

type AuthResponseBody struct {
	*Response
	Data AuthUserInfo `json:"data"`
}

type AccessTokenResponseBody struct {
	*Response
	AccessToken    string          `json:"accesstoken"`
	AccessTokenIat jwt.NumericDate `json:"access_token_iat"`
	AccessTokenExp jwt.NumericDate `json:"access_token_exp"`
}
