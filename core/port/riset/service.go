package riset

import (
	"go-hexagonal/core/model/mpay"
	"go-hexagonal/interface/api/extl/v1/riset/request"
)

type RisetService interface {
	GetMpayCustomers(params request.MpayCustReq) (*mpay.ResponseMpayCustomers, error)
	GetMpayCustomer(phone string) (*mpay.ResponseMpayCustomer, error)
}
