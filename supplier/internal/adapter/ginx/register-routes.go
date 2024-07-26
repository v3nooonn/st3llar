package ginx

import (
	"net/http"

	"github.com/v3nooom/st3llar/supplier/internal/api/handler"
	"github.com/v3nooom/st3llar/supplier/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func (g *GinAdaptor) RegisterRoutes() {
	g.GET("/up", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})

	oauthGroup := g.Engine.Group("/oauth")
	{
		oauthGroup.POST("/login", func(c *gin.Context) {
			handler.Login(c.Writer, c.Request)
		})

		oauthGroup.POST("/logout", wrapGinFunc([]http.HandlerFunc{handler.Login})...)
		oauthGroup.POST("/refresh", wrapGinFunc([]http.HandlerFunc{handler.Login})...)
	}

	lambdaGroup := g.Engine.Group("/lambda", middleware.Authentication(), middleware.Authorization())
	{
		lambdaGroup.POST("/register/:id", wrapGinFunc([]http.HandlerFunc{handler.Register})...)
	}
}
