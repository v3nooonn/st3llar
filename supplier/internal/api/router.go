package api

import "net/http"

type Router interface {
	RegisterRoutes()
	Run(port string)
	GET(path string, handlers ...http.HandlerFunc)
	POST(path string, handlers ...http.HandlerFunc)
	PUT(path string, handlers ...http.HandlerFunc)
	PATCH(path string, handlers ...http.HandlerFunc)
	DELETE(path string, handlers ...http.HandlerFunc)
	OPTIONS(path string, handlers ...http.HandlerFunc)
}
