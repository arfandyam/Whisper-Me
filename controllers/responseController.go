package controllers

import (
	"net/http"
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
	questionId := uuid.Must(uuid.Parse(ctx.Param("questionId")))
	request := &dto.CreateAnswerRequestBody{}
	answerResponse := controller.ResponseService.CreateResponse(ctx, request, questionId)

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
	cursorUrl := ctx.Query("cursor")

	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	answerResponse := controller.ResponseService.FindResponseByQuestionId(ctx, questionId, cursorUrl, accessToken)

	if len(ctx.Errors) > 0 || answerResponse == nil {
		return
	}

	answerResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil mendapatkan data",
	}

	ctx.JSON(http.StatusOK, answerResponse)
}

func (controller *ResponseController) SearchResponsesByKeyword(ctx *gin.Context){
	questionId := uuid.MustParse(ctx.Param("questionId"))
	keyword := ctx.Query("keyword")
	rankQuery := ctx.Query("rank")

	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	answerResponse := controller.ResponseService.SearchResponsesByKeyword(ctx, questionId, accessToken, keyword, rankQuery)

	if len(ctx.Errors) > 0 || answerResponse == nil {
		return
	}

	answerResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil mendapatkan data",
	}

	ctx.JSON(http.StatusOK, answerResponse)
}