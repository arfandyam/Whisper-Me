package dto

import "github.com/golang-jwt/jwt/v5"

type AuthResponseBody struct {
	*Response
	AccessToken    string `json:"access_token"`
	AccessTokenIat jwt.NumericDate `json:"access_token_iat"`
	AccessTokenExp jwt.NumericDate `json:"access_token_exp"`
	RefreshToken   string `json:"refresh_token"`
}

type AccessTokenResponseBody struct {
	*Response
	AccessToken string `json:"accesstoken"`
	AccessTokenIat jwt.NumericDate `json:"access_token_iat"`
	AccessTokenExp jwt.NumericDate `json:"access_token_exp"`
}
