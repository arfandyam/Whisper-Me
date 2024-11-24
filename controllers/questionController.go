package controllers

import (
	// "fmt"
	"net/http"
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

