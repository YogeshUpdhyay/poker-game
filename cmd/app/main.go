package main

import (
	"errors"
	"fmt"

	// "pokergame/pkg/deck"
	"pokergame/pkg/board"
)

func main() {
	fmt.Println("Lets play some poker")

	// get the no of players playing poker
	noOfPlayers := getNoOfPlayers()

	// creating players data
	players := make([]*board.Player, noOfPlayers)
	for currPlayerIdx := 0; currPlayerIdx < noOfPlayers; currPlayerIdx++ {
		players[currPlayerIdx] = &board.Player{}

		fmt.Println("Enter the player name and the buy in amount")
		var currPlayerName string
		var currPlayerBuyIn int
		_, err := fmt.Scanf("%s %d", &currPlayerName, &currPlayerBuyIn)
		if err != nil {
			panic(errors.New("error loading the player info"))
		}
		players[currPlayerIdx].SetName(currPlayerName)
		players[currPlayerIdx].SetChipsValue(currPlayerBuyIn)
	}

	fmt.Println(players)

	// deal 2 cards to each player from the deck

	// to do this we first need to
	// create a new deck
	// gameDeck := deck.NewDeck()

	// deal a hand of poker
	// hand := gameDeck.Deal(2)
	// fmt.Println((hand))
	// fmt.Println((gameDeck))
}

func getNoOfPlayers() int {
	var noOfPlayers int
	fmt.Println("Enter the no of players")
	_, err := fmt.Scanf("%d", &noOfPlayers)

	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return noOfPlayers
}
