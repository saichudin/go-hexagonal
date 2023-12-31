package routes

import (
	merchantCore "go-hexagonal/core/service/merchant"
	weborderAdapter "go-hexagonal/infrastructure/adapter/weborder"
	merchantRepo "go-hexagonal/infrastructure/repository/merchant/redis"
	merchantHandler "go-hexagonal/interface/api/extl/v1/merchant"
	risetHandler "go-hexagonal/interface/api/extl/v1/riset"
	risetService "go-hexagonal/core/service/riset"
	risetAdapter "go-hexagonal/infrastructure/adapter/riset"
	"go-hexagonal/interface/api/extl/v1/routes/middleware"

	"go-hexagonal/utils/config"
	"go-hexagonal/utils/logger"
)

func MerchantInjector() (handler merchantHandler.MerchantHandlerContract, weblinkMiddleware *middleware.WebLinkMiddleware) {
	weborderAdapter := weborderAdapter.NewWeborderAdapter()
	merchantRepo := merchantRepo.NewMerchantRepository(config.RedisClient)
	merchantCore := merchantCore.NewMerchantService(merchantRepo, weborderAdapter, logger.Logger)
	handler = merchantHandler.NewMerchantHandler(merchantCore)
	weblinkMiddleware = middleware.NewWebLinkMiddleware(merchantCore)
	return
}

func RisetInjector() (handler risetHandler.RisetHandlerContract) {
	risetAdapter := risetAdapter.NewRisetAdapter()
	risetService := risetService.NewRisetService(risetAdapter)
	handler = risetHandler.NewRisetHandler(risetService)

	return
}