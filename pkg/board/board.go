package board

import "poker-game/pkg/deck"

type Board struct {
	players       []*Player
	smallBlind    int
	bigBlind      int
	currDealerIdx int
	potValue      int
	deck          *deck.Deck
}

func CreateBoard(players []*Player, blindValue int) Board {
	deck := deck.NewDeck()
	return Board{
		players:       players,
		smallBlind:    blindValue,
		bigBlind:      blindValue * 2,
		currDealerIdx: 0,
		potValue:      0,
		deck:          &deck,
	}
}

func (board *Board) StartHand() {}

func (board *Board) LastHandWinner() string {
	return ""
}
