package routes

import (
	merchantService "e-menu-tentakel/core/service/merchant"
	merchantRepo "e-menu-tentakel/infrastructure/repository/merchant"
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

	merchantRepo := merchantRepo.NewMerchantRepository(config.RedisClient)
	merchantService := merchantService.NewMerchantService(merchantRepo)
	weblinkMiddleware := middleware.NewWebLinkMiddleware(merchantService)

	transaction := e.Group("/transaction")
	transaction.Use(weblinkMiddleware.Weblink())
	transaction.POST("/checkout/:outlet_code", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"outlet_id":   c.Get("outlet_id").(string),
			"merchant_id": c.Get("merchant_id").(string),
		})
	})
}
