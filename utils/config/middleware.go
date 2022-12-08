package config

import (
	"e-menu-tentakel/utils/middlewares"
	"e-menu-tentakel/utils/middlewares/concrete"
)

var (
	MiddlewareFactory = middlewares.NewMiddlewareFactory()
)

func RegisterRequiredMiddleware() {
	MiddlewareFactory.Register(&concrete.RemoveTrailingSlash{}, true)
	MiddlewareFactory.Register(&concrete.LogMiddleware{}, true)
}
