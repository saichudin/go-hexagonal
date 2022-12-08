package model

type DetailOutlet struct {
	IdOutlet         string        `json:"id_outlet"`
	IdMerchant       string        `json:"id_merchant"`
	IdUser           string        `json:"id_user"`
	OwnerName        string        `json:"owner_name"`
	ListCourier      []ListCourier `json:"list_courier"`
	ShipperLongitude string        `json:"shipper_longitude"`
	ShipperLatitude  string        `json:"shipper_latitude"`
}
