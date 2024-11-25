package service

import (
	"net/http"

	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResponseService struct {
	ResponseRepository repository.ResponseRepositoryInterface
	DB                 *gorm.DB
}

func NewResponseService(responseRepository repository.ResponseRepositoryInterface, DB *gorm.DB) ResponseServiceInterface {
	return &ResponseService{
		ResponseRepository: responseRepository,
		DB:                 DB,
	}
}

func (service *ResponseService) CreateResponse(ctx *gin.Context, request *dto.CreateEditAnswerRequestBody) *dto.CreateEditAnswerResponse {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid request body", err.Error())
		ctx.Error(err)
		return nil
	}

	responseId := uuid.New()
	response := &domain.Response{
		Id:         responseId,
		QuestionId: request.QuestionId,
		Response:   request.Response,
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	response, err := service.ResponseRepository.CreateResponse(tx, response)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to create request", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &dto.CreateEditAnswerResponse{
		Data: dto.ResponseDTO{
			Id: response.Id,
			QuestionId: response.QuestionId,
			Response: response.Response,
		},
	}
}
