package controllers

import (
	"net/http"
	"strings"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
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
	request := &dto.CreateQuestionRequest{}

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