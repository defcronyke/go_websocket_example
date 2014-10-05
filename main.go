package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net/http"
)

func EchoServer(ws *websocket.Conn) { // Simplest WebSocket Server
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	fmt.Printf("Listening on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
