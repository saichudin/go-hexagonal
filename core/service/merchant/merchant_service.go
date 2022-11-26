package merchant

import (
	"e-menu-tentakel/core/model"
	portMerchant "e-menu-tentakel/core/port/merchant"
)

type MerchantService struct {
	repoMerchant portMerchant.MerchantRepository
}

func NewMerchantService(repoMerchant portMerchant.MerchantRepository) portMerchant.MerchantService {
	return &MerchantService{
		repoMerchant: repoMerchant,
	}
}

func (srv *MerchantService) OutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error) {
	outletWebLink, err := srv.repoMerchant.GetOutletWebLinkInfo(outletCode)
	if err != nil {
		return nil, err
	}

	return outletWebLink, nil
}

func (srv *MerchantService) DetailOutlet(outletId string) (*model.DetailOutlet, error) {
	detailOutlet, err := srv.repoMerchant.GetDetailOutlet(outletId)
	if err != nil {
		return nil, err
	}

	return detailOutlet, nil
}
