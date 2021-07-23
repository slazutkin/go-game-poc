package app

const EnemyBaseHealth = 10

type Enemy struct {
	Level     int `json:"level"`
	MaxHealth int `json:"maxHealth"`
	Health    int `json:"health"`
}
