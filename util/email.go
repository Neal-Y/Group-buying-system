package util

import (
	"fmt"
	"net/smtp"
	"shopping-cart/config"
)

type EmailConfig struct {
	SMTPHost       string
	SMTPPort       string
	SenderEmail    string
	SenderPassword string
}

func NewEmailConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost:       "smtp.gmail.com",
		SMTPPort:       "587",
		SenderEmail:    config.AppConfig.Gmail,
		SenderPassword: config.AppConfig.GmailSecret,
	}
}

func (ec *EmailConfig) SendEmail(to, subject, body string) error {
	from := ec.SenderEmail
	pass := ec.SenderPassword
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(ec.SMTPHost+":"+ec.SMTPPort,
		smtp.PlainAuth("", from, pass, ec.SMTPHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

func SendResetCodeEmail(to, code string) error {
	ec := NewEmailConfig()
	subject := "Password Reset Request"
	body := fmt.Sprintf("Your verification code is: %s", code)
	return ec.SendEmail(to, subject, body)
}
