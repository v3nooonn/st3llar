package main

import (
	"net/http"

	"github.com/v3nooom/st3llar/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(ErrorMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/demo", handler.DemoHandler())

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
