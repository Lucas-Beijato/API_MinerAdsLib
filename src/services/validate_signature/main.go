package validatesignature

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// Func - Validate Signature
func ValidateSignature(req_message []byte, req_message_signature []byte, token []byte) bool {
	true_signature := hmac.New(sha256.New, token)
	true_signature.Write(req_message)
	expectedMAC := true_signature.Sum(nil)

	fmt.Println(req_message_signature, expectedMAC)
	return hmac.Equal(req_message_signature, expectedMAC)
}
