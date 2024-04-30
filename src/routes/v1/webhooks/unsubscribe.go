package webhooks

import (
	"fmt"
	"os"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	validatesignature "ApiExtention.com/src/services/validate_signature"
	paramssignaturetype "ApiExtention.com/src/types/params_signature"
	reqkiwifywhtype "ApiExtention.com/src/types/req_kiwify_wh"
	"github.com/gofiber/fiber/v2"
)

// For Unsubscribe
func Wh_Unsubscribe_Handler(c *fiber.Ctx) error {

	// Validate Signature
	key, isPresentKey := os.LookupEnv("TK_UNSUBSCRIBE")
	if !isPresentKey {
		fmt.Println("Unsubscribe Token Not Present")
		return c.SendStatus(400)
	}

	signature := new(paramssignaturetype.Params_Signature)
	if err := c.QueryParser(signature); err != nil {
		return err
	}
	bodyMessage := c.BodyRaw()
	isValidSignature := validatesignature.ValidateSignature(bodyMessage, []byte(signature.Signature), []byte(key))
	if !isValidSignature {
		fmt.Println("Not Valid Signature")
		return c.SendStatus(400)
	}
	// --------------------

	fmt.Println("[app]: Entrada no webhook '/unsubscribe'")

	b := new(reqkiwifywhtype.Req_Kiwify_Wh_Type)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body")
		return c.SendStatus(400)
	}

	if errToDelete := dbactionsservice.Delete_User(b); errToDelete != nil {
		fmt.Println("[app]: Error to delete in DB.")
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Usuário deletado do banco de dados.")

	to := []string{b.Costumer.Email, "mineradslib@gmail.com"}
	subject := "[Cancelado] - Um recado de MinerAdsLib."
	body := "Olá, sua inscrição foi excluída com sucesso!"

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	return c.SendStatus(200)
}
