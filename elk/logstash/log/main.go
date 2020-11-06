package main

import (
	"fmt"
	"log"
	"net"
	"flag"
)

var (
	M string
)

func main() {

	flag.StringVar(&M, "m", "ping logstash", "message to send")
	flag.Parse()

	logstash := &LogStash{Host: "127.0.0.1", Port: 5000}

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("[go-test-service] ")
	log.SetOutput(logstash)

	log.Println(M)
}

type LogStash struct {
	Host string
	Port int
}

func (s *LogStash) Write(b []byte) (int, error) {
	url := fmt.Sprintf("%v:%v", s.Host, s.Port)
	c, err := net.Dial("tcp", url)
	if err != nil {
		return 0, err
	}
	defer c.Close()
	n, err := c.Write(b)
	if err != nil {
		return 0, err
	}
	return n, nil
}
