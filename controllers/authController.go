package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arfandyam/Whisper-Me/config"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthController struct {
	AuthService service.AuthServiceInterface
	AppConfig   config.AppOauthConfigInterface
}

func NewAuthController(authService service.AuthServiceInterface, appConfig config.AppOauthConfigInterface) AuthControllerInterface {
	return &AuthController{
		AuthService: authService,
		AppConfig:   appConfig,
	}
}

func (controller *AuthController) LoginUser(ctx *gin.Context) {
	authReq := &dto.AuthRequestBody{}

	authResponse := controller.AuthService.LoginUser(ctx, authReq)

	if authResponse == nil {
		return
	}

	authResponse.Response = &dto.Response{
		Status:  "success",
		Message: "berhasil login",
	}

	ctx.JSON(http.StatusOK, authResponse)
}

func (controller *AuthController) UpdateAccessToken(ctx *gin.Context) {
	authReq := &dto.RefreshTokenRequestBody{}

	authResponse := controller.AuthService.UpdateAccessToken(ctx, authReq)

	if authResponse == nil {
		return
	}

	ctx.JSON(http.StatusOK, authResponse)
}

func (controller *AuthController) LogoutUser(ctx *gin.Context) {
	authReq := &dto.RefreshTokenRequestBody{}

	controller.AuthService.LogoutUser(ctx, authReq)

	if len(ctx.Errors) > 0 {
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "Berhasil logout",
	})
}

func (controller *AuthController) UserGoogleOauthLogin(ctx *gin.Context) {
	googleConfig := controller.AppConfig.GoogleConfig()
	url := googleConfig.AuthCodeURL("randomstate", oauth2.SetAuthURLParam("prompt", "select_account"))

	ctx.Redirect(http.StatusTemporaryRedirect, url)
	ctx.JSON(http.StatusTemporaryRedirect, gin.H{
		"message": "redirect to google login page",
	})
}

func (controller *AuthController) GoogleOauthCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	if state != "randomstate" {
		err := exceptions.NewCustomError(http.StatusBadRequest, "query not valid", "state query not matched")
		ctx.Error(err)
		return
	}

	googleConfig := controller.AppConfig.GoogleConfig()

	code := ctx.Query("code")
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "query not valid", "state query not matched")
		ctx.Error(err)
		return
	}

	client := googleConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "query not valid", "state query not matched")
		ctx.Error(err)
		return
	}

	userInfo := &dto.UserCreateOauthRequest{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "query not valid", "state query not matched")
		ctx.Error(err)
		return
	}

	fmt.Println("userInfo", userInfo)

	authResponse := controller.AuthService.OauthLoginUser(ctx, userInfo)

	if len(ctx.Errors) > 0 {
		return
	}

	ctx.JSON(http.StatusOK, authResponse)
}
