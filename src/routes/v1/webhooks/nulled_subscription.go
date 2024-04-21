package webhooks

import (
	"fmt"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	req_res_types "ApiExtention.com/src/types"
	"github.com/gofiber/fiber/v2"
)

// For Nulled Subscriptions
func Wh_nulled_sub_Handler(c *fiber.Ctx) error {

	fmt.Println("[app]: Entrada no webhook '/unsubscribe'")

	b := new(req_res_types.KiwifyResponse)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body")
		return c.SendStatus(400)
	}

	if errToDelete := dbactionsservice.Delete_User(b); errToDelete != nil {
		fmt.Println("[app]: Error to delete in DB.")
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Usuário deletado do banco de dados.")

	to := []string{"lucasbeijato0@gmail.com"}
	subject := "Um recado de MinerAdsLib."
	body := "Olá, sua inscrição foi excluída com sucesso!"

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	return c.SendStatus(200)
}
