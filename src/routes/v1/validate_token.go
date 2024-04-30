package routes

import (
	"fmt"

	dbactionsfuncs "ApiExtention.com/src/services/db_service/funcs/actions"
	tokengenservice "ApiExtention.com/src/services/token_gen_service/funcs"
	tokentype "ApiExtention.com/src/types/token"
	"github.com/gofiber/fiber/v2"
)

// Validate Token from client side
func Validate_Token_Handler(c *fiber.Ctx) error {
	b := new(tokentype.Validate_Token_Body)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("Erro no parse body")
	}

	if b.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    "Do you need to pass a token.",
			"isActive": false,
		})
	}

	isValid_token_db := dbactionsfuncs.Query_Token(b.Token)
	isValid_token_verification := tokengenservice.Validate_Token(b.Token)

	if !isValid_token_db || !isValid_token_verification {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"isActive": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"isActive": true,
	})
}
