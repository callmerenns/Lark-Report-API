package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"go.uber.org/zap"
)

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		logger.Log.Access.Info(
			"http_request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}
