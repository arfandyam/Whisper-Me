package service

import (
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseServiceInterface interface {
	CreateResponse(ctx *gin.Context, request *dto.CreateAnswerRequestBody, questionId uuid.UUID) *dto.CreateAnswerResponse
	FindResponseByQuestionId(ctx *gin.Context, questionId uuid.UUID, accessToken string, cursorUrl string) *dto.FindAnswerResponse
	SearchResponsesByKeyword(ctx *gin.Context, questionId uuid.UUID, accessToken string, keyword string, rankQuery string) *dto.SearchKeywordResponseByUserIdResponse
}