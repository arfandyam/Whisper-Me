package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type EmailServiceInterface interface {
	SendEmailVerification(ctx *gin.Context, emailProperties *dto.EmailVerificationProperties) error
}