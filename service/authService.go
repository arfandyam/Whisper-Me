package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/arfandyam/Whisper-Me/libs"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/domain"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := service.UserRepository.GetUserCredentials(tx, request.Username)
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to get user credentials", err.Error())
		tx.Rollback()
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
	accessToken, accessTokenIat, accessTokenExp, err := service.TokenManager.GenerateToken(user.Id, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))

	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to sign jwt", err.Error())
		ctx.Error(err)
		return nil
	}

	fmt.Println("testtt2")
	refreshTokenAge, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_AGE"))
	refreshToken, refreshTokenIat, refreshTokenExp, err := service.TokenManager.GenerateToken(user.Id, refreshTokenAge, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Refresh token not found.", err.Error())
		ctx.Error(err)
		return nil
	}

	fmt.Println("testtt3")

	if err := service.AuthRepository.InsertRefreshToken(tx, user.Id, refreshToken, refreshTokenIat.Time, refreshTokenExp.Time); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to store token", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}
	tx.Commit()

	return &dto.AuthResponseBody{
		Data: dto.AuthUserInfo{
			Id: user.Id,
			Username: user.Username,
			Firstname: user.Firstname,
			Lastname: user.Lastname,
			Email: user.Email,
			Is_oauth: user.Is_oauth,
			Is_verified: user.Is_verified,
			AccessToken: accessToken,
			AccessTokenIat: *accessTokenIat,
			AccessTokenExp: *accessTokenExp,
			RefreshToken: refreshToken,
		},
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

	claimsId, err := service.TokenManager.VerifyToken(request.RefreshToken, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Refresh token not found.", err.Error())
		ctx.Error(err)
		return nil
	}

	userId, _ := uuid.Parse(claimsId)

	accessTokenAge, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_AGE"))
	accessToken, iat, exp, err := service.TokenManager.GenerateToken(userId, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Refresh token not found.", err.Error())
		ctx.Error(err)
		return nil
	}

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
		tx.Rollback()
		ctx.Error(err)
		return
	}
	
	tx.Commit()
}

func (service *AuthService) OauthLoginUser(ctx *gin.Context, request *dto.UserCreateOauthRequest) *dto.AuthResponseBody {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
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
	user, _ := service.UserRepository.FindUserByEmail(tx, request.Email)
	fmt.Println("user setelah finduserbyemail", user)
	if user == nil {
		userId := uuid.New()
		user = &domain.User{
			Id:          userId,
			Username:    strings.Join([]string{uuid.NewString(), request.GivenName}, ""),
			Firstname:   request.GivenName,
			Lastname:    request.GivenName,
			Email:       request.Email,
			Password:    nil,
			Oauth_id:    &request.Sub,
			Is_oauth:    true,
			Is_verified: true,
		}

		fmt.Println("user dalam if", user)
		_, err := service.UserRepository.CreateUser(tx, user)
		if err != nil {
			fmt.Println("eaaaa")
			err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to create user", err.Error())
			tx.Rollback()
			ctx.Error(err)
			return nil
		}
	}

	fmt.Println("user luar if", user)
	accessTokenAge, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_AGE"))
	accessToken, accessTokenIat, accessTokenExp, err := service.TokenManager.GenerateToken(user.Id, accessTokenAge, os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to sign jwt", err.Error())
		ctx.Error(err)
		return nil
	}

	fmt.Println("testtt2")
	refreshTokenAge, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_AGE"))
	refreshToken, refreshTokenIat, refreshTokenExp, err := service.TokenManager.GenerateToken(user.Id, refreshTokenAge, os.Getenv("REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to sign jwt", err.Error())
		ctx.Error(err)
		return nil
	}

	if err := service.AuthRepository.InsertRefreshToken(tx, user.Id, refreshToken, refreshTokenIat.Time, refreshTokenExp.Time); err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to store token", err.Error())
		tx.Rollback()
		ctx.Error(err)
		return nil
	}

	tx.Commit()

	return &dto.AuthResponseBody{
		Data: dto.AuthUserInfo{
			Id: user.Id,
			Username: user.Username,
			Firstname: user.Firstname,
			Lastname: user.Lastname,
			Email: user.Email,
			AccessToken: accessToken,
			AccessTokenIat: *accessTokenIat,
			AccessTokenExp: *accessTokenExp,
			RefreshToken: refreshToken,
		},
	}
}
