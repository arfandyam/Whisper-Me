package service

import (
	"github.com/arfandyam/Whisper-Me/models/web"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context, request *web.UserCreateRequest) *web.UserCreateResponse
	// EditUser(ctx *gin.Context, userId uuid.UUID)
	// FindUserById(ctx *gin.Context, userId uuid.UUID)
	// ChangePassword(ctx *gin.Context, request web.UserChangePasswordRequest)
}
