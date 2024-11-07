package controllers

import (
	"net/http"

	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) UserControllerInterface {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	userReq := &dto.UserCreateRequest{}
	// fmt.Println("req.body", ctx.ShouldBindBodyWithJSON(&userReq))
	userResponse := controller.UserService.CreateUser(ctx, userReq)
	if userResponse == nil {
		return
	}

	userResponse.Response = &dto.Response{
		Status:  "success",
		Message: "Berhasil Menambahkan Data.",
	}

	ctx.JSON(http.StatusCreated, userResponse)
}

func (controller *UserController) EditUser(ctx *gin.Context){
	userReq := &dto.UserEditRequest{}
	userId := uuid.Must(uuid.Parse(ctx.Param("id")))

	userResponse := controller.UserService.EditUser(ctx, userReq, userId)
	if userResponse == nil {
		return
	}

	userResponse.Response = &dto.Response{
		Status: "success",
		Message: "Berhasil memperbarui data",
	}

	ctx.JSON(http.StatusOK, userResponse)
}
