package merchant

import (
	"e-menu-tentakel/core/model"
	portMerchant "e-menu-tentakel/core/port/merchant"

	"bitbucket.org/klopos/majoo-logger/log"
	"bitbucket.org/klopos/majoo-logger/logger"
)

type MerchantService struct {
	repoMerchant    portMerchant.MerchantRepository
	adapterWeborder portMerchant.WeborderAdapter
	log             logger.ILogger
}

func NewMerchantService(repoMerchant portMerchant.MerchantRepository, adapterWeborder portMerchant.WeborderAdapter, log logger.ILogger) portMerchant.MerchantService {
	return &MerchantService{
		repoMerchant:    repoMerchant,
		adapterWeborder: adapterWeborder,
		log:             log,
	}
}

func (srv *MerchantService) GetOutletWebLinkInfo(outletCode string) (*model.WebLinkUri, error) {
	srv.log.Info("GetOutletWebLinkInfo Logic")
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
				srv.log.Error(err.Error(), log.WithData(map[string]interface{}{
					"key":     outletCode,
					"payload": outletWebLink,
				}))
			}
		}
	}

	return outletWebLink, nil
}
