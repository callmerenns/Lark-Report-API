package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

func RecoveryLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error.Error(
					"panic_recovered",
					zap.Any("error", err),
					zap.String("path", c.FullPath()),
					zap.String("ip", c.ClientIP()),
				)

				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
