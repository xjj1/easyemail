// easyemail project easyemail.go
package easyemail

import (
	"crypto/tls"
	"strings"

	"gopkg.in/gomail.v2"
)

type EasyEmail struct {
	From        string
	To          string
	Server      string
	Subject     string
	BodyType    string
	Body        string
	Attachments []string
}

func (m *EasyEmail) SendMail() (err error) {

	d := &gomail.Dialer{Host: m.Server, Port: 25}
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	rcpt := strings.Split(m.To, ";")
	mail.SetHeader("To", rcpt...)
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody(m.BodyType, m.Body)
	for _, fname := range m.Attachments {
		mail.Attach(fname)
	}
	if err := d.DialAndSend(mail); err != nil {
		return (err)
	}

	return nil
}
