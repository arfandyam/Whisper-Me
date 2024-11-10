package auth

import (
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine, authController controllers.AuthControllerInterface){
    auth := route.Group("/auth")
    {
        auth.POST("", authController.LoginUser)
        auth.PUT("", authController.UpdateAccessToken)
    }
}