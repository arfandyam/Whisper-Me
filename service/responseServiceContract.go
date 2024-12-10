package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseServiceInterface interface {
	CreateResponse(ctx *gin.Context, request *dto.CreateEditAnswerRequestBody) *dto.CreateEditAnswerResponse
	FindResponseByQuestionId(ctx *gin.Context, questionId uuid.UUID, page int, accessToken string) *dto.FindAnswerResponse
}