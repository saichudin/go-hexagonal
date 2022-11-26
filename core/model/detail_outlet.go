package model

type DetailOutlet struct {
	IdOutlet         string        `json:"id_outlet"`
	IdMerchant       string        `json:"id_merchant"`
	OwnerName        string        `json:"owner_name"`
	ListCourier      []ListCourier `json:"list_courier"`
	ShipperLongitude string        `json:"shipper_longitude"`
	ShipperLatitude  string        `json:"shipper_latitude"`
}

type ListCourier struct {
	Code            string   `json:"code"`
	ServiceTypeCode []string `json:"service_types_code"`
}
