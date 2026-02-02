package constant

const (
	// Webhoox
	ErrorInvalidWebhookSecret = "INVALID_WEBHOOK_SECRET"
	ErrorMissingWebhookSecret = "MISSING_WEBHOOK_SECRET"

	// Auth & Security
	ErrorMissingAuthorization = "MISSING_AUTHORIZATION"
	ErrorInvalidAuthFormat    = "INVALID_AUTH_FORMAT"
	ErrorInvalidJWT           = "INVALID_JWT"
	ErrorForbiddenRole        = "FORBIDDEN_ROLE"

	// Rate Limit
	ErrorRateLimited = "RATE_LIMITED"

	// Payload
	ErrorInvalidPayload = "INVALID_PAYLOAD"

	// Server
	ErrorInternalServer = "INTERNAL_SERVER_ERROR"
)
