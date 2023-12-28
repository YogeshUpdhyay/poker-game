package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	TableId        int      `json:"tableId"`
	Players        []Player `json:"players"`
	AvailableSeats int      `json:"availableSeats"`
	MaxNoOfSeats   int      `json:"maxNoOfSeats"`
	PotValue       int      `json:"potValue"`
}

type Player struct {
	gorm.Model
	Name        string `json:"name"`
	BuyInAmount int    `json:"buyInAmount"`
	GameId      *uint  `json:"gameId" gorm:"foreignKey:GameID"`
}
