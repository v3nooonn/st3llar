package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/v3nooom/st3llar/supplier/internal/api/middleware"
	"net/http"

	"github.com/v3nooom/st3llar/supplier/internal/api/handler"
)

func (g *GinAdaptor) RegisterRoutes() {
	g.GET("/up", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})

	oauthGroup := g.Engine.Group("/oauth")
	{
		oauthGroup.POST("/login/:org", func(c *gin.Context) {
			handler.Login(c.Writer, c.Request)
		})

		oauthGroup.POST("/logout", wrapGinFunc([]http.HandlerFunc{handler.Login})...)
	}

	lambdaGroup := g.Engine.Group("/lambda", middleware.Authentication(), middleware.Authorization())
	{
		lambdaGroup.POST("/register/:id", wrapGinFunc([]http.HandlerFunc{handler.Login})...)
	}

	g.POST("/oauth/:id", handler.Login)
}
