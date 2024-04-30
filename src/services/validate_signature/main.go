package validatesignature

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

// Func - Validate Signature
func ValidateSignature(req_message []byte, req_message_signature []byte, token []byte) bool {
	true_signature := hmac.New(sha1.New, token)
	true_signature.Write(req_message)
	expectedMAC := true_signature.Sum(nil)

	fmt.Println("Mensagem recebida", string(req_message))
	fmt.Println("Assinatura recebida", string(req_message_signature))
	fmt.Println("Token do sistema \n", string(token))
	fmt.Println("Comparando as chaves:")
	fmt.Println(req_message_signature)
	fmt.Println(expectedMAC)

	return hmac.Equal(req_message_signature, expectedMAC)
}
