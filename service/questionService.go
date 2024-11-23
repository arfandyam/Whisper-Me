package service

import (
	"net/http"
	"os"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionService struct {
	QuestionRepository repository.QuestionRepositoryInterface
	TokenManager       tokenize.TokenManagerInterface
	DB                 *gorm.DB
}

func NewQuestionService(questionRepository *repository.QuestionRepositoryInterface, tokenManager tokenize.TokenManagerInterface, DB *gorm.DB) QuestionServiceInterface {
	return &QuestionService{
		QuestionRepository: *questionRepository,
		TokenManager:       tokenManager,
		DB:                 DB,
	}
}

func (service *QuestionService) CreateQuestion(ctx *gin.Context, accessToken string, request *dto.CreateQuestionRequest) *dto.CreateResponse {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid request body", err.Error())
		ctx.Error(err)
		return nil
	}

	claimsId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return nil
	}

	userId, err := uuid.Parse(claimsId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusInternalServerError, "Failed to parse token", err.Error())
		ctx.Error(err)
		return nil
	}

	questionId := uuid.New()
	slug := libs.ToSlug(request.Topic, questionId)

	question := &domain.Question{
		Id:       questionId,
		Slug:     slug,
		Topic:    request.Topic,
		Question: request.Question,
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	question, err = service.QuestionRepository.CreateQuestion(tx, userId, question)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "failed to store question", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &dto.CreateResponse{
		Id: question.Id,
	}
}
