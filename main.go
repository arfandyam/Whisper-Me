package main

import (
	"github.com/arfandyam/Whisper-Me/config"
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/router"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	// initializers.ConnDB()
}

func main() {

	db := initializers.ConnDB()
	r := gin.Default()


	r.Use(GlobalErrorHandler())

	//App Oauth Config
	appConfig := config.NewAppOauthConfig()

	//Token Manager
	tokenManager := tokenize.NewTokenManager()

	//Email Service
	emailService := service.NewEmailService()

	// User
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, tokenManager, db)
	userEmailService := service.NewUserEmailService(userService, emailService, tokenManager, db) //user and email service intermediary
	userController := controllers.NewUserController(userService, userEmailService)

	router.UserRoutes(r, userController)

	// Auth
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, userRepository, tokenManager, db)
	authController := controllers.NewAuthController(authService, appConfig)

	router.AuthRoutes(r, authController)

	//Question
	questionRepository := repository.NewQuestionRepository()
	questionService := service.NewQuestionService(&questionRepository, tokenManager, db)
	questionController := controllers.NewQuestionController(questionService)

	router.QuestionRoutes(r, questionController)

	//Response
	responseRepository := repository.NewResponseRepository()
	responseService := service.NewResponseService(responseRepository, db)
	responseController := controllers.NewResponseController(responseService)

	router.ResponseRoute(r, responseController)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}