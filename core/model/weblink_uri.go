package model

type WebLinkUri struct {
	No                      string        `json:"no"`
	IdOutlet                int           `json:"outlet_id"`
	IdMerchant              int           `json:"merchant_id"`
	IdUser                  string        `json:"user_id"`
	WebLinkUrl              string        `json:"weblink_url"`
	WebLinkQr               string        `json:"weblink_qr"`
	EateryType              string        `json:"eatery_type"`
	EaterySubType           string        `json:"eatery_sub_type"`
	ListCourier             []ListCourier `json:"shipper_agent_setup"`
	ShipperAddress          string        `json:"address"`
	ShipperAddressBenchmark string        `json:"address_benchmark"`
	ShipperLongitude        string        `json:"longitude"`
	ShipperLatitude         string        `json:"latitude"`
	ShipperPostalCode       string        `json:"postal_code"`
	IsPrimary               int           `json:"is_primary"`
	IsLite                  int           `json:"is_lite"`
	ProductView             string        `json:"product_view"`
}

type ListCourier struct {
	Code            string   `json:"code"`
	ServiceTypeCode []string `json:"service_types_code"`
}
