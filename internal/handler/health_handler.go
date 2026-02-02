package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/response"
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
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Welcome to API Lark",
		Data: gin.H{
			"name":    "Lark Webhook API",
			"version": "v1",
			"status":  "running",
		},
		Error: nil,
	})
}
