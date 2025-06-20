package service

import (
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/libs/mapper"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
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

func (service *ResponseService) CreateResponse(ctx *gin.Context, request *dto.CreateAnswerRequestBody, questionId uuid.UUID) *dto.CreateAnswerResponse {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid request body", err.Error())
		ctx.Error(err)
		return nil
	}

	responseId := uuid.New()
	response := &domain.Response{
		Id:         responseId,
		QuestionId: questionId,
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

	return &dto.CreateAnswerResponse{
		Data: mapper.MapResponseDomainToResponseDTO(*response),
	}
}

func (service *ResponseService) FindResponseByQuestionId(ctx *gin.Context, questionId uuid.UUID, cursorUrl string, accessToken string) *dto.FindAnswerResponse {
	userId, err := service.QuestionRepository.FindQuestionOwner(service.DB, questionId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusNotFound, "Failed to fetch user id", err.Error())
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

	var cursor *uuid.UUID
	if cursorUrl != "" {
		fetchedCursor := uuid.MustParse(cursorUrl)
		cursor = &fetchedCursor
	}

	if userIdFromClaims != *userId {
		err := exceptions.NewCustomError(http.StatusUnauthorized, "Unauthorized", "Source not allowed to access")
		ctx.Error(err)
		return nil
	}

	fetchPerPage, _ := strconv.Atoi(os.Getenv("FETCH_PER_PAGE"))
	responses, err := service.ResponseRepository.FindResponseByQuestionId(service.DB, questionId, fetchPerPage, cursor)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch responses/answers", err.Error())
		ctx.Error(err)
		return nil
	}

	prevCursor, err := service.ResponseRepository.FindPrevCursorResponse(service.DB, questionId, fetchPerPage, cursor)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch previous cursor", err.Error())
		ctx.Error(err)
		return nil
	}

	var nextCursor *uuid.UUID
	var responsesDTO []dto.ResponseDTO
	if len(responses) <= fetchPerPage {
		for i := 0; i < len(responses); i++ {
			responsesDTO = append(responsesDTO, mapper.MapResponseDomainToResponseDTO(responses[i]))
		}
		nextCursor = nil
	} else {
		for i := 0; i < fetchPerPage; i++ {
			responsesDTO = append(responsesDTO, mapper.MapResponseDomainToResponseDTO(responses[i]))
		}
		nextCursor = &responses[len(responses)-1].Id
	}

	return &dto.FindAnswerResponse{
		Data: responsesDTO,
		Meta: dto.PageCursorInfo{
			NextCursor: nextCursor,
			PrevCursor: prevCursor,
		},
	}
}

func (service *ResponseService) SearchResponsesByKeyword(ctx *gin.Context, questionId uuid.UUID, accessToken string, keyword string, rankQuery string, cursorUrl string) *dto.SearchKeywordResponseByUserIdResponse {
	userId, err := service.QuestionRepository.FindQuestionOwner(service.DB, questionId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusNotFound, "Failed to fetch user id", err.Error())
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

	var cursor *uuid.UUID
	if cursorUrl != "" {
		parsedUUID, err := uuid.Parse(cursorUrl)
		if err != nil {
			err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to parse cursor", err.Error())
			ctx.Error(err)
			return nil
		} else {
			cursor = &parsedUUID
		}
	}

	var rank *float64
	if rankQuery != "" {
		parsedRank, _ := strconv.ParseFloat(rankQuery, 64)
		rank = &parsedRank
	}

	fetchPerPage, _ := strconv.Atoi(os.Getenv("FETCH_PER_PAGE"))
	responses, err := service.ResponseRepository.SearchResponsesByKeyword(service.DB, questionId, fetchPerPage, keyword, rank, cursor)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch responses/answers", err.Error())
		ctx.Error(err)
		return nil
	}
	prevResponse, err := service.ResponseRepository.FindPrevRankResponse(service.DB, questionId, fetchPerPage, keyword, rank, cursor)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch prev rank response/answer", err.Error())
		ctx.Error(err)
		return nil
	}

	var responsesDTO []dto.ResponseDTO
	var nextRank *float64
	var nextCursor *uuid.UUID
	if len(responses) <= fetchPerPage {
		for i := 0; i < len(responses); i++ {
			responsesDTO = append(responsesDTO, mapper.MapResponseDomainToResponseDTO(responses[i]))
		}
		nextRank = nil
		nextCursor = nil
	} else {
		for i := 0; i < fetchPerPage; i++ {
			responsesDTO = append(responsesDTO, mapper.MapResponseDomainToResponseDTO(responses[i]))
		}
		nextRank = &responses[len(responses)-1].Rank
		nextCursor = &responses[len(responses)-1].Id
	}

	if prevResponse == nil {
		return &dto.SearchKeywordResponseByUserIdResponse{
			Data: responsesDTO,
			Meta: dto.PageRankInfo{
				NextCursor: nextCursor,
				PrevCursor: nil,
				NextRank:   nextRank,
				PrevRank:   nil,
			},
		}
	} else {
		return &dto.SearchKeywordResponseByUserIdResponse{
			Data: responsesDTO,
			Meta: dto.PageRankInfo{
				NextCursor: nextCursor,
				PrevCursor: &prevResponse.Id,
				NextRank:   nextRank,
				PrevRank:   &prevResponse.Rank,
			},
		}
	}
}
