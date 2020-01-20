package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Printf("client: %+v\n", client)
				if err := client.Conn.WriteJSON(Message{Source: client.ID, Type: 1, Body: "New User Joined..."}); err != nil {
					fmt.Println(err)
					return
				}

			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Source: client.ID, Type: 1, Body: "User Disconnected..."}); err != nil {
					fmt.Println(err)
					return
				}
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all client in Pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
