package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/khardi/trygo/notify/message"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

var mg MailgunMessageGateway

func main() {
	log.Println("Starting the application...")

	m, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		log.Fatalf("Mailgun create error: %v", err)
	}
	mg = MailgunMessageGateway{m}

	router := httprouter.New()
	router.POST("/sendmail/", sendHandler)

	port := 7070
	addr := fmt.Sprintf(":%v", port)
	log.Printf("ListenAndServe on %v\n", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// "/sendmail/"
func sendHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("request to: /sendmail/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("request body error: %v ", err)
	}
	log.Println("request body read ... OK!")

	var email *message.Email
	err = email.Unmarshal(body)
	if err != nil {
		log.Printf("email unmarshal error: %v ", err)
	}
	log.Println("email unmarshal ... OK!")

	log.Println("sending ...")
	result, err := mg.SendEmail(email)
	if err != nil {
		log.Printf("mailgun send error: %v ", err)
		// w.WriteHeader(500)
		// w.Write([]byte(err.Error()))
	}
	fmt.Println(string(result))
	// w.WriteHeader(200)
	// w.Write([]byte(result))
}

// ------------

// MailgunMessageGateway - gateway to MailGun Service
type MailgunMessageGateway struct {
	mg mailgun.Mailgun
}

// NewMailgunMessageGateway - create new Mailgun service
// $ export MG_API_KEY=your-api-key
// $ export MG_DOMAIN=your-domain
// $ export MG_PUBLIC_API_KEY=your-public-key
// $ export MG_URL="https://api.mailgun.net/v3"
// func NewMailgunMessageGateway() (MailgunMessageGateway, error) {
// 	mg, err := mailgun.NewMailgunFromEnv()
// 	if err != nil {
// 		return nil, fmt.Errorf("Mailgun create error: %v", err)
// 	}
// 	result := MailgunMessageGateway{mg}
// 	return result, nil
// }

// SendEmail - send email message to mailgun service
func (g *MailgunMessageGateway) SendEmail(email *message.Email) (string, error) {
	log.Println("SendEmail")

	msg := mailgun.NewMessage(
		email.Sender,
		email.Subject,
		email.Message,
		email.Recipient)
	log.Println("mailgun message created ... OK!")
	log.Printf("msg = %v", msg)
	// resp, id, err := g.mg.Send(msg)
	// if err != nil {
	// 	return "", err
	// }
	// result := fmt.Sprintf("%s: %s\n", id, resp)
	// return result, nil
	return "", nil
}
