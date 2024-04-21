package main

import (
	"fmt"
	"log"
	"os"

	"ApiExtention.com/src/routes"
	"ApiExtention.com/src/routes/v1/webhooks"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := fiber.New()

	// v1 version
	v1 := app.Group("/v1/")

	// WebHooks
	v1.Post("/new_sale", webhooks.New_Sale_Handler)
	v1.Post("/unsubscribe", webhooks.Wh_Nulled_Sub_Handler)
	v1.Post("/overdue_subscription", webhooks.Wh_Overdue_Sub_Handler)
	v1.Post("/renewed_subscription", webhooks.Wh_Renewed_Sub_Handler)

	// Routes
	v1.Post("/validate_token", routes.Validate_Token_Handler)

	// Start Server
	fmt.Println("[App]: Essa bagaça ta rodando! 🚀 ")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
