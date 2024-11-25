package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type ResponseServiceInterface interface {
	CreateResponse(ctx *gin.Context, request *dto.CreateEditAnswerRequestBody) *dto.CreateEditAnswerResponse
}