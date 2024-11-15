package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepository repository.UserRepositoryInterface
	TokenManager   tokenize.TokenManagerInterface
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepositoryInterface, tokenManager tokenize.TokenManagerInterface, DB *gorm.DB) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
		TokenManager:   tokenManager,
		DB:             DB,
	}
}

func (service *UserService) CreateUser(ctx *gin.Context, request *dto.UserCreateRequest) *dto.UserCreateResponse {
	// Melakukan validasi berdasarkan UserCreateBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid Request Body", err.Error())
		ctx.Error(err)
		return nil
	}

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

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err = service.UserRepository.CreateUser(tx, user)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User failed to add", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}
	tx.Commit()

	return &dto.UserCreateResponse{
		Id: user.Id,
	}
}

func (service *UserService) EditUser(ctx *gin.Context, request *dto.UserEditRequest, userId uuid.UUID) *dto.UserEditResponse {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
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

	user, err := service.UserRepository.FindUserById(tx, userId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User not found", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	user.Firstname = request.Firstname
	user.Lastname = request.Lastname

	fmt.Println("user dari service sebelum edit", user)

	user, err = service.UserRepository.EditUser(tx, user)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User failed to add", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return &dto.UserEditResponse{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}
}

func (service *UserService) FindUserById(ctx *gin.Context, userId uuid.UUID) *dto.UserFindByIdResponse {
	user, err := service.UserRepository.FindUserById(service.DB, userId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User not found.", err.Error())
		ctx.Error(err)
		return nil
	}

	return &dto.UserFindByIdResponse{
		Id:        user.Id,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
}

func (service *UserService) ChangePassword(ctx *gin.Context, request *dto.UserChangePasswordRequest, accessToken string) {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid request Body", err.Error())
		ctx.Error(err)
		return
	}

	userId, err := service.TokenManager.VerifyToken(accessToken, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "invalid access token", err.Error())
		ctx.Error(err)
		return
	}

	password, err := service.UserRepository.GetUserPassword(service.DB, *userId)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "User not found.", err.Error())
		ctx.Error(err)
		return
	}

	if !libs.CheckPasswordHash(request.Oldpassword, *password) {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Wrong credentials", "Password not matched")
		ctx.Error(err)
		return
	}

	newPassword, err := libs.HashPassword(request.Newpassword)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to hash password", err.Error())
		ctx.Error(err)
		return
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = service.UserRepository.ChangeUserPassword(tx, *userId, newPassword)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to update credentials", err.Error())
		ctx.Error(err)
		return
	}

	tx.Commit()
}
