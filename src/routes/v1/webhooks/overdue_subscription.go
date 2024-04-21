package webhooks

import (
	"fmt"

	req_res_types "ApiExtention.com/src/types"
	"github.com/gofiber/fiber/v2"
)

// For Overdue Subscription
func Wh_Overdue_Sub_Handler(c *fiber.Ctx) error {
	fmt.Println("[app]: Entrada no webhook '/unsubscribe'")

	b := new(req_res_types.KiwifyResponse)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body")
		return c.SendStatus(400)
	}

	return c.SendStatus(200)
}
