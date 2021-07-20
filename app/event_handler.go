package app

import (
	"go-game-poc/pkg/event"
)

func NewHandler() event.Handler {
	return &EventHandler{
		players: make(map[string]*Player),
	}
}

type EventHandler struct {
	players map[string]*Player
}

func (eh *EventHandler) HandleEvent(evt *event.Event) {
	switch evt.Type {
	case event.EventConnected:
		p, err := NewPlayer(evt.ClientID)
		if err != nil {
			return
		}
		eh.players[evt.ClientID] = p
	case event.EventDisconnected:
		delete(eh.players, evt.ClientID)
	case event.EventData:
		p := eh.players[evt.ClientID]
		p.AddScore(1)
	}
}
