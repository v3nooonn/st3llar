package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("---> Authorization middleware")
	}
}
