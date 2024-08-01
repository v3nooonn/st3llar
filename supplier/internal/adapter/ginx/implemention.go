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

func (g *GinAdaptor) Run(port string) {
	g.Engine.Run(port)
}

func (g *GinAdaptor) GET(path string, handlers ...http.HandlerFunc) {
	g.Engine.GET(path, wrapGinFunc(handlers)...)
}

func (g *GinAdaptor) POST(path string, handlers ...http.HandlerFunc) {
	g.Engine.POST(path, wrapGinFunc(handlers)...)
}

func (g *GinAdaptor) PUT(path string, handlers ...http.HandlerFunc) {
	g.Engine.PUT(path, wrapGinFunc(handlers)...)
}

func (g *GinAdaptor) PATCH(path string, handlers ...http.HandlerFunc) {
	g.Engine.PATCH(path, wrapGinFunc(handlers)...)
}

func (g *GinAdaptor) DELETE(path string, handlers ...http.HandlerFunc) {
	g.Engine.DELETE(path, wrapGinFunc(handlers)...)
}

func (g *GinAdaptor) OPTIONS(path string, handlers ...http.HandlerFunc) {
	g.Engine.OPTIONS(path, wrapGinFunc(handlers)...)
}
