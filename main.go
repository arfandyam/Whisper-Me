package main

import (
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/repository"
	"github.com/arfandyam/Whisper-Me/router/user"
	"github.com/arfandyam/Whisper-Me/service"
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

	// User
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)
	
	user.UserRoutes(r, userController)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}