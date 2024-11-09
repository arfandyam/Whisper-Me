package tokenize

import (
	"github.com/google/uuid"
)

type TokenManagerInterface interface {
	GenerateToken(id uuid.UUID, tokenAge int, secretKey string) (string, error)
	// GenerateRefreshToken(id uuid.UUID) (string, error)
}