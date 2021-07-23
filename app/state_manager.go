package app

import "go-game-poc/pkg/event"

func NewStateManager(inevt, outevt chan *event.Event) *StateManager {
	return &StateManager{
		players: make(map[string]*Player),
		inevt:   inevt,
		outevt:  outevt,
	}
}

type StateManager struct {
	players map[string]*Player
	inevt   chan *event.Event
	outevt  chan *event.Event
}

func (eh *StateManager) Start() {
	for {
		evt := <-eh.inevt
		eh.HandleEvent(evt)
	}
}

func (eh *StateManager) HandleEvent(evt *event.Event) {
	switch evt.Type {
	case event.EventConnected:
		p := NewPlayer(evt.ClientID)
		eh.players[evt.ClientID] = p
		eh.outevt <- p.PlayerStateEvent()
	case event.EventDisconnected:
		delete(eh.players, evt.ClientID)
	case event.EventData:
		p := eh.players[evt.ClientID]
		p.AddCredits()
		p.DamageEnemy()
		eh.outevt <- p.PlayerStateEvent()
	}
}
