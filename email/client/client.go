package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/khardi/trygo/notify/message"
)

var (
	host = "http://127.0.0.1"
	port = 7070
)

func main() {
	mg := NewMailGateway(host, port)

	email := message.NewEmail()

	resp, err := mg.SendMail(email)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(resp))
}

// MailGateway - modeling of Mailgun Service
type MailGateway struct {
	host string
	port int
}

// NewMailGateway - returns new Mailgun Service
func NewMailGateway(host string, port int) MailGateway {
	return MailGateway{host, port}
}

func (g *MailGateway) sendURL() string {
	return fmt.Sprintf("%s:%v/sendmail/", g.host, g.port)
}

// SendMail - endpoint to send email
func (g *MailGateway) SendMail(email *message.Email) ([]byte, error) {
	emailJSON, err := email.Marshal()
	if err != nil {
		return nil, fmt.Errorf("Email marshal error: %v ", err)
	}

	appJSON := "application/json"
	resp, err := http.Post(g.sendURL(), appJSON, bytes.NewBuffer(emailJSON))
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Email send error: %v ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Server responce read error: %v ", err)
	}

	return body, nil
}
