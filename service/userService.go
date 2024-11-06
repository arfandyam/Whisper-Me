package service

import (
	"fmt"
	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/web"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepositoryInterface, DB *gorm.DB) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *UserService) CreateUser(ctx *gin.Context, request *web.UserCreateRequest) *web.UserCreateResponse {
	// Melakukan validasi berdasarkan UserCreateBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid Request Body", err.Error())
		ctx.Error(err)
		return nil
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create Id
	uuid := uuid.New()

	// Hashing Password
	password, err := libs.HashPassword(request.Password)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to hash password", err.Error())
		ctx.Error(err)
		return nil
	}

	user := &domain.User{
		Id:          uuid,
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		Username:    request.Username,
		Email:       request.Email,
		Password:    &password,
		Oauth_id:    nil,
		Is_oauth:    false,
		Is_verified: false,
	}

	fmt.Println("user dari service", user)
	fmt.Println("user.Password", *user.Password)

	user, err = service.UserRepository.CreateUser(tx, user)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User failed to add", err.Error())
		ctx.Error(err)
	}

	return &web.UserCreateResponse{
		Status: "success",
		Message: "Berhasil Menambahkan Data.",
		Id: &user.Id,
	}
}
