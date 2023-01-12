package helpers

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail() {
	from := "steam-clone@diplomski.com"

	user := "5ab7a2880ac4ef"
	password := "4089efa556ef43"

	to := []string{
		"radanovic.m@hotmail.com",
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	msg := []byte("From: john.doe@example.com\r\n" +
		"To: radanovic.m@hotmail.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Email sent successfully to %s\n", addr)

}
