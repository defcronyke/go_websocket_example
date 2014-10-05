package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var num_connections uint64 = 0
var connections map[uint64]*websocket.Conn

// Simplest WebSocket Server
func EchoServer(ws *websocket.Conn) {

	conn_num := num_connections // Keep track of all connections.
	connections[conn_num] = ws
	num_connections++

	log.Printf("connection id %v opened", conn_num)

	defer func() {
		err := ws.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}

		log.Printf("connection id %v closed", conn_num)
		delete(connections, conn_num)
	}()
	io.Copy(ws, ws)
}

// Send a message to all clients every time a new client connects.
func EmitServer(ws *websocket.Conn) {

	conn_num := num_connections // will be used as an identifier for the connection.
	connections[conn_num] = ws  // Save the connection object in a map.
	num_connections++

	log.Printf("connection id %v opened", conn_num)

	defer func() { // On connection closed
		err := ws.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}

		log.Printf("connection id %v closed", conn_num)
		delete(connections, conn_num) // Remove connection from map.
	}()

	for _, val := range connections { // Write to all clients concurrently.
		go val.Write([]byte("New client connected: ID " + strconv.FormatInt(int64(num_connections-1), 10)))
	}

	io.Copy(ioutil.Discard, ws) // Keep connection open.
}

func main() {
	connections = make(map[uint64]*websocket.Conn)

	http.Handle("/echo-ws", websocket.Handler(EchoServer))
	http.Handle("/emit-ws", websocket.Handler(EmitServer))

	http.Handle("/", http.FileServer(http.Dir("webroot")))

	fmt.Printf("Listening on: http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
