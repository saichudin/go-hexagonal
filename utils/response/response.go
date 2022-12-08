package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (r *Response) SetResponse(message string, data any) {
	r.Message = message
	r.Data = data
}

type ResponseMsg struct {
	Status  bool   `json:"status"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func (r *ResponseMsg) SetResponseMsg(status bool, message string, data any) {
	r.Status = status
	r.Message = message
	r.Data = data
}

type ResponseWithMeta struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    Meta   `json:"meta"`
}

type Meta struct {
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page"`
	PrevPage    int   `json:"prev_page"`
	NextPage    int   `json:"next_page"`
	TotalPage   int   `json:"total_page"`
}

func NewResponse(ctx echo.Context, statusCode int, message string, data interface{}) error {
	res := Response{
		Message: message,
		Data:    data,
	}

	return ctx.JSON(statusCode, res)
}

func NewResponseWithMeta(ctx echo.Context, statusCode int, message string, data interface{}, meta map[string]interface{}) error {
	res := ResponseWithMeta{
		Message: message,
		Data:    data,
		Meta:    setMeta(meta),
	}

	return ctx.JSON(statusCode, res)
}

func setMeta(meta map[string]interface{}) Meta {
	var currentPage, totalPage, prevPage, nextPage int
	currentPage = meta["currentPage"].(int)
	totalPage = meta["totalPage"].(int)

	if currentPage == totalPage {
		nextPage = currentPage
	} else {
		nextPage = currentPage + 1
	}

	if currentPage == 1 {
		prevPage = currentPage
	} else {
		prevPage = currentPage - 1
	}

	return Meta{
		Limit:       meta["limit"].(int),
		Total:       meta["total"].(int64),
		CurrentPage: meta["currentPage"].(int),
		PrevPage:    prevPage,
		NextPage:    nextPage,
		TotalPage:   totalPage,
	}
}
