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

type PlayerState struct {
	Type string `json:"type"`
	Data struct {
		Score int `json:"score"`
	} `json:"data"`
}

func NewPlayerState(score int) PlayerState {
	return PlayerState{
		Type: "EVT_PLAYER_STATE",
		Data: struct {
			Score int `json:"score"`
		}{Score: score},
	}
}
