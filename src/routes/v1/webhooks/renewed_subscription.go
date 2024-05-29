package webhooks

import (
	"fmt"

	dbactionsservice "ApiExtention.com/src/services/db_service/funcs/actions"
	emailservice "ApiExtention.com/src/services/email_service/funcs"
	tokengenservice "ApiExtention.com/src/services/token_gen_service/funcs"
	reqkiwifywhtype "ApiExtention.com/src/types/req_kiwify_wh"
	usertype "ApiExtention.com/src/types/user"
	"github.com/gofiber/fiber/v2"
)

// For Renewed Subscription
func Wh_Renewed_Sub_Handler(c *fiber.Ctx) error {

	// // Validate Signature
	// key, isPresentKey := os.LookupEnv("TK_RENEWED_SUBSCRIPTION")
	// if !isPresentKey {
	// 	fmt.Println("Renewed Subscription Token Not Present")
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

	fmt.Println("[app]: Entrada no webhook '/renewed_subscription'")

	b := new(reqkiwifywhtype.Req_Kiwify_Wh_Type)
	if err := c.BodyParser(b); err != nil {
		fmt.Println("[app]: Error to parse body: ", err)
		return c.SendStatus(400)
	}

	// GEN TOKEN
	token, errToGenToken := tokengenservice.Gen_Token(&b.Subscription_id)
	if errToGenToken != nil {
		fmt.Println(errToGenToken)
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Token gerado: ", token)

	New_User := usertype.User{
		Data_User:       b,
		Token:           token,
		Subscription_ID: b.Subscription_id,
	}

	if errToCleanToken := dbactionsservice.Update_Token_User(&New_User); errToCleanToken != nil {
		fmt.Println("[app]: Error to clean token in DB.")
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Token Atualizado no DB!")

	to := []string{New_User.Data_User.Costumer.Email, "mineradslib@gmail.com"}
	subject := "[Renovado] - Um recado de MinerAdsLib."
	body := "Olá, sua assinatura foi renovada.\n\nSegue seu novo Token: " + token

	if err := emailservice.Send_Email(to, subject, body); err != nil {
		return c.SendStatus(500)
	}

	fmt.Println("[app]: Email enviado.")

	return c.SendStatus(200)
}
