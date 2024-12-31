package main

import (
	"github.com/arfandyam/Whisper-Me/libs/exceptions"
	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *exceptions.CustomError:
				ctx.AbortWithStatusJSON(e.Status, gin.H{
					"status":      "failed",
					"description": e.Description,
					"message":     e.Message,
				})
			default:
				ctx.JSON(500, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
			}
		}

		ctx.Abort()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
