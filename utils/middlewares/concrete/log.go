package concrete

import (
	"go-hexagonal/utils/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LogMiddleware struct {
	MiddlewareProto
}

func (hv *LogMiddleware) GetCallback() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: log.APILogHandler,
		Skipper: log.APILogSkipper,
	})
}

func (hv *LogMiddleware) GetName() string {
	return "engine.middleware.core.LogMiddleware"
}

func (hv *LogMiddleware) IsPre() bool {
	return false
}
