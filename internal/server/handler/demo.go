package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DemoHandler() gin.HandlerFunc {
	type Payload struct {
		Field1 string `json:"field1"`
	}

	var payload Payload

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fmt.Println("Received log from AppRunner: ", payload.Field1)
	}
}
