package merchant

import "e-menu-tentakel/core/model"

type MerchantRepository interface {
	GetOutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error)
	StoreWeblink(key string, value interface{}) error
	DeleteWeblink(key string) error
}
