package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Authentication middleware")
	}
}
