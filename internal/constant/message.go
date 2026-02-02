package constant

const (
	MsgTooManyRequests      = "Too many requests"
	MsgMissingWebhookSecret = "Missing X-Webhook-Secret header"
	MsgInvalidWebhookSecret = "Invalid webhook secret"
	MsgMissingAuthorization = "Missing Authorization header"
	MsgInvalidAuthFormat    = "Invalid Authorization format"
	MsgInvalidJWT           = "Invalid or expired token"
	MsgForbiddenRole        = "Invalid token role"
)
