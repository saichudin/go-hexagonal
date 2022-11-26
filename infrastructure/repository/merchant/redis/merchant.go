package redis

import (
	"context"
	"e-menu-tentakel/core/model"
	portMerchant "e-menu-tentakel/core/port/merchant"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type MerchantRepository struct {
	Redis *redis.Client
}

func NewMerchantRepository(redis *redis.Client) portMerchant.MerchantRepository {
	return &MerchantRepository{
		Redis: redis,
	}
}

func (repo *MerchantRepository) GetOutletWebLinkInfo(outletCode string) (weblink *model.WebLinkUri, err error) {
	redisWebLink, err := repo.Redis.Get(context.TODO(), "web_link_uri:"+outletCode).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(redisWebLink), &weblink)
	if err != nil {
		return nil, err
	}

	return weblink, nil
}

func (repo *MerchantRepository) GetDetailOutlet(outletId string) (detailOutlet *model.DetailOutlet, err error) {
	redisDetailOutlet, err := repo.Redis.Get(context.TODO(), "detail_outlet:"+outletId).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(redisDetailOutlet), &detailOutlet)
	if err != nil {
		return nil, err
	}

	return detailOutlet, nil
}
