package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check lite routes!")
	})

	merchant := v1.Group("/merchant")
	merchantHandler, merchantMiddleware := MerchantInjector()
	merchant.GET("/:weblink_url", merchantHandler.GetOutletWebLinkInfo)

	test := v1.Group("/test")
	test.Use(merchantMiddleware.Weblink())
	test.GET("/health/:outlet_code", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check lite routes!")
	})

	risetHandler := RisetInjector()
	risetRoute := e.Group("/riset")
	risetRoute.GET("/mpay/customers", risetHandler.GetMpayCustomers)
	risetRoute.GET("/mpay/customer/:id", risetHandler.GetMpayCustomer)
	risetRoute.POST("/mpay/customer", risetHandler.CreateMpayCustomer)
	risetRoute.PUT("/mpay/customer/:id", risetHandler.UpdateMpayCustomer)
	risetRoute.DELETE("/mpay/customer/:id", risetHandler.DeleteMpayCustomer)
}
