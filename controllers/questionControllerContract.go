package controllers

import "github.com/gin-gonic/gin"

type QuestionControllerInterface interface {
	CreateQuestion(ctx *gin.Context)
	EditQuestion(ctx *gin.Context)
	FindQuestionById(ctx *gin.Context)
	FindQuestionsByUserId(ctx *gin.Context)
}