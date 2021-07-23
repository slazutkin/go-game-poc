package app

type DamageUpgrade struct {
	Level       int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int
}
