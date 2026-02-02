package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/database"
)

const luaScript = `
redis.call('ZREMRANGEBYSCORE', KEYS[1], 0, ARGV[2])
redis.call('ZADD', KEYS[1], ARGV[1], ARGV[1])

local count = redis.call('ZCARD', KEYS[1])

if count == 1 then
	redis.call('EXPIRE', KEYS[1], ARGV[4])
end

if count > tonumber(ARGV[3]) then
	return {0, count}
end

return {1, count}
`

func LuaRateLimiter(
	cfg *config.Config,
	prefix string,
	limit int,
	window time.Duration,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		rdb := database.Redis
		if rdb == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorRateLimited,
					Details: "Redis not initialized",
				},
			})
			return
		}

		ctx := context.Background()

		key := fmt.Sprintf(
			"rate:%s:%s",
			prefix,
			c.ClientIP(),
		)

		now := time.Now().UnixNano()
		windowStart := time.Now().Add(-window).UnixNano()

		result, err := rdb.Eval(
			ctx,
			luaScript,
			[]string{key},
			now,
			windowStart,
			limit,
			int(window.Seconds()),
		).Result()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorRateLimited,
					Details: "Rate limiter execution failed",
				},
			})
			return
		}

		// SAFE CAST
		res, ok := result.([]interface{})
		if !ok || len(res) == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorRateLimited,
					Details: "Invalid rate limiter response",
				},
			})
			return
		}

		allowed, ok := res[0].(int64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.APIResponse{
				Success: false,
				Message: "Internal server error",
				Error: &response.APIError{
					Code:    constant.ErrorRateLimited,
					Details: "Invalid rate limiter flag",
				},
			})
			return
		}

		if allowed == 0 {
			c.Header("Retry-After", strconv.Itoa(int(window.Seconds())))

			errorBody := &response.APIError{
				Code:    constant.ErrorRateLimited,
				Details: constant.MsgTooManyRequests,
			}

			// DEV ONLY DEBUG INFO
			if cfg.AppEnv == "dev" {
				errorBody.Window = window.String()
				errorBody.Limit = limit
			}

			c.AbortWithStatusJSON(http.StatusTooManyRequests, response.APIResponse{
				Success: false,
				Message: "Request blocked",
				Error:   errorBody,
			})
			return
		}

		c.Next()
	}
}
