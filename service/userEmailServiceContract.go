package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type UserEmailServiceInterface interface {
	CreateUserAndSendEmailVerification(ctx *gin.Context, request *dto.UserCreateRequest) *dto.UserCreateResponse
}