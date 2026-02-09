package database

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

func NewPostgres(cfg *config.Config) *pgxpool.Pool {

	logger.Log.App.Info(
		"initializing_postgres_connection",
	)

	dsn := cfg.DatabaseURL
	if dsn == "" {
		logger.Log.Error.Error(
			"database_url_not_set",
		)
		panic("DATABASE_URL is required")
	}

	// enforce sslmode
	if !strings.Contains(dsn, "sslmode=") {
		if strings.Contains(dsn, "?") {
			dsn += "&sslmode=require"
		} else {
			dsn += "?sslmode=require"
		}
	}

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Log.Error.Error(
			"failed_to_parse_database_dsn",
			zap.Error(err),
		)
		panic(err)
	}

	// pool tuning
	poolCfg.MaxConns = 10
	poolCfg.MinConns = 2
	poolCfg.MaxConnIdleTime = 5 * time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		logger.Log.Error.Error(
			"failed_to_create_postgres_pool",
			zap.Error(err),
		)
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		logger.Log.Error.Error(
			"failed_to_ping_postgres",
			zap.Error(err),
		)
		panic(err)
	}

	logger.Log.App.Info(
		"postgres_connected",
		zap.Int32("max_conns", poolCfg.MaxConns),
		zap.Int32("min_conns", poolCfg.MinConns),
	)

	return db
}
