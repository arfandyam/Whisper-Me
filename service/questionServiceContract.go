package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuestionServiceInterface interface {
	CreateQuestion(ctx *gin.Context, accessToken string, request *dto.CreateEditQuestionRequest) *dto.CreateEditQuestionResponse
	EditQuestion(ctx *gin.Context, accessToken string, questionId uuid.UUID, request *dto.CreateEditQuestionRequest) *dto.CreateEditQuestionResponse
	FindQuestionById(ctx *gin.Context, accessToken string, questionId uuid.UUID) *dto.FindQuestionResponse
	FindQuestionsByUserId(ctx *gin.Context, accessToken string, cursorUrl string) *dto.FindQuestionsByUserIdResponse
	SearchQuestionsByKeyword(ctx *gin.Context, accessToken string, cursorUrl string, keyword string) *dto.FindQuestionsByUserIdResponse
}