package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	LoginUser(ctx *gin.Context, request *dto.AuthRequestBody) *dto.AuthResponseBody
	
}