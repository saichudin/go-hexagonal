package riset

import (
	"fmt"
	"go-hexagonal/core/model/mpay"
	port "go-hexagonal/core/port/riset"
	"go-hexagonal/interface/api/extl/v1/riset/request"
	"go-hexagonal/utils/paginator"
)

type RisetService struct {
	adapterRiset port.RisetAdapter
}

func NewRisetService(adapterRiset port.RisetAdapter) port.RisetService {
	return &RisetService{
		adapterRiset: adapterRiset,
	}
}

func (srv RisetService) GetMpayCustomers(params request.MpayCustReq) (*mpay.ResponseMpayCustomers, error) {
	mpayCustomers, count, err := srv.adapterRiset.GetMpayCustomers(params)
	if err != nil {

	}

	var response mpay.ResponseMpayCustomers
	response.Message = "success"
	response.Data = mpayCustomers
	response.Meta = paginator.Paginator(params.Page, params.Limit, count)

	return &response, err
}

func (srv RisetService) GetMpayCustomer(phone string) (*mpay.ResponseMpayCustomer, error) {
	var response mpay.ResponseMpayCustomer
	mpayCustomer, err := srv.adapterRiset.GetMpayCustomer(phone)
	fmt.Println()
	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		return &response, err
	}

	response.Message = "success"
	response.Data = *mpayCustomer

	return &response, err
}
