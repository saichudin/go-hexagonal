package merchant

import "e-menu-tentakel/core/model"

type MerchantService interface {
	GetOutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error)
}
