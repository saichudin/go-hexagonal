package request

import "time"

type MpayCustReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type MpayCustPayload struct {
	// ID               int       `json:"id"`
	CustomerId       int       `json:"customer_id"`
	PhoneNumber      string    `json:"phone_number"`
	MpayCustomerCode string    `json:"mpay_customer_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type MpayCustUpdatePayload struct {
	CustomerId       int       `json:"customer_id"`
	PhoneNumber      string    `json:"phone_number"`
	MpayCustomerCode string    `json:"mpay_customer_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
