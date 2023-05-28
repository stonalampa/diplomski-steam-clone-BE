package helpers

import (
	"fmt"
	"log"
	"main/utils"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

func sendEmail(toEmail, subject string, body string) bool {
	from := "steam-clone@diplomski.com"
	user := "5ab7a2880ac4ef"
	password := "4089efa556ef43"
	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	to := []string{toEmail}
	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html\r\n\r\n"+
		"%s\r\n", from, toEmail, subject, body))

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Printf("Email sent successfully to %s\n", toEmail)
	return true
}

func GenerateNewPasswordEmail(toEmail, newPassword string, htmlFilePath string) bool {
	fileContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("emailTemplate").Parse(string(fileContent))
	if err != nil {
		log.Fatal(err)
		return false
	}

	var bodyBuilder strings.Builder
	err = tmpl.Execute(&bodyBuilder, struct{ Password string }{Password: newPassword})
	if err != nil {
		log.Fatal(err)
		return false
	}

	body := bodyBuilder.String()
	return sendEmail(toEmail, "Steam Clone - New Password", body)
}

func ConfirmationEmail(toEmail, htmlFilePath string) bool {
	fileContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("emailTemplate").Parse(string(fileContent))
	if err != nil {
		log.Fatal(err)
		return false
	}

	token, err := utils.GenerateConfirmationToken(toEmail)
	if err != nil {
		log.Fatal(err)
		return false
	}

	var bodyBuilder strings.Builder
	err = tmpl.Execute(&bodyBuilder, struct{ ConfirmationLink string }{ConfirmationLink: "http://localhost:3030/api/users/confirm?token=" + token})
	if err != nil {
		log.Fatal(err)
		return false
	}

	body := bodyBuilder.String()
	return sendEmail(toEmail, "Steam Clone - Account confirmation email", body)
}

func GenerateSuccessfulPurchaseEmail(toEmail, title, price, htmlFilePath string) bool {
	fileContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("emailTemplate").Parse(string(fileContent))
	if err != nil {
		log.Fatal(err)
		return false
	}

	var bodyBuilder strings.Builder
	err = tmpl.Execute(&bodyBuilder, struct{ Title, Price string }{
		Title: title,
		Price: price,
	})
	if err != nil {
		log.Fatal(err)
		return false
	}

	body := bodyBuilder.String()
	return sendEmail(toEmail, "Steam Clone - Congratulations on Your Purchase", body)
}
