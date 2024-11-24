package mail

import (
	"github.com/ChaotenHG/auth-server/config"
	gomail "gopkg.in/mail.v2"
)

var mailConfig config.SmtpConfig
var dialer *gomail.Dialer

func LoadConfig(cfg *config.Config) {
	mailConfig = cfg.Mail

	dialer = gomail.NewDialer(mailConfig.Domain, mailConfig.Port, mailConfig.Username, mailConfig.Password)

	loadTemplate()
}

func SendMail(To string, Subject string, plainBody string, htmlBody string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", mailConfig.FromMail)
	message.SetHeader("To", To)
	message.SetHeader("Subject", Subject)

	message.SetBody("text/plain", plainBody)
	message.AddAlternative("text/html", htmlBody)

	return dialer.DialAndSend(message)
}
