package main

import (
	"github.com/arfandyam/Whisper-Me/config"
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/router"
	"github.com/arfandyam/Whisper-Me/service"
	"github.com/arfandyam/Whisper-Me/tokenize"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	// initializers.ConnDB()
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *exceptions.CustomError:
				ctx.AbortWithStatusJSON(e.Status, gin.H{
					"status":      "failed",
					"description": e.Description,
					"message":     e.Message,
				})
			default:
				ctx.JSON(500, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
			}
		}

		ctx.Abort()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	db := initializers.ConnDB()
	r := gin.Default()

	r.Use(CorsMiddleware())
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
	responseService := service.NewResponseService(responseRepository, questionRepository, tokenManager, db)
	responseController := controllers.NewResponseController(responseService)

	router.ResponseRoute(r, responseController)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
