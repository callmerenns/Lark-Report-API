package router

import (
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tsaqif-19/lark-report-api/docs"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/handler"
	"github.com/tsaqif-19/lark-report-api/internal/middleware"
)

func RegisterRoutes(
	r *gin.Engine,
	cfg *config.Config,
	larkHandler *handler.LarkHandler,
	healthHandler *handler.HealthHandler,
	tokenHandler *handler.TokenHandler,
) {

	// Root
	r.GET("/", healthHandler.Welcome)

	// Webhook (JWT + Secret + Rate Limit)
	webhook := r.Group("/webhook")
	webhook.Use(
		middleware.LuaRateLimiter(cfg, "lark_webhook", 30, time.Minute),
		middleware.WebhookSecret(cfg),
		middleware.JWT(cfg),
	)
	{
		webhook.POST("/lark", larkHandler.HandleWebhook)
	}

	// Internal (SUPER STRICT)
	internal := r.Group("/internal")
	internal.Use(
		middleware.LuaRateLimiter(cfg, "internal", 5, time.Minute),
		middleware.WebhookSecret(cfg),
	)
	{
		internal.GET("/generate-lark-token", tokenHandler.GenerateLarkToken)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
