package webhooks

import (
	"fmt"
	"os"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	tokengenservice "ApiExtention.com/src/services/token_gen_service/funcs"
	validatesignature "ApiExtention.com/src/services/validate_signature"
	req_res_types "ApiExtention.com/src/types"
	"github.com/gofiber/fiber/v2"
)

// Handle for a new sale
func New_Sale_Handler(c *fiber.Ctx) error {

	fmt.Println("[app]: Entrada no webhook '/new_sale'")

	b := new(req_res_types.KiwifyResponse)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body")
		return c.SendStatus(400)
	}

	// Validate Signature
	key, isPresentKey := os.LookupEnv("TK_NEW_SALE")
	if !isPresentKey {
		fmt.Println("New Sale Token Not Present")
		return c.SendStatus(400)
	}
	signature := new(req_res_types.ParamsSignature)
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

	// DB
	if err := dbactionsservice.Isert_New_User_in_DB(&New_User); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Adicionado ao banco de dados.")

	// EMAIL
	to := []string{New_User.Data_User.Costumer.Email, "mineradslib@gmail.com"}
	subject := "[Registrado] - Um recado de MinerAdsLib."
	// body := "Olá, você foi registrado dentro do nosso sistema, seu token de acesso a extensão é: " + token
	body := fmt.Sprintf("Olá, você foi registrado dentro do nosso sistema.\n\nSeu Token é: %s", token)

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	// FINAL RESPONSE
	return c.SendStatus(200)
}
