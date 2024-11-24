package mail

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
)

//go:embed mail.html
var otpEmailTemplate string

type EmailData struct {
	Name string
	OTP  string
}

var tmpl *template.Template

func loadTemplate() {
	var err error
	tmpl, err = template.New("otpEmail").Parse(otpEmailTemplate)
	if err != nil {
		log.Fatalf("Failed to parse embedded template: %v", err)
	}
}

func SendOTPMail(email string, code string) error {

	data := EmailData{
		Name: email,
		OTP:  code,
	}

	var tplBuffer bytes.Buffer
	if err := tmpl.Execute(&tplBuffer, data); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	SendMail(email, "One Time Password",
		"Your Code is "+code+"\n This code will expire in 10 minutes.\n If you did not request this code,please ignore this email.", tplBuffer.String())
	return nil
}
