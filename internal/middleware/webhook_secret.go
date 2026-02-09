package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"github.com/tsaqif-19/lark-report-api/internal/response"
	"go.uber.org/zap"
)

func WebhookSecret(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("X-Webhook-Secret")
		ip := c.ClientIP()

		if secret == "" {
			logger.Log.Security.Warn(
				"missing_webhook_secret",
				zap.String("ip", ip),
				zap.String("path", c.FullPath()),
				zap.String("method", c.Request.Method),
			)

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.APIResponse{
					Success: false,
					Message: "Unauthorized",
					Error: &response.APIError{
						Code:    constant.ErrorMissingWebhookSecret,
						Details: constant.MsgMissingWebhookSecret,
					},
				},
			)
			return
		}

		if secret != cfg.WebhookSecret {
			logger.Log.Security.Warn(
				"invalid_webhook_secret",
				zap.String("ip", ip),
				zap.String("path", c.FullPath()),
				zap.String("method", c.Request.Method),
			)

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.APIResponse{
					Success: false,
					Message: "Unauthorized",
					Error: &response.APIError{
						Code:    constant.ErrorInvalidWebhookSecret,
						Details: constant.MsgInvalidWebhookSecret,
					},
				},
			)
			return
		}

		c.Next()
	}
}
