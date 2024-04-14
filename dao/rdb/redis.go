package rdb

import (
	"bluebell/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const RedisTimeOut = time.Second * 3

var rdb *redis.Client

func Init() error {
	redisCfg, err := config.Cfg.Redis()
	if err != nil {
		return fmt.Errorf("read redis config failed: %w", err)
	}

	op := redis.Options{
		Username: redisCfg.User,
		Password: redisCfg.Password,
		Addr:     redisCfg.Addr,
	}
	c := redis.NewClient(&op)

	ctx, cancel := context.WithTimeout(context.Background(), RedisTimeOut)
	defer cancel()
	if err = c.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("connect redis failed: %w", err)
	}

	rdb = c
	return nil
}
