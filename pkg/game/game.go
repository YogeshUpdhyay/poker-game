package game

import "poker-game/pkg/deck"

type Game struct {
	players       []*Player
	smallBlind    int
	bigBlind      int
	currDealerIdx int
	potValue      int
	deck          *deck.Deck
}

func CreateBoard(players []*Player, blindValue int) Game {
	deck := deck.NewDeck()
	return Game{
		players:       players,
		smallBlind:    blindValue,
		bigBlind:      blindValue * 2,
		currDealerIdx: 0,
		potValue:      0,
		deck:          &deck,
	}
}

func (Game *Game) StartHand() {}

func (Game *Game) LastHandWinner() string {
	return ""
}
