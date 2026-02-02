package config

import "os"

type Config struct {
	AppEnv        string
	DatabaseURL   string
	JWTSecret     string
	Port          string
	WebhookSecret string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		AppEnv:        os.Getenv("APP_ENV"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		WebhookSecret: os.Getenv("WEBHOOK_SECRET"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       0,
		Port:          port,
	}
}
