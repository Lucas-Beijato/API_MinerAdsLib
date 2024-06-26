package webhooks

import (
	"fmt"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	reqkiwifywhtype "ApiExtention.com/src/types/req_kiwify_wh"
	"github.com/gofiber/fiber/v2"
)

// For Overdue Subscription
func Wh_Overdue_Sub_Handler(c *fiber.Ctx) error {

	// // Validate Signature
	// key, isPresentKey := os.LookupEnv("TK_OVERDUE_SUBSCRIPTION")
	// if !isPresentKey {
	// 	fmt.Println("Overdue Subscription Token Not Present")
	// 	return c.SendStatus(400)
	// }

	// signature := new(paramssignaturetype.Params_Signature)
	// if err := c.QueryParser(signature); err != nil {
	// 	return err
	// }
	// bodyMessage := c.BodyRaw()
	// isValidSignature := validatesignature.ValidateSignature(bodyMessage, []byte(signature.Signature), []byte(key))
	// if !isValidSignature {
	// 	fmt.Println("Not Valid Signature")
	// 	return c.SendStatus(400)
	// }
	// // --------------------

	fmt.Println("[app]: Entrada no webhook '/overdue_subscription'")

	b := new(reqkiwifywhtype.Req_Kiwify_Wh_Type)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body: ", err)
		return c.SendStatus(400)
	}

	if errToCleanToken := dbactionsservice.Clean_Token_User(b); errToCleanToken != nil {
		fmt.Println("[app]: Error to clean token in DB.")
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Token resetado!")

	to := []string{b.Costumer.Email, "mineradslib@gmail.com"}
	subject := "[Em Atraso] - Um recado de MinerAdsLib."
	body := "Olá, sua assinatura foi suspensa por falta de pagamento, entre em contato conosco!"

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	return c.SendStatus(200)
}
