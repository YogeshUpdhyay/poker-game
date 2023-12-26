package main

import (
	"errors"
	"fmt"
	"pokergame/pkg/board"
)

func main() {
	fmt.Println("Lets play some poker")

	// get the no of players playing poker
	noOfPlayers := getNoOfPlayers()

	// // creating players data and initializing a board
	players := GetPlayersData(noOfPlayers)
	board := board.CreateBoard(players, 10)

	fmt.Printf("Starting the hand")
	board.StartHand()
	fmt.Printf("The hand winner is %s", board.LastHandWinner())
}

func GetPlayersData(noOfPlayers int) []*board.Player {
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
	return players
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
