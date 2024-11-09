package controllers

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	LoginUser(ctx *gin.Context)
}