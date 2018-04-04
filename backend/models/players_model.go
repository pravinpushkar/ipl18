package models

type PlayersModel struct {
	Players []Player `json:"players"`
}

type Player struct {
	PlayerId string `json:"id"`
	Name     string `json:"name"`
	TeamId   string `json:"teamId"`
	Role     string `json:"role"`
}