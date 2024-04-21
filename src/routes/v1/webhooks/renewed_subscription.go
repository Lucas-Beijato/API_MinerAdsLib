package webhooks

import "github.com/gofiber/fiber/v2"

// For Renewed Subscription
func Wh_Renewed_Sub_Handler(c *fiber.Ctx) error {

	// Para implementar

	return c.SendStatus(200)
}
