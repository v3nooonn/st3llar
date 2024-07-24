package ginx

import (
	"net/http"
)

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
