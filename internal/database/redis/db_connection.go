package redis

import (
	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	return client, nil
}