package tokenize

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct{}

func NewTokenManager() TokenManagerInterface {
	return &TokenManager{}
}

func (tokenManager *TokenManager) GenerateToken(id uuid.UUID, tokenAge int, secretKey string) (string, error) {
	byteKey := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Second * time.Duration(tokenAge)).Unix(),
			"iat": time.Now().Unix(),
		})
	tokenString, err := token.SignedString(byteKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
