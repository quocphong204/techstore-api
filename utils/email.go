package utils

import (
	"gopkg.in/gomail.v2"
	"techstore-api/config"
)

func SendOrderEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTPUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)

	return d.DialAndSend(m)
}
