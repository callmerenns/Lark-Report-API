package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/database"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"github.com/tsaqif-19/lark-report-api/internal/response"
	"go.uber.org/zap"
)

func SimpleRateLimiter(
	cfg *config.Config,
	prefix string,
	limit int,
	window time.Duration,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		rdb := database.Redis
		if rdb == nil {
			logger.Log.Error.Error(
				"redis_not_initialized",
				zap.String("middleware", "simple_rate_limiter"),
			)

			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorInternalServer,
					Details: "Redis not initialized",
				},
			})
			return
		}

		ctx := context.Background()
		ip := c.ClientIP()
		key := fmt.Sprintf("rate:%s:%s", prefix, ip)

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			logger.Log.Error.Error(
				"rate_limiter_redis_error",
				zap.Error(err),
				zap.String("key", key),
				zap.String("ip", ip),
			)

			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorInternalServer,
					Details: "Rate limiter error",
				},
			})
			return
		}

		if count == 1 {
			_ = rdb.Expire(ctx, key, window).Err()
		}

		// DEV LOG (AMAN)
		if cfg.AppEnv == "dev" {
			logger.Log.App.Info(
				"rate_limiter_hit",
				zap.String("key", key),
				zap.Int64("count", count),
				zap.Int("limit", limit),
			)
		}

		if count > int64(limit) {
			logger.Log.Security.Warn(
				"rate_limit_exceeded",
				zap.String("prefix", prefix),
				zap.String("ip", ip),
				zap.Int64("count", count),
				zap.Int("limit", limit),
				zap.Duration("window", window),
				zap.String("path", c.FullPath()),
				zap.String("method", c.Request.Method),
			)

			c.Header("Retry-After", strconv.Itoa(int(window.Seconds())))

			c.AbortWithStatusJSON(http.StatusTooManyRequests, response.APIResponse{
				Success: false,
				Message: "Request blocked",
				Error: &response.APIError{
					Code:    constant.ErrorRateLimited,
					Details: constant.MsgTooManyRequests,
				},
			})
			return
		}

		c.Next()
	}
}
