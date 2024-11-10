package controllers

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	LoginUser(ctx *gin.Context)
	UpdateAccessToken(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
}