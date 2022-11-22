package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check v1 routes!")
	})
}
