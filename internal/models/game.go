package models

import (
	"poker-game/pkg/game"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	TableId        int            `json:"tableId"`
	Players        []*game.Player `json:"players"`
	AvailableSeats int            `json:"availableSeats"`
	MaxNoOfSeats   int            `json:"maxNoOfSeats"`
	PotValue       int            `json:"potValue"`
}
