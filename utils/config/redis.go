package config

import (
	"context"
	"e-menu-tentakel/utils/conv"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

type RedisConfig struct {
	Host     string
	Port     int
	Auth     string
	Database int
}

func (redisConfig *RedisConfig) SetConfigRedis() *RedisConfig {
	redisConfig.Host = os.Getenv("REDISV6_HOST")
	redisConfig.Port = conv.StringToInt(os.Getenv("REDISV6_PORT"), 0)
	redisConfig.Auth = os.Getenv("REDISV6_AUTH")
	redisConfig.Database = conv.StringToInt(os.Getenv("REDISV6_DB"), 0)

	return redisConfig
}

func (redisConfig *RedisConfig) ConnectRedis() (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Auth,
		DB:       redisConfig.Database,
	}

	RedisClient = redis.NewClient(opt)

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return RedisClient, nil
}
