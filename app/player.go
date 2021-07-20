package app

type Player struct {
	id    string
	score int
}

func NewPlayer(id string) (*Player, error) {
	return &Player{
		id:    id,
		score: 0,
	}, nil
}

func (p *Player) ID() string {
	return p.id
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) AddScore(score int) {
	p.score += score
}
