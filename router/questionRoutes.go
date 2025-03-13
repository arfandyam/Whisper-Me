package router

import (
	"github.com/arfandyam/Whisper-Me/controllers"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(route *gin.Engine, questionController controllers.QuestionControllerInterface) {
	question := route.Group("/question")
	{
		question.POST("", questionController.CreateQuestion)
		question.PUT(":id", questionController.EditQuestion)
		// question.GET(":id", questionController.FindQuestionById)
		question.GET(":slug", questionController.FindQuestionBySlug)
		question.GET("", questionController.FindQuestionsByUserId)
		question.GET("/r/:urlKey", questionController.ShortenUrl)
		question.GET("/search", questionController.SearchQuestionsByKeyword)
	}
}
