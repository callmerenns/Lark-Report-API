package main

import "github.com/tsaqif-19/lark-report-api/cmd/server"

// @title           API Lark Webhook
// @version         1.0
// @description     API untuk menerima webhook dari Lark
// @termsOfService  https://example.com/terms

// @contact.name   API Support
// @contact.email  tsaqif@adamata.co

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:9200
// @BasePath  /

// 🔐 SECURITY DEFINITION (DI SINI)
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @Param Authorization header string true "Bearer JWT"
// @Param X-Webhook-Secret header string true "Static secret from Lark"

// @title Lark Report API
// @version 1.0
// @description Lark Report API Service
// @host g3p06lb7-9200.asse.devtunnels.ms
// @schemes https
// @BasePath /

func main() {
	server.Run()
}
