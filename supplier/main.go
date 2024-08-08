package main

import (
	"github.com/v3nooom/st3llar/supplier/internal/adapter/ginx"
)

func main() {
	router := ginx.NewWithOpts(
		ginx.WithCustomRecovery(),
		ginx.WithCORS(),
		ginx.WithHeaderHandler(),
		ginx.WithErrorHandler(),
	)

	router.RegisterRoutes()

	router.Run(":8080")
}
