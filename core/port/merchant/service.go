package merchant

import "go-hexagonal/core/model"

type MerchantService interface {
	GetOutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error)
}
