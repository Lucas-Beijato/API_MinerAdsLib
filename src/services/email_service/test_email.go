package emailserviceTest

import (
	"fmt"

	emailservice "ApiExtention.com/src/services/email_service/funcs"
)

// Test Func Send Email
func Send_Email_Test() {

	// Mensagem de e-mail
	to := []string{"lucasbeijato0@gmail.com", "andrei.maicon648@gmail.com", "ferzinhabastosmoreira@gmail.com"}
	subject := "Um recado de MinerAdsLib"
	body := "Olá, esse negócio tá saindo do papel, já to mandando emails! hahahahahaha!"

	err := emailservice.Send_Email(to, subject, body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email enviado com sucesso!")
	}
}
