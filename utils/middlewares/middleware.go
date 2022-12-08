package middlewares

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

// Middleware Struct
type Middleware struct{}

// Log method is used to activate middleware logger request
func (_m Middleware) Log() echo.MiddlewareFunc {
	return mid.Logger()
}
