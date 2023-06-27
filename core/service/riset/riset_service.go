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

func (srv RisetService) GetMpayCustomer(id int) (*mpay.ResponseMpayCustomer, error) {
	var response mpay.ResponseMpayCustomer
	mpayCustomer, err := srv.adapterRiset.GetMpayCustomer(id)
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

func (srv RisetService) CreateMpayCustomer(payload request.MpayCustPayload) (*mpay.ResponseMpayCustomer, error) {
	var response mpay.ResponseMpayCustomer
	err := srv.adapterRiset.CreateMpayCustomer(payload)

	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		return &response, err
	}
	mpayCustomer, err := srv.adapterRiset.GetMpayCustomerByPhone(payload.PhoneNumber)

	response.Message = "success create data"
	response.Data = *mpayCustomer

	return &response, err
}

func (srv RisetService) UpdateMpayCustomer(id int, payload request.MpayCustPayload) (*mpay.ResponseMpayCustomer, error) {
	var response mpay.ResponseMpayCustomer
	err := srv.adapterRiset.UpdateMpayCustomer(id, payload)

	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		return &response, err
	}
	mpayCustomer, err := srv.adapterRiset.GetMpayCustomer(id)

	response.Message = "success update data"
	response.Data = *mpayCustomer

	return &response, err
}

func (srv RisetService) DeleteMpayCustomer(id int) (*mpay.ResponseMpayCustomer, error) {
	var response mpay.ResponseMpayCustomer
	err := srv.adapterRiset.DeleteMpayCustomer(id)
	if err != nil {
		response.Message = err.Error()
		response.Data = nil

		return &response, err
	}

	response.Message = "success delete data"

	return &response, err
}
