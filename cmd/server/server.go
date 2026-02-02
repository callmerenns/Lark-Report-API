package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/database"
	"github.com/tsaqif-19/lark-report-api/internal/handler"
	"github.com/tsaqif-19/lark-report-api/internal/repository"
	"github.com/tsaqif-19/lark-report-api/internal/router"
	"github.com/tsaqif-19/lark-report-api/internal/service"
)

func Run() {
	// Load ENV variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	// INIT DATABASE
	db := database.NewPostgres(cfg)

	defer db.Close()
	// INIT REDIS (WAJIB SEBELUM ROUTER)
	database.InitRedis(
		cfg.RedisAddr,
		cfg.RedisPassword,
		cfg.RedisDB,
	)
	defer database.CloseRedis()

	// Dependency Injection
	repo := repository.NewRecordRepository(db)
	svc := service.NewRecordService(repo)

	healthHandler := handler.NewHealthHandler()
	larkHandler := handler.NewLarkHandler(svc)
	tokenHandler := handler.NewTokenHandler(cfg)

	r := gin.Default()

	router.RegisterRoutes(
		r,
		cfg,
		larkHandler,
		healthHandler,
		tokenHandler,
	)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
