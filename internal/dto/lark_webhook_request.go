package dto

import "github.com/tsaqif-19/lark-report-api/internal/domain"

type LarkWebhookRequest struct {
	Data domain.Record `json:"data"`
}
