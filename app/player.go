package app

import "go-game-poc/pkg/event"

type Player struct {
	id      string
	credits int
	damage  int
	enemy   *Enemy
}

func NewPlayer(id string) *Player {
	return &Player{
		id:      id,
		credits: 0,
		damage:  1,
		enemy: &Enemy{
			Level:     1,
			Health:    EnemyBaseHealth,
			MaxHealth: EnemyBaseHealth,
		},
	}
}

func (p *Player) AddCredits() {
	p.credits += p.damage
}

func (p *Player) DamageEnemy() {
	p.enemy.Health = p.enemy.Health - p.damage
	if p.enemy.Health <= 0 {
		p.enemy.Level++
		p.enemy.MaxHealth = EnemyBaseHealth * p.enemy.Level
		p.enemy.Health = p.enemy.MaxHealth
	}
}

func (p *Player) PlayerStateEvent() *event.Event {
	return &event.Event{
		Type:     "EVT_PLAYER_STATE",
		ClientID: p.id,
		Data: struct {
			Credits int    `json:"credits"`
			Damage  int    `json:"damage"`
			Enemy   *Enemy `json:"enemy"`
		}{

			Credits: p.credits,
			Damage:  p.damage,
			Enemy:   p.enemy,
		},
	}
}
