package validatesignature

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

// Func - Validate Signature
func ValidateSignature(req_message []byte, req_message_signature []byte, token []byte) bool {

	// Método principal
	true_signature := hmac.New(sha1.New, token)
	true_signature.Write(req_message)
	expectedMAC := hex.EncodeToString(true_signature.Sum(nil))
	// fmt.Println("metodo1: ", expectedMAC)

	// fmt.Println("Assinatura recebida: ", string(req_message_signature))
	// fmt.Println("Assinatura Gerada pelo sistema: ", expectedMAC)

	return hmac.Equal(req_message_signature, []byte(expectedMAC))
}
