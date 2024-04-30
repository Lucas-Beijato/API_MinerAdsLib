package validatesignature

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// Func - Validate Signature
func ValidateSignature(req_message []byte, req_message_signature []byte, token []byte) bool {

	// Método principal
	true_signature := hmac.New(sha1.New, token)
	true_signature.Write(req_message)
	expectedMAC := hex.EncodeToString(true_signature.Sum(nil))
	fmt.Println("metodo1: ", expectedMAC)

	metodo224 := hmac.New(sha256.New224, token)
	metodo224.Write(req_message)
	expectedMAC224 := metodo224.Sum(nil)
	fmt.Println("metodo224: ", expectedMAC224)

	metodo256 := hmac.New(sha256.New, token)
	metodo256.Write(req_message)
	expectedMAC256 := metodo256.Sum(nil)
	fmt.Println("metodo256: ", expectedMAC256)

	metodo384 := hmac.New(sha512.New384, token)
	metodo384.Write(req_message)
	expectedMAC384 := metodo384.Sum(nil)
	fmt.Println("metodo384: ", expectedMAC384)

	metodo512 := hmac.New(sha512.New, token)
	metodo512.Write(req_message)
	expectedMAC512 := metodo512.Sum(nil)
	fmt.Println("metodo512: ", expectedMAC512)

	fmt.Println("Assinatura recebida: ", req_message_signature)

	return hmac.Equal(req_message_signature, []byte(expectedMAC))
}
