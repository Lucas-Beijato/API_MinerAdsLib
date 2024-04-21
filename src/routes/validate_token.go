package routes

import (
	"fmt"

	dbactionsfuncs "ApiExtention.com/src/services/db_service/funcs/actions"
	req_res_types "ApiExtention.com/src/types"
	"github.com/gofiber/fiber/v2"
)

// Validate Token from client side
func Validate_Token_Handler(c *fiber.Ctx) error {
	b := new(req_res_types.Validate_Token_Body)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("Erro no parse body")
	}

	if b.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "Do you need to pass a token.",
			"isActive": false,
		})
	}

	if isValid := dbactionsfuncs.Query_Token(b.Token); !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"isActive": false,
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"isActive": true,
	})
}
