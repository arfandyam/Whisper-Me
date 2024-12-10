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

type ResponseController struct {
	ResponseService service.ResponseServiceInterface
}

func NewResponseController(responseService service.ResponseServiceInterface) ResponseControllerInterface {
	return &ResponseController{
		ResponseService: responseService,
	}
}

func (controller *ResponseController) CreateResponse(ctx *gin.Context){
	request := &dto.CreateEditAnswerRequestBody{}
	answerResponse := controller.ResponseService.CreateResponse(ctx, request)

	if len(ctx.Errors) > 0 {
		return
	}

	answerResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil menambahkan data.",
	}

	ctx.JSON(http.StatusCreated, answerResponse)
}

func (controller *ResponseController) FindResponseByQuestionId(ctx *gin.Context){
	questionId := uuid.MustParse(ctx.Param("questionId"))
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "invalid page query",
		})
	}

	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	answerResponse := controller.ResponseService.FindResponseByQuestionId(ctx, questionId, page, accessToken)

	if len(ctx.Errors) > 0 || answerResponse == nil {
		return
	}

	answerResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil mendapatkan data",
	}

	ctx.JSON(http.StatusOK, answerResponse)
}