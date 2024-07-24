package ginx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// wrapGinFunc, wraps the http.HandlerFunc into gin.HandlerFunc
// And also, sets the path parameters in the request context
func wrapGinFunc(hs []http.HandlerFunc) []gin.HandlerFunc {
	gHFs := make([]gin.HandlerFunc, 0, len(hs))
	for _, hanFunc := range hs {
		gHFs = append(
			gHFs,
			func(c *gin.Context) {
				ctx := c.Request.Context()
				for _, param := range c.Params {
					ctx = context.WithValue(ctx, param.Key, param.Value)
				}
				c.Request = c.Request.WithContext(ctx)
				hanFunc(c.Writer, c.Request)
			},
		)
	}

	return gHFs
}
