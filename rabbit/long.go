package main

import (
	"log"

	"github.com/streadway/amqp"
)

func xxx() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
}
