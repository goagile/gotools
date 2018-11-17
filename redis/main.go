package main

import (
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	defer conn.Close()

	resp := conn.Cmd("SET", "a", 10)
	if resp.Err != nil {
		log.Fatal(resp.Err)
	}

	log.Println(resp)
}
