package routes

import (
	merchantCore "e-menu-tentakel/core/service/merchant"
	weborderAdapter "e-menu-tentakel/infrastructure/adapter/weborder"
	merchantRepo "e-menu-tentakel/infrastructure/repository/merchant/redis"
	merchantHandler "e-menu-tentakel/interface/api/extl/v1/merchant"
	"e-menu-tentakel/interface/api/extl/v1/routes/middleware"

	"e-menu-tentakel/utils/config"
	"e-menu-tentakel/utils/logger"
)

func MerchantInjector() (handler merchantHandler.MerchantHandlerContract, weblinkMiddleware *middleware.WebLinkMiddleware) {
	weborderAdapter := weborderAdapter.NewWeborderAdapter()
	merchantRepo := merchantRepo.NewMerchantRepository(config.RedisClient)
	merchantCore := merchantCore.NewMerchantService(merchantRepo, weborderAdapter, logger.Logger)
	handler = merchantHandler.NewMerchantHandler(merchantCore)
	weblinkMiddleware = middleware.NewWebLinkMiddleware(merchantCore)
	return
}
