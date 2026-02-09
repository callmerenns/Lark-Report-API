package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/database"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"github.com/tsaqif-19/lark-report-api/internal/response"
	"go.uber.org/zap"
)

func JWT(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		ip := c.ClientIP()
		endpoint := c.FullPath()

		if auth == "" {
			logger.Log.Security.Warn(
				"missing_authorization_header",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
			)

			unauthorized(c, constant.ErrorMissingAuthorization, constant.MsgMissingAuthorization)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {

			logger.Log.Security.Warn(
				"invalid_authorization_format",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
			)

			unauthorized(c, constant.ErrorInvalidAuthFormat, constant.MsgInvalidAuthFormat)
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {

			logger.Log.Security.Warn(
				"invalid_jwt_token",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
				zap.Error(err),
			)

			unauthorized(c, constant.ErrorInvalidJWT, constant.MsgInvalidJWT)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			logger.Log.Security.Warn(
				"invalid_jwt_claims",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
			)

			unauthorized(c, constant.ErrorInvalidJWT, constant.MsgInvalidJWT)
			return
		}

		if claims["role"] != "lark_webhook" {

			logger.Log.Security.Warn(
				"forbidden_jwt_role",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
				zap.Any("role", claims["role"]),
			)

			forbidden(c, constant.ErrorForbiddenRole, constant.MsgForbiddenRole)
			return
		}

		// 🔐 VALIDASI JTI
		jti, ok := claims["jti"].(string)
		if !ok || jti == "" {

			logger.Log.Security.Warn(
				"missing_jti_claim",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
			)

			unauthorized(c, constant.ErrorInvalidJWT, "Missing token id")
			return
		}

		activeJTI, err := database.Redis.Get(
			context.Background(),
			"lark:active_jti",
		).Result()

		if err != nil || activeJTI != jti {

			logger.Log.Security.Warn(
				"jwt_token_revoked_or_mismatch",
				zap.String("ip", ip),
				zap.String("endpoint", endpoint),
				zap.String("jti", jti),
			)

			unauthorized(c, constant.ErrorInvalidJWT, "Token has been revoked")
			return
		}

		// ✅ JWT valid → lanjut (TIDAK LOG)
		c.Set("jwt_claims", claims)
		c.Next()
	}
}

func unauthorized(c *gin.Context, code, detail string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, response.APIResponse{
		Success: false,
		Message: "Unauthorized",
		Error: &response.APIError{
			Code:    code,
			Details: detail,
		},
	})
}

func forbidden(c *gin.Context, code, detail string) {
	c.AbortWithStatusJSON(http.StatusForbidden, response.APIResponse{
		Success: false,
		Message: "Forbidden",
		Error: &response.APIError{
			Code:    code,
			Details: detail,
		},
	})
}
