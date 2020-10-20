package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	S = new(logstash)
)

func main() {
	flag.StringVar(&S.Host, "h", "127.0.0.1", "logstash UDP host")
	flag.IntVar(&S.Port, "p", 5000, "logstash UDP port")
	flag.Parse()

	log.Printf("Start UDP with %v\n", S)

	for {
		fmt.Print("> ")
		var m string
		_, err := fmt.Scan("%v", &m)
		if err != nil {
			log.Fatal(err)
		}
		n, err := S.Write([]byte(m))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Write %v bytes message %v\n", n, m)
	}
}

type logstash struct {
	Host string
	Port int
}

//
// Write
//
func (s *logstash) Write(b []byte) (int, error) {
	url := fmt.Sprintf("%v:%v", s.Host, s.Port)
	log.Println("logstash url:", url)

	c, err := net.Dial("udp", url)
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
