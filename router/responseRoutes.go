package router

import (
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/gin-gonic/gin"
)

func ResponseRoute(route *gin.Engine, responseController controllers.ResponseControllerInterface){
	response := route.Group("/response")
	{
		response.POST("", responseController.CreateResponse)
	}
}