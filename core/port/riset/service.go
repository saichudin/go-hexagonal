package riset

import (
	"go-hexagonal/core/model/mpay"
	"go-hexagonal/interface/api/extl/v1/riset/request"
)

type RisetService interface {
	GetMpayCustomers(params request.MpayCustReq) (*mpay.ResponseMpayCustomers, error)
	GetMpayCustomer(id int) (*mpay.ResponseMpayCustomer, error)
	CreateMpayCustomer(payload request.MpayCustPayload) (*mpay.ResponseMpayCustomer, error)
	UpdateMpayCustomer(id int, payload request.MpayCustPayload) (*mpay.ResponseMpayCustomer, error)
	DeleteMpayCustomer(id int) (*mpay.ResponseMpayCustomer, error)
}
