package riset

import (
	"go-hexagonal/core/model/mpay"
	"go-hexagonal/interface/api/extl/v1/riset/request"
)

type RisetAdapter interface {
	GetMpayCustomers(params request.MpayCustReq) (*[]mpay.MpayCustomer, int64, error)
	GetMpayCustomer(phone string) (*mpay.MpayCustomer, error)
}
