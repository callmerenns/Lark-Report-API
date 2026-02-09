package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/logger"
	"github.com/tsaqif-19/lark-report-api/internal/response"
	"go.uber.org/zap"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Welcome godoc
// @Summary Welcome endpoint
// @Description Health check & welcome message
// @Tags Health
// @Produce json
// @Success 200 {object} response.WelcomeSuccessExample "Welcome to API Lark"
// @Router / [get]
func (h *HealthHandler) Welcome(c *gin.Context) {

	logger.Log.App.Info(
		"health_check",
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.String("service", "lark-api"),
		zap.String("status", "running"),
		zap.String("endpoint", c.FullPath()),
	)

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Welcome to API Lark",
		Data: gin.H{
			"name":       "Lark Webhook API",
			"version":    "v1",
			"status":     "running",
			"created_by": "Al Tsaqif",
		},
		Error: nil,
	})
}
