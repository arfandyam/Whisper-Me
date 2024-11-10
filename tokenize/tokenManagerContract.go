package tokenize

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenManagerInterface interface {
	GenerateToken(id uuid.UUID, tokenAge int, secretKeyString string) (string, *jwt.NumericDate, *jwt.NumericDate,  error)
	VerifyToken(tokenString string, secretKeyString string) (*uuid.UUID, error)
}