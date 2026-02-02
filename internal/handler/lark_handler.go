package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsaqif-19/lark-report-api/internal/domain"
	"github.com/tsaqif-19/lark-report-api/internal/response"
	"github.com/tsaqif-19/lark-report-api/internal/service"
)

type LarkHandler struct {
	service *service.RecordService
}

func NewLarkHandler(s *service.RecordService) *LarkHandler {
	return &LarkHandler{service: s}
}

// HandleWebhook godoc
// @Summary      Lark Webhook
// @Description  Receive webhook from Lark and create record
// @Tags         Webhook
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        X-Webhook-Secret header string true "Static webhook secret"
// @Param        body body dto.LarkWebhookRequest true "Payload from Lark"
// @Success 	 201 {object} response.RecordCreatedSuccessExample "Record created successfully"
// @Failure		 400 {object} response.BadRequestErrorExample "Invalid payload"
// @Failure		 401 {object} response.UnauthorizedErrorExample "Unauthorized"
// @Failure		 403 {object} response.ForbiddenErrorExample "Forbidden"
// @Failure		 429 {object} response.RateLimitErrorExample "Rate limited"
// @Failure		 500 {object} response.InternalServerErrorExample "Internal server error"
// @Router       /webhook/lark [post]
func (h *LarkHandler) HandleWebhook(c *gin.Context) {
	var req struct {
		Data domain.Record `json:"data"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: "Invalid request payload",
			Error: &response.APIError{
				Code:    "INVALID_JSON",
				Details: err.Error(),
			},
		})
		return
	}

	id, err := h.service.CreateRecord(c.Request.Context(), req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.APIResponse{
			Success: false,
			Message: "Failed to create record",
			Error: &response.APIError{
				Code: "DATABASE_ERROR",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Success: true,
		Message: "Record created successfully",
		Data: response.RecordCreatedData{
			ID:                  id,
			IncidentTitle:       req.Data.IncidentTitle,
			IncidentDescription: req.Data.IncidentDescription,
			MachineModel:        req.Data.MachineModel,
			LightboxType:        req.Data.LightboxType,
			SetType:             req.Data.SetType,
			TypesOfSpareparts:   req.Data.TypesOfSpareparts,
			HowToHandle:         req.Data.HowToHandle,
			CollectDataset:      req.Data.HowToHandle,
			ProcessingModel:     req.Data.ProcessingModel,
			DateOfIncident:      req.Data.DateOfIncident,
			TimeOfIncident:      req.Data.TimeOfIncident,
			DayOfIncident:       req.Data.DayOfIncident,
			ProcessReason:       req.Data.ProcessReason,
			CreatedBy:           req.Data.CreatedBy,
			ModifiedBy:          req.Data.ModifiedBy,
			Status:              req.Data.Status,
			CreatedAt:           time.Now().UTC(),
		},
		Error: nil,
	})
}
