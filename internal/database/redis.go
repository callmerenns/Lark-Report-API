package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(addr, password string, db int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed connect Redis:", err)
	}

	log.Println("Connected to Redis")
}

func CloseRedis() {
	if Redis != nil {
		Redis.Close()
	}
}
