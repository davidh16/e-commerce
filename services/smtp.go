package services

import (
	"bytes"
	"e-commerce/config"
	"e-commerce/templates"
	"fmt"
	"github.com/iancoleman/strcase"
	"html/template"
	"net/smtp"
	"strings"
)

func (s Service) SendVerificationEmail(emailAddress string, firstName string, verificationToken string) error {

	cfg := config.GetConfig()

	templateName := "verification_email.html"

	t1 := template.Must(template.New(templateName).Funcs(template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToLower":      strings.ToLower,
		"ToCamel":      strcase.ToCamel,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).ParseFS(templates.Files, templateName))

	type verificationEmail struct {
		Name    string
		Token   string
		BaseUrl string
	}

	var tpl bytes.Buffer
	if err := t1.Execute(&tpl, verificationEmail{
		Name:    firstName,
		Token:   verificationToken,
		BaseUrl: cfg.BaseUrl,
	}); err != nil {
		fmt.Println("Email Had Not Been Sent Successfully!")
	}

	// Receiver email address.
	to := []string{
		emailAddress,
	}

	message := []byte("Subject: " + "Verifikacija email adrese" + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		tpl.String())

	// Authentication.
	auth := smtp.PlainAuth("", cfg.SmtpFrom, cfg.GoogleAppPassword, cfg.SmtpHost)

	// Sending email.
	err := smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.SmtpFrom, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Email Sent Successfully!")
	return nil
}
