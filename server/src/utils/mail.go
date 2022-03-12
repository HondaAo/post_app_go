package utils

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

type MailBody struct {
	From  string
	To    string
	Title string
	Text  string
}

func AuthMail(body MailBody) {
	m := gomail.NewMessage()

	m.SetHeader("From", body.From)
	m.SetHeader("To", body.To)
	m.SetHeader("Subject", body.Title)
	m.SetBody("text/plain", body.Text)

	d := gomail.NewDialer("smtp.gmail.com", 587, body.From, "<email_password>")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ChangePasswordMail(address string, body string) {

}
