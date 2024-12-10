package controllers

import "github.com/gin-gonic/gin"

type ResponseControllerInterface interface {
	CreateResponse(ctx *gin.Context)
	FindResponseByQuestionId(ctx *gin.Context)
}
