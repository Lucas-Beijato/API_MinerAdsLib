package tokengenservice

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// To load easy the secret key
func Get_secret_key() ([]byte, error) {
	secret_key, is_present_secret_key := os.LookupEnv("SECRET")
	if !is_present_secret_key {
		return nil, errors.New("erro ao encontrar a SECRET")
	}
	secret_key_bytes := []byte(secret_key)

	return secret_key_bytes, nil
}

// Validate Token
func Validate_Token(tokenString string) error {

	secretKey, err := Get_secret_key()
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// Gen a New Token
func Gen_Token(t_id *string) (string, error) {

	secret_key_bytes, err := Get_secret_key()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"transaction_ID": t_id,
		// "exp":            time.Now().Add(time.Hour * 24 * 32).Unix(),
		"exp": time.Now().Add(time.Minute + 2).Unix(),
	})

	tokenString, err := token.SignedString(secret_key_bytes)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
