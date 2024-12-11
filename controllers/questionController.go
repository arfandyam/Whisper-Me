package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QuestionController struct {
	QuestionService service.QuestionServiceInterface
}

func NewQuestionController(questionService service.QuestionServiceInterface) QuestionControllerInterface {
	return &QuestionController{
		QuestionService: questionService,
	}
}

func (controller *QuestionController) CreateQuestion(ctx *gin.Context){
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	request := &dto.CreateEditQuestionRequest{}

	questionResponse := controller.QuestionService.CreateQuestion(ctx, accessToken, request)

	if len(ctx.Errors) > 0 {
		return
	}

	questionResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil menambahkan data.",
	}

	ctx.JSON(http.StatusCreated, questionResponse)
}

func (controller *QuestionController) EditQuestion(ctx *gin.Context){
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	questionId := uuid.Must(uuid.Parse(ctx.Param("id")))

	// fmt.Println("accessToken", accessToken)

	request := &dto.CreateEditQuestionRequest{}

	questionResponse := controller.QuestionService.EditQuestion(ctx, accessToken, questionId, request)

	if len(ctx.Errors) > 0 {
		return
	}

	questionResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil memperbarui data.",
	}

	ctx.JSON(http.StatusOK, questionResponse)
}

func (controller *QuestionController) FindQuestionById(ctx *gin.Context){
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	questionId := uuid.Must(uuid.Parse(ctx.Param("id")))

	// fmt.Println("accessToken", accessToken)

	questionResponse := controller.QuestionService.FindQuestionById(ctx, accessToken, questionId)

	if len(ctx.Errors) > 0 {
		return
	}

	questionResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil memperbarui data.",
	}

	ctx.JSON(http.StatusOK, questionResponse)
}

func (controller *QuestionController) FindQuestionsByUserId(ctx *gin.Context){
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "invalid page query",
		})
	}	

	questionResponse := controller.QuestionService.FindQuestionsByUserId(ctx, accessToken, page)

	if len(ctx.Errors) > 0 {
		return
	}

	questionResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil mendapatkan data",
	}

	ctx.JSON(http.StatusOK, questionResponse)
}

func (controller *QuestionController) SearchQuestionsByKeyword(ctx *gin.Context){
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "invalid page query",
		})
	}

	keyword := ctx.Query("keyword")
	questionResponse := controller.QuestionService.SearchQuestionsByKeyword(ctx, accessToken, page, keyword)

	if len(ctx.Errors) > 0 {
		return
	}

	questionResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil mendapatkan data",
	}

	ctx.JSON(http.StatusOK, questionResponse)
}