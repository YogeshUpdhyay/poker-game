package serializers

import "poker-game/pkg/game"

type JoinAGame struct {
	GameId int          `json:"gameId"`
	Player *game.Player `json:"player"`
}
