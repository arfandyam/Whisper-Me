package controllers

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	CreateUser(ctx *gin.Context)
	EditUser(ctx *gin.Context)
	FindUserById(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}