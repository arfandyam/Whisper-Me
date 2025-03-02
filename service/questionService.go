package service

import (
	"net/http"
	"os"
	"strconv"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/libs/mapper"
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

func (service *QuestionService) CreateQuestion(ctx *gin.Context, accessToken string, request *dto.CreateEditQuestionRequest) *dto.CreateEditQuestionResponse {
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
		UserId:   userId,
		Slug:     slug,
		Topic:    request.Topic,
		Question: request.Question,
		UrlKey:   libs.SlugToBase62(slug),
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	question, err = service.QuestionRepository.CreateQuestion(tx, question)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "failed to store question", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &dto.CreateEditQuestionResponse{
		Id:       question.Id,
		UserId:   question.UserId,
		Slug:     question.Slug,
		Topic:    question.Topic,
		Question: question.Question,
	}
}

func (service *QuestionService) EditQuestion(ctx *gin.Context, accessToken string, questionId uuid.UUID, request *dto.CreateEditQuestionRequest) *dto.CreateEditQuestionResponse {
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

	slug := libs.ToSlug(request.Topic, questionId)
	question := &domain.Question{
		Id:       questionId,
		UserId:   userId,
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
	question, err = service.QuestionRepository.EditQuestion(tx, question)

	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to edit item", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &dto.CreateEditQuestionResponse{
		Id:       question.Id,
		Slug:     question.Slug,
		Topic:    question.Topic,
		Question: question.Question,
	}
}

func (service *QuestionService) FindQuestionById(ctx *gin.Context, accessToken string, questionId uuid.UUID) *dto.FindQuestionResponse {
	_, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return nil
	}

	question, err := service.QuestionRepository.FindQuestionById(service.DB, questionId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "failed to fetch data", err.Error())
		ctx.Error(err)
		return nil
	}

	return &dto.FindQuestionResponse{
		Data: dto.QuestionDTO{
			Id:        question.Id,
			UserId:    question.UserId,
			Slug:      question.Slug,
			Topic:     question.Topic,
			Question:  question.Question,
			UrlKey:    question.UrlKey,
			CreatedAt: question.CreatedAt,
		},
	}
}

func (service *QuestionService) FindQuestionBySlug(ctx *gin.Context, accessToken string, slug string) *dto.FindQuestionResponse {
	_, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return nil
	}

	question, err := service.QuestionRepository.FindQuestionBySlug(service.DB, slug)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "failed to fetch data", err.Error())
		ctx.Error(err)
		return nil
	}

	return &dto.FindQuestionResponse{
		Data: dto.QuestionDTO{
			Id:        question.Id,
			UserId:    question.UserId,
			Slug:      question.Slug,
			Topic:     question.Topic,
			Question:  question.Question,
			UrlKey:    question.UrlKey,
			CreatedAt: question.CreatedAt,
		},
	}
}

func (service *QuestionService) FindQuestionsByUserId(ctx *gin.Context, accessToken string, cursorUrl string) *dto.FindQuestionsByUserIdResponse {
	claimsId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return nil
	}

	userId := uuid.Must(uuid.Parse(claimsId))

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

	fetchPerPage, _ := strconv.Atoi(os.Getenv("FETCH_PER_PAGE"))
	questions, err := service.QuestionRepository.FindQuestionsByUserId(service.DB, userId, cursor, fetchPerPage)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch questions", err.Error())
		ctx.Error(err)
		return nil
	}

	prevCursor, err := service.QuestionRepository.FindPrevCursorQuestion(service.DB, userId, cursor, fetchPerPage)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch prevCursor", err.Error())
		ctx.Error(err)
		return nil
	}

	var questionsDTO []dto.QuestionDTO
	var nextCursor *uuid.UUID
	if len(questions) <= fetchPerPage {
		nextCursor = nil
		for i := 0; i <= len(questions)-1; i++ {
			questionsDTO = append(questionsDTO, mapper.MapQuestionDomainToQuestionDTO(questions[i]))
		}
	} else {
		nextCursor = &questions[len(questions)-1].Id
		for i := 0; i <= fetchPerPage-1; i++ {
			questionsDTO = append(questionsDTO, mapper.MapQuestionDomainToQuestionDTO(questions[i]))
		}
	}

	return &dto.FindQuestionsByUserIdResponse{
		Data: questionsDTO,
		Meta: dto.PageCursorInfo{
			NextCursor: nextCursor,
			PrevCursor: prevCursor,
		},
	}
}

func (service *QuestionService) SearchQuestionsByKeyword(ctx *gin.Context, accessToken string, keyword string, rankQuery string) *dto.SearchKeywordQuestionsByUserIdResponse {
	claimsId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return nil
	}

	userId := uuid.Must(uuid.Parse(claimsId))

	var rank *float64
	if rankQuery != "" {
		parsedRank, _ := strconv.ParseFloat(rankQuery, 64)
		rank = &parsedRank
	}

	fetchPerPage, _ := strconv.Atoi(os.Getenv("FETCH_PER_PAGE"))

	questions, err := service.QuestionRepository.SearchQuestionsByKeyword(service.DB, userId, fetchPerPage, keyword, rank)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch questions", err.Error())
		ctx.Error(err)
		return nil
	}

	prevRank, err := service.QuestionRepository.FindPrevRankQuestion(service.DB, userId, fetchPerPage, keyword, rank)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to fetch prev rank", err.Error())
		ctx.Error(err)
		return nil
	}

	var questionsDTO []dto.QuestionDTO
	var nextRank *float64
	if len(questions) <= fetchPerPage {
		nextRank = nil
		for i := 0; i < len(questions); i++ {
			questionsDTO = append(questionsDTO, mapper.MapQuestionDomainToQuestionDTO(questions[i]))
		}
	} else {
		nextRank = &questions[len(questions)-1].Rank
		for i := 0; i < fetchPerPage; i++ {
			questionsDTO = append(questionsDTO, mapper.MapQuestionDomainToQuestionDTO(questions[i]))
		}
	}

	return &dto.SearchKeywordQuestionsByUserIdResponse{
		Data: questionsDTO,
		Meta: dto.PageRankInfo{
			NextRank: nextRank,
			PrevRank: prevRank,
		},
	}
}
