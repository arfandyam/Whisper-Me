package tokenize

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManager struct{}

func NewTokenManager() TokenManagerInterface {
	return &TokenManager{}
}

func (tokenManager *TokenManager) GenerateToken(id uuid.UUID, tokenAge int, secretKeyString string) (string, *jwt.NumericDate, *jwt.NumericDate, error) {
	secretKey := []byte(secretKeyString)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(tokenAge))),
			"iat": jwt.NewNumericDate(time.Now()),
		})
	tokenString, err := token.SignedString(secretKey)
	claims, _ := token.Claims.(jwt.MapClaims)
	exp := claims["exp"].(*jwt.NumericDate)
	iat := claims["iat"].(*jwt.NumericDate)
	if err != nil  {
		return "", nil, nil, err
	}

	return tokenString, iat, exp, nil
}

func (tokenManager *TokenManager) VerifyToken(tokenString string, secretKeyString string) (*uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKeyString), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			id, err := uuid.Parse(id)
			if err != nil {
				return nil, err
			}

			return &id, nil
		}
	}

	return nil, fmt.Errorf("invalid access token")
}
