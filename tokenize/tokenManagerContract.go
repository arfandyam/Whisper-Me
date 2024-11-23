package tokenize

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenManagerInterface interface {
	GenerateToken(id interface{}, tokenAge int, secretKeyString string) (string, *jwt.NumericDate, *jwt.NumericDate,  error)
	VerifyToken(tokenString string, secretKeyString string) (string, error)
}