package message

import (
	"encoding/json"
	"fmt"
)

const mail = "daniefroz911@gmail.com"

// Email - is modeling email message
type Email struct {
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Recipient string `json:"recipient"`
}

// NewEmail - Create New Email Instance
func NewEmail() *Email {
	return &Email{
		Sender:    mail,
		Subject:   "Hello",
		Message:   "Hello Froz!",
		Recipient: mail,
	}
}

// Marshal Email message to JSON
func (e *Email) Marshal() ([]byte, error) {
	emailJSON, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("email marshal error: %v ", err)
	}
	return emailJSON, err
}

// Unmarshal - decode selef Email fields from JSON
func (e *Email) Unmarshal(body []byte) error {
	err := json.Unmarshal(body, &e)
	if err != nil {
		return fmt.Errorf("email unmarshal error: %v ", err)
	}
	return nil
}
