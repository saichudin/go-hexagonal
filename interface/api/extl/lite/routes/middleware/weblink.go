package middleware

import (
	port "e-menu-tentakel/core/port/merchant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebLinkMiddleware struct {
	service port.MerchantService
}

func NewWebLinkMiddleware(service port.MerchantService) *WebLinkMiddleware {
	return &WebLinkMiddleware{
		service: service,
	}
}

func (mdw WebLinkMiddleware) Weblink() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			weblinkUri, err := mdw.service.OutletWebLinkInfo(c.Param("outlet_code"))
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "failed parsing outlet code")
			}

			c.Set("outlet_id", weblinkUri.IdOutlet)
			c.Set("merchant_id", weblinkUri.IdMerchant)

			return next(c)
		}
	}
}
