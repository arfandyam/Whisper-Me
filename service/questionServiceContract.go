package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
)

type QuestionServiceInterface interface {
	CreateQuestion(ctx *gin.Context, accessToken string, request *dto.CreateQuestionRequest) *dto.CreateResponse
}