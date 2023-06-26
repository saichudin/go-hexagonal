package mpay

import (
	"go-hexagonal/utils/paginator"
)

type MpayCustomer struct {
	Id                 int    `json:"id"`
	Customer_id        int    `json:"customer_id"`
	Phone_number       string `json:"phone_number"`
	Mpay_customer_code string `json:"mpay_customer_code"`
}

type ResponseMpayCustomer struct {
	Message string       `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseMpayCustomers struct {
	Message string          `json:"message"`
	Data    *[]MpayCustomer `json:"data"`
	Meta 	paginator.Meta	`json:"meta"`
}

type MpayCustomerRequest struct {
	Phone_number string `json:"phone_number"`
}