package rdb

import (
	"bluebell/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	timeout       = time.Second * 3
	articleExpire = time.Hour * 24 * 7
)
const (
	prefix          = "bluebell:"
	articleCreateAt = "article:time"
	articleScore    = "article:score"
	articleVoter    = "article:voter:"
)

var db *redis.Client

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

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
	db = redis.NewClient(&op)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err = db.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("connect redis failed: %w", err)
	}

	return nil
}

func formatKey(key string, val ...string) string {
	res := prefix + key
	if len(val) > 0 {
		res += val[0]
	}
	return res
}
