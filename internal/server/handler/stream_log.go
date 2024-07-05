package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StreamLogHandler() gin.HandlerFunc {
	type logStream struct {
		Log string `json:"log"`
	}

	var stream logStream

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWithJSON(&stream); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		fmt.Println("Received log from AppRunner: ", stream.Log)
	}
}
