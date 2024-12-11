package service

import (
	"net/http"
	"os"
	"strconv"

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

type ResponseService struct {
	ResponseRepository repository.ResponseRepositoryInterface
	QuestionRepository repository.QuestionRepositoryInterface
	TokenManager       tokenize.TokenManagerInterface
	DB                 *gorm.DB
}

func NewResponseService(responseRepository repository.ResponseRepositoryInterface, questionRepository repository.QuestionRepositoryInterface, tokenManager tokenize.TokenManagerInterface, DB *gorm.DB) ResponseServiceInterface {
	return &ResponseService{
		ResponseRepository: responseRepository,
		QuestionRepository: questionRepository,
		TokenManager:       tokenManager,
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
			Id:         response.Id,
			QuestionId: response.QuestionId,
			Response:   response.Response,
		},
	}
}

func (service *ResponseService) FindResponseByQuestionId(ctx *gin.Context, questionId uuid.UUID, page int, accessToken string) *dto.FindAnswerResponse {
	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userId, err := service.QuestionRepository.FindQuestionOwner(tx, questionId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusNotFound, "Failed to fetch user id", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}

	claimsId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to verify token", err.Error())
		ctx.Error(err)
		return nil
	}

	userIdFromClaims := uuid.MustParse(claimsId)

	if userIdFromClaims != *userId {
		err := exceptions.NewCustomError(http.StatusUnauthorized, "Unauthorized", "Source not allowed to access")
		ctx.Error(err)
		return nil
	}

	fetchPerPage := 10
	offset := (page - 1) * fetchPerPage

	responses, err := service.ResponseRepository.FindResponseByQuestionId(tx, questionId, fetchPerPage, offset)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch responses/answers", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}

	var responsesDTO []dto.FullResponseDTO
	for i:=0; i < len(responses); i++ {
		responsesDTO = append(responsesDTO, dto.FullResponseDTO{
			ResponseDTO: dto.ResponseDTO{
				Id: responses[i].Id,
				QuestionId: responses[i].QuestionId,
				Response: responses[i].Response,
			},

			ResponseTimestampDTO: dto.ResponseTimestampDTO{
				CreatedAt: responses[i].CreatedAt,
				UpdatedAt: responses[i].UpdatedAt,
			},
		})
	}

	return &dto.FindAnswerResponse{
		Data: responsesDTO,
	}
}

func (service *ResponseService) SearchResponsesByKeyword(ctx *gin.Context, questionId uuid.UUID, keyword string, page int, accessToken string) *dto.FindAnswerResponse {
	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	userId, err := service.QuestionRepository.FindQuestionOwner(tx, questionId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusNotFound, "Failed to fetch user id", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}

	claimsId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to verify token", err.Error())
		ctx.Error(err)
		return nil
	}

	userIdFromClaims := uuid.MustParse(claimsId)

	if userIdFromClaims != *userId {
		err := exceptions.NewCustomError(http.StatusUnauthorized, "Unauthorized", "Source not allowed to access")
		ctx.Error(err)
		return nil
	}

	fetchPerPage, _ := strconv.Atoi(os.Getenv("FETCH_PER_PAGE"))
	responses, err := service.ResponseRepository.SearchResponsesByKeyword(tx, keyword, questionId, fetchPerPage, libs.CalculateOffset(page, fetchPerPage))

	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch responses/answers", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}

	var responsesDTO []dto.FullResponseDTO
	for i:=0; i < len(responses); i++ {
		responsesDTO = append(responsesDTO, dto.FullResponseDTO{
			ResponseDTO: dto.ResponseDTO{
				Id: responses[i].Id,
				QuestionId: responses[i].QuestionId,
				Response: responses[i].Response,
			},

			ResponseTimestampDTO: dto.ResponseTimestampDTO{
				CreatedAt: responses[i].CreatedAt,
				UpdatedAt: responses[i].UpdatedAt,
			},
		})
	}

	return &dto.FindAnswerResponse{
		Data: responsesDTO,
	}
}


