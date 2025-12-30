package tools

import (
	"context"
	"quick-start/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (rc *redis.Client, err error) {
	cfg := config.Cfg.Redis
	rc = redis.NewClient(&redis.Options{
		Addr:       cfg.Addr,
		ClientName: cfg.ClientName,
		Username:   cfg.Username,
		Password:   cfg.Password,
		DB:         cfg.DB,
	})

	if err := rc.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rc, nil
}
