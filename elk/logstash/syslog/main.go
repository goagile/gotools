package main

import (
	"log"
	"log/syslog"
)

func main() {
	SysLog, err := syslog.New(syslog.LOG_NOTICE, "go-test-logstash")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(SysLog)

	log.Println("Hello, SysLog")
}
