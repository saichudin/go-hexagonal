package config

import (
	"go-hexagonal/utils/middlewares"
	"go-hexagonal/utils/middlewares/concrete"
)

var (
	MiddlewareFactory = middlewares.NewMiddlewareFactory()
)

func RegisterRequiredMiddleware() {
	MiddlewareFactory.Register(&concrete.RemoveTrailingSlash{}, true)
	MiddlewareFactory.Register(&concrete.LogMiddleware{}, true)
}
