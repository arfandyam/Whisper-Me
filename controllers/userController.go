package controllers

import (
	"github.com/arfandyam/Whisper-Me/models/web"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	userReq := &web.UserCreateRequest{}
	// fmt.Println("req.body", ctx.ShouldBindBodyWithJSON(&userReq))
	userResponse := controller.UserService.CreateUser(ctx, userReq)

	if userResponse.Id == nil {
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Berhasil menambahkan data.",
		"data":    userResponse.Id,
	})
}
