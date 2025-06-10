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
	FindQuestionBySlug(ctx *gin.Context, slug string) *dto.FindQuestionResponse
	FindQuestionsByUserId(ctx *gin.Context, accessToken string, cursorUrl string) *dto.FindQuestionsByUserIdResponse
	SearchQuestionsByKeyword(ctx *gin.Context, accessToken string, keyword string, rankQuery string, cursorUrl string) *dto.SearchKeywordQuestionsByUserIdResponse
	FindQuestionSlugByUrlKey(ctx *gin.Context, urlKey string) *string
}
