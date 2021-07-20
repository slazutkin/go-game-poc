package ws

import (
	"go-game-poc/pkg/event"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Incoming   chan InboundMessage
	Handler    event.Handler
}

func NewPool(handler event.Handler) *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Incoming:   make(chan InboundMessage),
		Handler:    handler,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			evt := &event.Event{Type: event.EventConnected, ClientID: client.ID}
			pool.Handler.HandleEvent(evt)
			client.Conn.WriteJSON(&OutboundMessage{Type: MessageConnected, Data: ClientData{ID: client.ID}})
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			evt := &event.Event{Type: event.EventConnected, ClientID: client.ID}
			pool.Handler.HandleEvent(evt)
			client.Conn.WriteJSON(&OutboundMessage{Type: MessageDisconnected, Data: ClientData{ID: client.ID}})
		case message := <-pool.Incoming:
			evt := &event.Event{Type: event.EventData, ClientID: message.ClientID, Data: message.Data}
			pool.Handler.HandleEvent(evt)
		}
	}
}

func (pool *Pool) Size() int {
	return len(pool.Clients)
}
