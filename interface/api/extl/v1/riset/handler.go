package riset

import (
	port "go-hexagonal/core/port/riset"
	"net/http"
	"strconv"
	"time"

	"go-hexagonal/interface/api/extl/v1/riset/request"

	"github.com/labstack/echo/v4"
)

type RisetHandlerContract interface {
	GetMpayCustomers(c echo.Context) error
	GetMpayCustomer(c echo.Context) error
	CreateMpayCustomer(c echo.Context) error
	UpdateMpayCustomer(c echo.Context) error
	DeleteMpayCustomer(c echo.Context) error
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
	id , _ := strconv.Atoi(c.Param("id")) 

	// fmt.Println(payload)
	res, err := h.service.GetMpayCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}

	// return &response, nil
	return c.JSON(http.StatusOK, res)
}

func (h RisetHandler) CreateMpayCustomer(c echo.Context) error {
	var payload request.MpayCustPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	// payload.ID = 0
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	res, err := h.service.CreateMpayCustomer(payload)
	return c.JSON(http.StatusOK, res)
}

func (h RisetHandler) UpdateMpayCustomer(c echo.Context) error {
	var payload request.MpayCustPayload
	id , _ := strconv.Atoi(c.Param("id")) 
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	res, _ := h.service.UpdateMpayCustomer(id, payload)

	return c.JSON(http.StatusOK, res)
}

func (h RisetHandler) DeleteMpayCustomer(c echo.Context) error {
	id , _ := strconv.Atoi(c.Param("id")) 
	res, _ := h.service.DeleteMpayCustomer(id)

	return c.JSON(http.StatusOK, res)
}