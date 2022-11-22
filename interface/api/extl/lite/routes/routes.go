package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	lite := e.Group("/lite")
	lite.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check lite routes!")
	})
}
