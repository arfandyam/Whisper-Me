package controllers

import (
	"net/http"

	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthServiceInterface
}

func NewAuthController(authService service.AuthServiceInterface) AuthControllerInterface {
	return &AuthController{
		AuthService: authService,
	}
}

func (controller *AuthController) LoginUser(ctx *gin.Context){
	authReq := &dto.AuthRequestBody{}

	authResponse := controller.AuthService.LoginUser(ctx, authReq)

	if authResponse == nil {
		return
	}

	authResponse.Response = &dto.Response{
		Status: "success",
		Message: "berhasil login",
	}

	ctx.JSON(http.StatusOK, authResponse)
}
