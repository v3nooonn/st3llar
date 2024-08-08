package ginx

import (
	"net/http"

	"github.com/v3nooom/st3llar/supplier/internal/api"

	"github.com/gin-gonic/gin"
)

type (
	GinAdaptor struct {
		Engine *gin.Engine
	}
	GinAdaptorOpts func(*GinAdaptor)
)

func NewWithOpts(opts ...GinAdaptorOpts) api.Router {
	sv := &GinAdaptor{Engine: gin.Default()}

	for _, opt := range opts {
		opt(sv)
	}

	return sv
}

func WithCustomRecovery() GinAdaptorOpts {
	return func(adaptor *GinAdaptor) {
		adaptor.Engine.Use(
			gin.CustomRecovery(func(c *gin.Context, err interface{}) {
				e, ok := err.(error)
				if ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": false,
						"message": "internal error",
						"msg":     e.Error(),
					})
				}
			},
			),
		)
	}
}

func WithCORS() GinAdaptorOpts {
	return func(adaptor *GinAdaptor) {
		adaptor.Engine.Use(corsHandler())
	}
}

func WithHeaderHandler() GinAdaptorOpts {
	return func(adaptor *GinAdaptor) {
		adaptor.Engine.Use(headerHandler())
	}
}

func WithErrorHandler() GinAdaptorOpts {
	return func(adaptor *GinAdaptor) {
		adaptor.Engine.Use(errorHandler())
	}
}
