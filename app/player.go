package app

import "go-game-poc/pkg/event"

type Player struct {
	id      string
	credits int
	damage  int
}

func NewPlayer(id string) (*Player, error) {
	return &Player{
		id:      id,
		credits: 0,
		damage:  1,
	}, nil
}

func (p *Player) AddCredits() {
	p.credits += p.damage
}

func (p *Player) PlayerStateEvent() *event.Event {
	return &event.Event{
		Type:     "EVT_PLAYER_STATE",
		ClientID: p.id,
		Data: struct {
			Credits int `json:"credits"`
			Damage  int `json:"damage"`
		}{

			Credits: p.credits,
			Damage:  p.damage},
	}
}
