package controllers

import (
	"fmt"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type UserController struct {
	UserService      service.UserServiceInterface
	UserEmailService service.UserEmailServiceInterface
}

func NewUserController(userService service.UserServiceInterface, userEmailService service.UserEmailServiceInterface) UserControllerInterface {
	return &UserController{
		UserService: userService,
		UserEmailService: userEmailService,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	userReq := &dto.UserCreateRequest{}
	// fmt.Println("req.body", ctx.ShouldBindBodyWithJSON(&userReq))
	userResponse := controller.UserEmailService.CreateUserAndSendEmailVerification(ctx, userReq)
	if userResponse == nil {
		return
	}

	userResponse.Response = &dto.Response{
		Status:  "success",
		Message: "Berhasil Menambahkan Data.",
	}

	ctx.JSON(http.StatusCreated, userResponse)
}

func (controller *UserController) EditUser(ctx *gin.Context) {
	userReq := &dto.UserEditRequest{}
	userId := uuid.Must(uuid.Parse(ctx.Param("id")))

	userResponse := controller.UserService.EditUser(ctx, userReq, userId)
	if userResponse == nil {
		return
	}

	userResponse.Response = &dto.Response{
		Status:  "success",
		Message: "Berhasil memperbarui data",
	}

	ctx.JSON(http.StatusOK, userResponse)
}

func (controller *UserController) FindUserById(ctx *gin.Context) {
	userId := uuid.Must(uuid.Parse(ctx.Param("id")))

	userResponse := controller.UserService.FindUserById(ctx, userId)
	if userResponse == nil {
		return
	}

	fmt.Println("userResponse controller", userResponse)

	userResponse.Response = &dto.Response{
		Status:  "success",
		Message: "Berhasil mendapatkan data.",
	}

	ctx.JSON(http.StatusOK, userResponse)
}

func (controller *UserController) ChangePassword(ctx *gin.Context) {
	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	userReq := &dto.UserChangePasswordRequest{}
	controller.UserService.ChangePassword(ctx, userReq, accessToken)

	if len(ctx.Errors) > 0 {
		return
	}

	ctx.JSON(http.StatusOK, &dto.Response{
		Status:  "success",
		Message: "Berhasil memperbarui data.",
	})
}

func (controller *UserController) VerifyUsersEmail(ctx *gin.Context){
	queryToken := ctx.Query("token")
	controller.UserService.VerifyUsersEmail(ctx, queryToken)

	if len(ctx.Errors) > 0 {
		return
	}

	ctx.JSON(http.StatusOK, &dto.Response{
		Status: "success",
		Message: "Berhasil memperbarui data.",
	})
}
