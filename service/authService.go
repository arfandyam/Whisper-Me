package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	user, err := service.UserRepository.GetUserCredentials(service.DB, request.Username)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to get user credentials", err.Error())
		ctx.Error(err)
		return nil
	}

	if !libs.CheckPasswordHash(request.Password, *user.Password) {
		err := exceptions.NewCustomError(http.StatusBadRequest, "credentials not matched", "wrong password")
		ctx.Error(err)
		return nil
	}

	fmt.Println("testtt1")
	accessTokenAge, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_AGE"))
	accessToken, iat, exp, err := service.TokenManager.GenerateToken(user.Id, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to sign jwt", err.Error())
		ctx.Error(err)
		return nil
	}

	fmt.Println("testtt2")
	refreshTokenAge, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_AGE"))
	refreshToken, _, _, err := service.TokenManager.GenerateToken(user.Id, refreshTokenAge, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))

	fmt.Println("testtt3")
	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
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
		AccessToken:    accessToken,
		AccessTokenIat: *iat,
		AccessTokenExp: *exp,
		RefreshToken:   refreshToken,
	}
}

func (service *AuthService) UpdateAccessToken(ctx *gin.Context, request *dto.RefreshTokenRequestBody) *dto.AccessTokenResponseBody {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid Request Body", err.Error())
		ctx.Error(err)
		return nil
	}

	if err := service.AuthRepository.VerifyRefreshToken(service.DB, &request.RefreshToken); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Refresh token not found.", err.Error())
		ctx.Error(err)
		return nil
	}

	userId, err := service.TokenManager.VerifyToken(request.RefreshToken, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Refresh token not found.", err.Error())
		ctx.Error(err)
		return nil
	}

	accessTokenAge, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_AGE"))
	accessToken, iat, exp, err := service.TokenManager.GenerateToken(*userId, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	return &dto.AccessTokenResponseBody{
		AccessToken:    accessToken,
		AccessTokenIat: *iat,
		AccessTokenExp: *exp,
	}
}

func (service *AuthService) LogoutUser(ctx *gin.Context, request *dto.RefreshTokenRequestBody) {
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Invalid Request Body", err.Error())
		ctx.Error(err)
		return
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := service.AuthRepository.DeleteRefreshToken(tx, request.RefreshToken); err != nil {
		fmt.Println("eaeaeaea")
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to delete refresh token", err.Error())
		// tx.Rollback()
		ctx.Error(err)
		return
	}

	fmt.Println("bassanggg")
	tx.Commit()
}
