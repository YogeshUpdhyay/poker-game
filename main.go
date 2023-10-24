package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating a new deck")
	fmt.Println("Lets play some poker")

	// get the no of players playing poker
	// noOfPlayers := getNoOfPlayers()

	// deal 2 cards to each player from the deck

	// to do this we first need to
	// create a new deck
	gameDeck := newDeck()
	fmt.Println(gameDeck)
	fmt.Println(len(gameDeck))

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
