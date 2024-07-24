package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			c.Header("Content-Type", "application/json")
			// Log or handle the errors here as needed
			// For example, return a JSON response with the error
			c.JSON(
				// TODO: update the status code of error
				http.StatusInternalServerError,
				gin.H{
					"status":  http.StatusInternalServerError,
					"message": c.Errors.Last().Error(),
					"notes":   "currently all errors are treated as internal server error",
				},
			)

			// Prevent calling any subsequent handlers
			c.Abort()
		}
	}
}

func corsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement this if necessary
	}
}

func headerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement this, like timezone, locale etc
	}
}
