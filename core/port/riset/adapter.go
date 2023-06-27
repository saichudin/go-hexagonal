package riset

import (
	"go-hexagonal/core/model/mpay"
	"go-hexagonal/interface/api/extl/v1/riset/request"
)

type RisetAdapter interface {
	GetMpayCustomers(params request.MpayCustReq) (*[]mpay.MpayCustomer, int64, error)
	GetMpayCustomer(id int) (*mpay.MpayCustomer, error)
	GetMpayCustomerByPhone(phone string) (*mpay.MpayCustomer, error)
	CreateMpayCustomer(payload request.MpayCustPayload) (error)
	UpdateMpayCustomer(id int, payload request.MpayCustPayload) error
	DeleteMpayCustomer(id int) error
}
