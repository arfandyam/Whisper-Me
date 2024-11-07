package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context, request *dto.UserCreateRequest) *dto.UserCreateResponse
	EditUser(ctx *gin.Context, request *dto.UserEditRequest, userId uuid.UUID) *dto.UserEditResponse
	FindUserById(ctx *gin.Context, userId uuid.UUID) *dto.UserFindByIdResponse
	// ChangePassword(ctx *gin.Context, request web.UserChangePasswordRequest)
}
