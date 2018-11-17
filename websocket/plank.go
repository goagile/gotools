package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Port int

func main() {
	// parse cmd flags
	flag.IntVar(&Port, "p", 8081, "port")
	flag.Parse()

	// handle static files
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// handle favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	// handle websocket requests
	http.HandleFunc("/ws", ws)

	// start server
	log.Printf("Server start. Port: %v", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", Port), nil)
	if err != nil {
		log.Fatalf("fatal: %v", err)
	}
}

// websocked config
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// handle websocket requests
func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ws error: %v", err)
		http.Error(w, "ws error", http.StatusBadRequest)
		return
	}
	for {
		_, bytmsg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ws read message error: %v", err)
			http.Error(w, "ws read message error", http.StatusBadRequest)
			return
		}
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(bytmsg))
		go serve(bytmsg, conn)
	}
}

func serve(request []byte, conn *websocket.Conn) {
	err := conn.WriteMessage(1, request)
	if err != nil {
		log.Printf("ws write message error: %v", err)
		return
	}
}
