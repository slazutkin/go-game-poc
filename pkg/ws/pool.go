package ws

import "go-game-poc/pkg/event"

type Pool struct {
	Register      chan *Client
	Unregister    chan *Client
	Clients       map[string]*Client
	Incoming      chan InboundMessage
	events        chan *event.Event
	notifications chan *event.Event
}

func NewPool(events, notifications chan *event.Event) *Pool {
	return &Pool{
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Clients:       make(map[string]*Client),
		Incoming:      make(chan InboundMessage),
		events:        events,
		notifications: notifications,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client.ID] = client
			pool.events <- &event.Event{Type: event.EventConnected, ClientID: client.ID}
			client.Conn.WriteJSON(&OutboundMessage{Type: MessageConnected, Data: ClientData{ID: client.ID}})
		case client := <-pool.Unregister:
			delete(pool.Clients, client.ID)
			pool.events <- &event.Event{Type: event.EventDisconnected, ClientID: client.ID}
			client.Conn.WriteJSON(&OutboundMessage{Type: MessageDisconnected, Data: ClientData{ID: client.ID}})
		case message := <-pool.Incoming:
			pool.events <- &event.Event{Type: event.EventData, ClientID: message.ClientID, Data: message.Data}
		case evt := <-pool.notifications:
			client, ok := pool.Clients[evt.ClientID]
			if !ok {
				return
			}
			client.Conn.WriteJSON(evt.Data)
		}
	}
}

func (pool *Pool) Size() int {
	return len(pool.Clients)
}
