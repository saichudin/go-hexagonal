package routes

import (
	merchantService "e-menu-tentakel/core/service/merchant"
	merchantRedis "e-menu-tentakel/infrastructure/repository/merchant/redis"
	"e-menu-tentakel/interface/api/extl/lite/routes/middleware"
	"e-menu-tentakel/utils/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	lite := e.Group("/lite")
	lite.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check lite routes!")
	})

	merchantRedis := merchantRedis.NewMerchantRepository(config.RedisClient)
	merchantService := merchantService.NewMerchantService(merchantRedis)
	weblinkMiddleware := middleware.NewWebLinkMiddleware(merchantService)

	transaction := e.Group("/transaction")
	transaction.Use(weblinkMiddleware.Weblink())
	transaction.POST("/checkout/:outlet_code", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"outlet_id":        c.Get("outlet_id"),
			"merchant_id":      c.Get("merchant_id"),
			"outlet_longitude": c.Get("outlet_longitude"),
			"outlet_latitude":  c.Get("outlet_latitude"),
			"outlet_couriers":  c.Get("outlet_couriers"),
		})
	})
}
