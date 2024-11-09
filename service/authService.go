package service

import (
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

type AuthService struct {
	AuthRepository repository.AuthRepositoryInterface
	UserRepository repository.UserRepositoryInterface
	TokenManager   tokenize.TokenManagerInterface
	DB             *gorm.DB
}

func NewAuthService(authRepository repository.AuthRepositoryInterface, userRepository repository.UserRepositoryInterface, tokenManager tokenize.TokenManagerInterface, DB *gorm.DB) AuthServiceInterface {
	return &AuthService{
		AuthRepository: authRepository,
		UserRepository: userRepository,
		TokenManager:   tokenManager,
		DB:             DB,
	}
}

func (service *AuthService) LoginUser(ctx *gin.Context, request *dto.AuthRequestBody) *dto.AuthResponseBody {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid request body", err.Error())
		ctx.Error(err)
		return nil
	}

	user := &domain.User{
		Username: request.Username,
		Password: &request.Password,
	}

	if err := service.UserRepository.VerifyUserCredentials(service.DB, user); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Credentials not matched", err.Error())
		ctx.Error(err)
		return nil
	}

	accessTokenAge, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_AGE"))
	accessToken, err := service.TokenManager.GenerateToken(user.Id, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to sign jwt", err.Error())
		ctx.Error(err)
		return nil
	}

	refreshTokenAge, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_AGE"))
	refreshToken, err := service.TokenManager.GenerateToken(user.Id, refreshTokenAge, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))

	tx := service.DB.Begin()
	defer func(){
		if r:=recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := service.AuthRepository.LoginUser(tx, &refreshToken); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to store token", err.Error())
		ctx.Error(err)
		return nil
	}
	tx.Commit()

	return &dto.AuthResponseBody{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

// func UpdateAccessToken(ctx *gin.Context)
