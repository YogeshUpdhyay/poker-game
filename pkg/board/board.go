package board

type Board struct {
	players       []Player
	smallBlind    int
	bigBlind      int
	currDealerIdx int
	potValue      int
}

func CreateBoard(players []Player, blindValue int) Board {
	return Board{
		players:       players,
		smallBlind:    blindValue,
		bigBlind:      blindValue * 2,
		currDealerIdx: 0,
		potValue:      0,
	}
}
