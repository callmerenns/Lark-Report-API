package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/response"
)

func WebhookSecret(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("X-Webhook-Secret")

		if secret == "" {
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
