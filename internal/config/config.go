package config

import (
	"os"

	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

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

	logger.Log.App.Info(
		"loading_application_config",
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

		logger.Log.App.Warn(
			"port_not_set_using_default",
			zap.String("default_port", port),
		)
	}

	cfg := &Config{
		AppEnv:        os.Getenv("APP_ENV"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		WebhookSecret: os.Getenv("WEBHOOK_SECRET"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       0,
		Port:          port,
	}

	// ⚠️ Validasi config penting
	validateConfig(cfg)

	logger.Log.App.Info(
		"application_config_loaded",
		zap.String("env", cfg.AppEnv),
		zap.String("port", cfg.Port),
		zap.Bool("redis_enabled", cfg.RedisAddr != ""),
	)

	return cfg
}

func validateConfig(cfg *Config) {

	if cfg.AppEnv == "" {
		logger.Log.Error.Error(
			"app_env_not_set",
		)
		panic("APP_ENV is required")
	}

	if cfg.DatabaseURL == "" {
		logger.Log.Error.Error(
			"database_url_not_set",
		)
		panic("DATABASE_URL is required")
	}

	if cfg.JWTSecret == "" {
		logger.Log.Error.Error(
			"jwt_secret_not_set",
		)
		panic("JWT_SECRET is required")
	}
}
