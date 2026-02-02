package response

type GenerateTokenResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Token generated successfully"`
	Data    struct {
		AccessToken string `json:"access_token" example:"eyJhbGciOi..."`
		TokenType   string `json:"token_type" example:"Bearer"`
		ExpiresIn   int64  `json:"expires_in" example:"31536000"`
	} `json:"data"`
}
