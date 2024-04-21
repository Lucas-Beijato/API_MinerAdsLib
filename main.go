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
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		return c.SendString("Hello Word!")
	})

	// WebHooks
	v1.Post("/newSale/", webhooks.New_Sale_Handle)

	// Routes
	v1.Post("/validate_token/", routes.Validate_Token_Handle)

	// Start Server
	fmt.Println("[App]: Essa bagaça ta rodando! 🚀 ")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
