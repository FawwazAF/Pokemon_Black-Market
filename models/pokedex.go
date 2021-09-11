package models

type Pokedex struct {
	Id        uint          `json:"id"`
	Name      string        `json:"name"`
	Height    int           `json:"height"`
	Weight    int           `json:"weight"`
	Abilities []interface{} `json:"abilities"`
	Forms     []interface{} `json:"forms"`
}
