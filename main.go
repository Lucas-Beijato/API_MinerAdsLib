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
	v1 := app.Group("/v1")

	Handler_newSale := webhooks.New_Sale_Handle
	Handler_Val_tok := routes.Validate_Token_Handle

	// WebHooks
	v1.Post("/newSale", Handler_newSale)

	// Routes
	v1.Post("/validate_token", Handler_Val_tok)

	// Start Server
	fmt.Println("[App]: Essa bagaça ta rodando! 🚀 ")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
