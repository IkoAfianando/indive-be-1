package utils

import (
	"gopkg.in/mail.v2"
	"os"
	"strconv"
)

func SendVerificationEmail(to string, code string) error {
	m := mail.NewMessage()
	m.SetHeader("From", "noreply@yourapp.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", "Your verification code is: "+code)
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpUrl := os.Getenv("SMTP_URL")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpPortInt, _ := strconv.Atoi(smtpPort)

	d := mail.NewDialer(smtpUrl, smtpPortInt, smtpUser, smtpPassword)

	return d.DialAndSend(m)
}
