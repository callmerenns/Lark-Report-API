package router

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
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

	// ✅ HANDLE PREFLIGHT
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ✅ CORS FIX (DEVTUNNELS COMPATIBLE)
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return strings.HasSuffix(origin, ".devtunnels.ms")
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Root
	r.GET("/", healthHandler.Welcome)

	// Webhook
	webhook := r.Group("/webhook")
	webhook.Use(
		middleware.LuaRateLimiter(cfg, "lark_webhook", 30, time.Minute),
		middleware.WebhookSecret(cfg),
		middleware.JWT(cfg),
	)
	{
		webhook.POST("/lark", larkHandler.HandleWebhook)
	}

	// Internal
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
