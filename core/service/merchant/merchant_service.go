package merchant

import (
	"e-menu-tentakel/core/model"
	portMerchant "e-menu-tentakel/core/port/merchant"
)

type MerchantService struct {
	repoMerchant    portMerchant.MerchantRepository
	adapterWeborder portMerchant.WeborderAdapter
}

func NewMerchantService(repoMerchant portMerchant.MerchantRepository, adapterWeborder portMerchant.WeborderAdapter) portMerchant.MerchantService {
	return &MerchantService{
		repoMerchant:    repoMerchant,
		adapterWeborder: adapterWeborder,
	}
}

func (srv *MerchantService) GetOutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error) {
	outletWebLink, err := srv.repoMerchant.GetOutletWebLinkInfo(outletCode)
	if err != nil && err.Error() != "redis: nil" {
		return nil, err
	}

	if outletWebLink == nil {
		outletWebLink, err = srv.adapterWeborder.GetDetailWeblink(outletCode)
		if err != nil {
			return nil, err
		}

		if outletWebLink != nil {
			err = srv.repoMerchant.StoreWeblink(outletCode, outletWebLink)
			if err != nil {
				//log aja
			}
		}
	}

	return outletWebLink, nil
}
