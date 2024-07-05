package main

import (
	"net/http"

	"github.com/v3nooom/st3llar-helper/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(ErrorMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/lambda/:input", handler.MessageHandler())
	r.POST("/hooks/stream/log", handler.StreamLogHandler())

	r.Run(":8080")
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.Header("Content-Type", "application/json")

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": c.Errors.Last().Error(),
			})

			c.Abort()
			return
		}
	}
}
