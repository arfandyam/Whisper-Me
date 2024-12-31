package service

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/arfandyam/Whisper-Me/models/dto"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserEmailService struct {
	UserService  UserServiceInterface
	EmailService EmailServiceInterface
	TokenManager tokenize.TokenManagerInterface
	DB           *gorm.DB
}

func NewUserEmailService(userService UserServiceInterface, emailService EmailServiceInterface, tokenManager tokenize.TokenManagerInterface, DB *gorm.DB) UserEmailServiceInterface {
	return &UserEmailService{
		UserService:  userService,
		EmailService: emailService,
		TokenManager: tokenManager,
		DB:           DB,
	}
}

func (service *UserEmailService) CreateUserAndSendEmailVerification(ctx *gin.Context, request *dto.UserCreateRequest) *dto.CreateResponse {
	
	start := time.Now()
	defer func() {
		log.Printf("Total time taken for user creation: %v", time.Since(start))
	}()
	
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

	createUserStart := time.Now()
	userResponse := service.UserService.CreateUser(ctx, tx, request)
	if userResponse == nil {
		tx.Rollback()
		return nil
	}
	log.Printf("Time taken to create user: %v", time.Since(createUserStart))

	tokenGenStart := time.Now()
	var receiver []string
	emailVerificationTokenAge, _ := strconv.Atoi(os.Getenv("EMAIL_TOKEN_AGE"))
	emailVerificationToken, emailVerificationTokenIat, emailVerificationTokenExp, err := service.TokenManager.GenerateToken(request.Email, emailVerificationTokenAge, os.Getenv("EMAIL_TOKEN_SECRET_KEY"))
	if err != nil {
		err := exceptions.NewCustomError(http.StatusBadRequest, "Failed to generate token", err.Error())
		ctx.Error(err)
		tx.Rollback()
		return nil
	}
	log.Printf("Time taken to generate token: %v", time.Since(tokenGenStart))

	emailLinkStart := time.Now()
	verificationLink := os.Getenv("PROTOCOL") + "://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/user/verification?token=" + emailVerificationToken
	emailProperties := &dto.EmailVerificationProperties{
		ToEmail:          append(receiver, request.Email),
		Subject:          "WhisperMe Email Verification",
		VerificationLink: verificationLink,
		IssuedAt:         emailVerificationTokenIat.Time.Format("Monday, 02 January 2006"),
		ExpiredAt:        emailVerificationTokenExp.Time.Format("Monday, 02 January 2006"),
	}
	log.Printf("Time taken to prepare email verification link: %v", time.Since(emailLinkStart))

	emailSendStart := time.Now()
	err = service.EmailService.SendEmailVerification(ctx, emailProperties)
	if err != nil {
		tx.Rollback()
		return nil
	}
	log.Printf("Time taken to send email: %v", time.Since(emailSendStart))

	tx.Commit()

	return userResponse
}
