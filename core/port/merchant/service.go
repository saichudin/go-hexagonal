package merchant

import "e-menu-tentakel/core/model"

type MerchantService interface {
	OutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error)
	DetailOutlet(outletId string) (*model.DetailOutlet, error)
}
