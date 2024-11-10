package dto

import "github.com/golang-jwt/jwt/v5"

type AuthResponseBody struct {
	*Response
	AccessToken    string `json:"accesstoken"`
	AccessTokenIat jwt.NumericDate `json:"accesstokeniat"`
	AccessTokenExp jwt.NumericDate `json:"accesstokenexp"`
	RefreshToken   string `json:"refreshtoken"`
}

type AccessTokenResponseBody struct {
	*Response
	AccessToken string `json:"accesstoken"`
	AccessTokenIat jwt.NumericDate `json:"accesstokeniat"`
	AccessTokenExp jwt.NumericDate `json:"accesstokenexp"`
}
