package database

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(addr, password string, db int) {
	var opt *redis.Options
	var err error

	// 🔥 Support rediss:// (Upstash)
	if strings.HasPrefix(addr, "redis://") || strings.HasPrefix(addr, "rediss://") {
		opt, err = redis.ParseURL(addr)
		if err != nil {
			log.Fatal("Failed parse Redis URL:", err)
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
		log.Fatal("Failed connect Redis:", err)
	}

	log.Println("Connected to Redis")
}

func CloseRedis() {
	if Redis != nil {
		Redis.Close()
	}
}
