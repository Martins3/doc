package main;

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func main() {
	// compose the message
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "13006369675@163.com"}
	m.To = []string{"hubachelar@qq.com"}

	// add attachments
	if err := m.Attach("/home/shen/.SpaceVim.d/.ycm_extra_conf.py"); err != nil {
		log.Fatal(err)
	}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "13006369675@163.com", "pwd", "smtp.zoho.com")
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		log.Fatal(err)
	}
}
