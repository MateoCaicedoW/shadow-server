package websocket

import (
	"encoding/json"
	"fmt"
)

// Listener is used to listen for messages as a hub
type listener struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var listen = listener{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func Broadcast(m []byte) {
	listen.broadcast <- m
}

// Run starts the listener
func Run() {
	for {
		select {
		case client := <-listen.register:
			listen.clients[client] = true
		case client := <-listen.unregister:
			if _, ok := listen.clients[client]; ok {
				delete(listen.clients, client)
				close(client.send)
			}
		case message := <-listen.broadcast:
			for client := range listen.clients {
				//send message only to the clients in the same room
				var jsonMap map[string]interface{}
				err := json.Unmarshal(message, &jsonMap)
				if err != nil {
					fmt.Println("Error unmarshalling message", err)
				}

				if client.section != jsonMap["element_id"] {
					continue
				}

				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(listen.clients, client)
				}
			}
		}
	}
}
