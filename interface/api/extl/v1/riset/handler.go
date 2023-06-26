package riset

import (
	"fmt"
	"go-hexagonal/core/model/mpay"
	port "go-hexagonal/core/port/riset"
	"net/http"
	"strconv"

	"go-hexagonal/interface/api/extl/v1/riset/request"

	"github.com/labstack/echo/v4"
)

type RisetHandlerContract interface {
	GetMpayCustomers(c echo.Context) error
	GetMpayCustomer(c echo.Context) error
}

type RisetHandler struct {
	service port.RisetService
}

func NewRisetHandler(svc port.RisetService) RisetHandlerContract {
	return &RisetHandler{
		service: svc,
	}
}

func (h RisetHandler) GetMpayCustomers(c echo.Context) error {
	var page int
	var limit int
	page , _ = strconv.Atoi(c.QueryParam("page")) 
	limit , _ = strconv.Atoi(c.QueryParam("limit"))
	if page == 0 {page = 1} 
	if limit == 0 {limit = 10}

	params := new(request.MpayCustReq)
	params.Page = page
	params.Limit = limit

	res, err := h.service.GetMpayCustomers(*params)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	// return &response, nil
	return c.JSON(http.StatusOK, res)
}

func (h RisetHandler) GetMpayCustomer(c echo.Context) error {
	var payload mpay.MpayCustomerRequest

	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	fmt.Println(payload)
	res, err := h.service.GetMpayCustomer(payload.Phone_number)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	// return &response, nil
	return c.JSON(http.StatusOK, res)
}
