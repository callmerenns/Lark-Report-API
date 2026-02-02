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
	"github.com/tsaqif-19/lark-report-api/internal/response"
)

func JWT(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			unauthorized(c, constant.ErrorMissingAuthorization, constant.MsgMissingAuthorization)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			unauthorized(c, constant.ErrorInvalidAuthFormat, constant.MsgInvalidAuthFormat)
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			unauthorized(c, constant.ErrorInvalidJWT, constant.MsgInvalidJWT)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			unauthorized(c, constant.ErrorInvalidJWT, constant.MsgInvalidJWT)
			return
		}

		if claims["role"] != "lark_webhook" {
			forbidden(c, constant.ErrorForbiddenRole, constant.MsgForbiddenRole)
			return
		}

		// VALIDASI JTI
		jti, ok := claims["jti"].(string)
		if !ok || jti == "" {
			unauthorized(c, constant.ErrorInvalidJWT, "Missing token id")
			return
		}

		activeJTI, err := database.Redis.Get(
			context.Background(),
			"lark:active_jti",
		).Result()

		if err != nil || activeJTI != jti {
			unauthorized(c, constant.ErrorInvalidJWT, "Token has been revoked")
			return
		}

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
