package main

import (
	"context"
	"go-hexagonal/utils/config"
	"go-hexagonal/utils/conv"
	"go-hexagonal/utils/logger"
	"go-hexagonal/utils/validation"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"time"

	routeLite "go-hexagonal/interface/api/extl/lite/routes"
	routeV1 "go-hexagonal/interface/api/extl/v1/routes"
	routeWeborder "go-hexagonal/interface/api/extl/weborder/v1/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.SetConfig()
	config.RegisterRequiredMiddleware()
	logger.LoadLogger()
	AppStart()
}

func AppStart() {
	e := echo.New()

	e.Use(middleware.Recover())

	// toggle debug mode
	if conv.StringToBool(os.Getenv("APP_DEBUG")) {
		e.Debug = true
	}
	DispatchMiddleware(e)

	// register validator
	val := validator.New()
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	e.Validator = &validation.Validator{Validator: val}

	// assign route
	routeLite.API(e)
	routeWeborder.API(e)
	routeV1.API(e)

	go func() {
		if err := e.Start(":" + os.Getenv("APP_PORT")); err != nil {
			logger.Logger.Info("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Logger.Error(err.Error())
	}
}

/**
 * Dispatch registered middleware.
 *
 * @param *echo.Echo Echo framework instance (as reference.)
 * @return void
 */
func DispatchMiddleware(e *echo.Echo) {
	middlewares := config.MiddlewareFactory.GetAll()

	for _, mf := range middlewares {
		if mf.IsPre() {
			e.Pre(mf.GetCallback())
		} else {
			e.Use(mf.GetCallback())
		}
	}
}
