package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context, tx *gorm.DB, request *dto.UserCreateRequest) *dto.CreateResponse
	EditUser(ctx *gin.Context, request *dto.UserEditRequest, userId uuid.UUID) *dto.UserEditResponse
	FindUserById(ctx *gin.Context, userId uuid.UUID) *dto.UserFindByIdResponse
	ChangePassword(ctx *gin.Context, request *dto.UserChangePasswordRequest, accessToken string)
	VerifyUsersEmail(ctx *gin.Context, queryToken string)
}
