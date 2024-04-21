package webhooks

import (
	"fmt"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	tokengenservice "ApiExtention.com/src/services/token_gen_service/funcs"
	req_res_types "ApiExtention.com/src/types"
	"github.com/gofiber/fiber/v2"
)

// For Renewed Subscription
func Wh_Renewed_Sub_Handler(c *fiber.Ctx) error {

	fmt.Println("[app]: Entrada no webhook '/renewed_subscription'")

	b := new(req_res_types.KiwifyResponse)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body")
		return c.SendStatus(400)
	}

	// GEN TOKEN
	token, errToGenToken := tokengenservice.Gen_Token(&b.Subscription_ID)
	if errToGenToken != nil {
		fmt.Println(errToGenToken)
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Token gerado.")

	New_User := req_res_types.User{
		Data_User: b,
		Token:     token,
	}

	if errToCleanToken := dbactionsservice.Update_Token_User(&New_User); errToCleanToken != nil {
		fmt.Println("[app]: Error to clean token in DB.")
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Token Atualizado no sistema!")

	to := []string{"lucasbeijato0@gmail.com"}
	subject := "Um recado de MinerAdsLib."
	body := "Olá, sua assinatura foi renovada. Segue seu novo Token: " + token

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	return c.SendStatus(200)
}
