package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// создание соединения
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErr(err)
	defer conn.Close()

	// создание канала
	ch, err := conn.Channel()
	failOnErr(err)
	defer ch.Close()

	// создание очереди
	q, err := ch.QueueDeclare(
		"my_queue", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // wait time for processing
		nil,        // args
	)
	failOnErr(err)

	// тело сообщения
	body := "HELLO"

	// публикация сообщений
	err = ch.Publish(
		"",     // exchange
		q.Name, // key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	fmt.Printf("Sent message: %v\n", body)
	failOnErr(err)

	// получение сообщений
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // non-local
		false,  // no-wait
		nil,    // args
	)
	failOnErr(err)

	for d := range msgs {
		fmt.Printf("Recieved: %v\n", string(d.Body))
	}
}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
