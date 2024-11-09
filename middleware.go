package main

import (
	"github.com/gin-gonic/gin"
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type){
			case *exceptions.CustomError:
				ctx.AbortWithStatusJSON(e.Status, gin.H{
					"status": "failed",
					"description": e.Description,
					"message": e.Message,
				})
			default:
				ctx.JSON(500, gin.H{
					"status": "failed",
					"message": err.Error(),
				})
			}
		}

		ctx.Abort()
	}
}