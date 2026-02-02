package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tsaqif-19/lark-report-api/internal/config"
	"github.com/tsaqif-19/lark-report-api/internal/constant"
	"github.com/tsaqif-19/lark-report-api/internal/database"
	"github.com/tsaqif-19/lark-report-api/internal/response"
)

type TokenHandler struct {
	cfg *config.Config
}

func NewTokenHandler(cfg *config.Config) *TokenHandler {
	return &TokenHandler{cfg: cfg}
}

// GenerateLarkToken godoc
// @Summary      Generate JWT token for Lark webhook
// @Description  Generate JWT token used by Lark to access webhook endpoint
// @Tags         Internal
// @Produce      json
// @Param        X-Webhook-Secret header string true "Static webhook secret"
// @Success      200 {object} response.GenerateTokenResponse
// @Failure		 400 {object} response.BadRequestErrorExample "Invalid payload"
// @Failure		 401 {object} response.UnauthorizedErrorExample "Unauthorized"
// @Failure		 429 {object} response.RateLimitErrorExample "Rate limited"
// @Failure		 500 {object} response.InternalServerErrorExample "Internal server error"
// @Router       /internal/generate-lark-token [get]
func (h *TokenHandler) GenerateLarkToken(c *gin.Context) {
	now := time.Now()
	expiresAt := now.Add(9 * time.Hour)
	jti := uuid.NewString()

	claims := jwt.MapClaims{
		"iss":  "lark",
		"role": "lark_webhook",
		"exp":  expiresAt.Unix(),
		"iat":  now.Unix(),
		"jti":  jti,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Success: false,
			Message: "Failed to generate token",
			Error: &response.APIError{
				Code:    constant.ErrorInternalServer,
				Details: "JWT signing failed",
			},
		})
		return
	}

	// SIMPAN JTI AKTIF (invalidate token lama)
	key := "lark:active_jti"
	if err := database.Redis.Set(
		context.Background(),
		key,
		jti,
		time.Until(expiresAt),
	).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Success: false,
			Message: "Failed to store token",
			Error: &response.APIError{
				Code:    constant.ErrorInternalServer,
				Details: "Redis error",
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Token generated successfully",
		Data: gin.H{
			"access_token": tokenStr,
			"token_type":   "Bearer",
			"expires_in":   int64(time.Until(expiresAt).Seconds()),
		},
		Error: nil,
	})
}
