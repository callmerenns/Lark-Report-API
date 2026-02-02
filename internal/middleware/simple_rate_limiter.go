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
	"github.com/tsaqif-19/lark-report-api/internal/response"
)

func SimpleRateLimiter(
	cfg *config.Config,
	prefix string,
	limit int,
	window time.Duration,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		rdb := database.Redis
		ctx := context.Background()

		key := fmt.Sprintf("rate:%s:%s", prefix, c.ClientIP())

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorInternalServer,
					Details: err.Error(),
				},
			})
			return
		}

		// set TTL only on first request
		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		if count > int64(limit) {
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
