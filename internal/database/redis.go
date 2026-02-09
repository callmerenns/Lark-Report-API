package database

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

var Redis *redis.Client

func InitRedis(addr, password string, db int) {

	logger.Log.App.Info(
		"initializing_redis_connection",
		zap.String("addr", addr),
	)

	var opt *redis.Options
	var err error

	// 🔥 Support redis:// & rediss:// (Upstash)
	if strings.HasPrefix(addr, "redis://") || strings.HasPrefix(addr, "rediss://") {
		opt, err = redis.ParseURL(addr)
		if err != nil {
			logger.Log.Error.Error(
				"failed_to_parse_redis_url",
				zap.Error(err),
			)
			panic(err)
		}
	} else {
		opt = &redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}
	}

	Redis = redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Redis.Ping(ctx).Err(); err != nil {
		logger.Log.Error.Error(
			"failed_to_connect_redis",
			zap.Error(err),
		)
		panic(err)
	}

	logger.Log.App.Info(
		"redis_connected",
		zap.Int("db", db),
	)
}

func CloseRedis() {
	if Redis != nil {
		_ = Redis.Close()
		logger.Log.App.Info(
			"redis_connection_closed",
		)
	}
}
