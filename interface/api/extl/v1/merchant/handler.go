package merchant

import (
	"net/http"

	"github.com/labstack/echo/v4"

	port "go-hexagonal/core/port/merchant"
	"go-hexagonal/utils/constants"
	"go-hexagonal/utils/response"
)

type MerchantHandlerContract interface {
	GetOutletWebLinkInfo(c echo.Context) error
}

type MerchantHandler struct {
	service port.MerchantService
}

func NewMerchantHandler(svc port.MerchantService) MerchantHandlerContract {
	return &MerchantHandler{service: svc}
}

func (h *MerchantHandler) GetOutletWebLinkInfo(c echo.Context) error {
	resp := new(response.Response)

	outletWeblink, err := h.service.GetOutletWebLinkInfo(c.Param("weblink_url"))
	if err != nil {
		resp.SetResponse(err.Error(), nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if outletWeblink == nil {
		resp.SetResponse(constants.NotFoundResponse, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp.SetResponse(constants.SuccessResponse, outletWeblink)
	return c.JSON(http.StatusBadRequest, resp)
}
