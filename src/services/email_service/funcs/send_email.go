package emailservice

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
)

// Send Email
func Send_Email(to []string, subject string, body string) error {

	smtpServer, is_present_smtpServer := os.LookupEnv("SMTP_SERVER")
	smtpPort, is_present_smtpPort := os.LookupEnv("SMTP_PORT")
	senderEmail, is_present_senderEmail := os.LookupEnv("SENDER_EMAIL")
	senderPassword, is_present_senderPassword := os.LookupEnv("SENDER_PASS")

	switch {
	case !is_present_smtpServer:
		{
			fmt.Println("SMTP SERVER IS NOT DEFINED IN AMBIENT VARIABLES")
			return errors.New("SMTP SERVER IS NOT DEFINED IN AMBIENT VARIABLES")
		}
	case !is_present_smtpPort:
		{
			fmt.Println("SMTP PORT IS NOT DEFINED IN AMBIENT VARIABLES")
			return errors.New("SMTP PORT IS NOT DEFINED IN AMBIENT VARIABLES")
		}
	case !is_present_senderEmail:
		{
			fmt.Println("SMTP SENDER EMAIL ADDRESS IS NOT DEFINED IN AMBIENT VARIABLES")
			return errors.New("SMTP SENDER EMAIL ADDRESS IS NOT DEFINED IN AMBIENT VARIABLES")
		}
	case !is_present_senderPassword:
		{
			fmt.Println("SMTP SENDER EMAIL PASSWORD IS NOT DEFINED IN AMBIENT VARIABLES")
			return errors.New("SMTP SENDER EMAIL PASSWORD IS NOT DEFINED IN AMBIENT VARIABLES")
		}
	}

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	msg := "To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	err_to_send_email := smtp.SendMail(fmt.Sprintf("%s:%s", smtpServer, smtpPort), auth, senderEmail, to, []byte(msg))
	if err_to_send_email != nil {
		return errors.New("SMTP SENDER EMAIL PASSWORD IS NOT DEFINED IN AMBIENT VARIABLES")
	}

	return nil

}
