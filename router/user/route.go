package user

import (
	"github.com/gin-gonic/gin"
	"github.com/arfandyam/Whisper-Me/controllers"
)

func UserRoutes(route *gin.Engine, userController controllers.UserControllerInterface){
	user := route.Group("/user")
	{
		// Create User
		user.POST("", userController.CreateUser)
		user.PUT("/:id", userController.EditUser)
	}
}