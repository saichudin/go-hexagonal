package redis

import (
	"context"
	"encoding/json"
	"go-hexagonal/core/model"
	portMerchant "go-hexagonal/core/port/merchant"
	"go-hexagonal/utils/constants"

	"github.com/go-redis/redis/v8"
)

type MerchantRepository struct {
	Redis *redis.Client
}

var ctx = context.Background()

func NewMerchantRepository(redis *redis.Client) portMerchant.MerchantRepository {
	return &MerchantRepository{
		Redis: redis,
	}
}

func (repo *MerchantRepository) GetOutletWebLinkInfo(outletCode string) (weblink *model.WebLinkUri, err error) {
	redisWebLink, err := repo.Redis.Get(ctx, "web_link_uri:"+outletCode).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(redisWebLink), &weblink)
	if err != nil {
		return nil, err
	}

	return weblink, nil
}

func (repo *MerchantRepository) StoreWeblink(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = repo.Redis.Set(ctx, "web_link_uri:"+key, data, constants.RedisExpireDuration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (repo *MerchantRepository) DeleteWeblink(key string) error {
	return repo.Redis.Del(ctx, "web_link_uri:"+key).Err()
}
