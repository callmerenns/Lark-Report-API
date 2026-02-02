package response

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    string `json:"code"`
	Details string `json:"details,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Window  string `json:"window,omitempty"`
}
