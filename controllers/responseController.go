package controllers

import (
	"net/http"

	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
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