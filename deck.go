package main

type Deck []Card

func newDeck() Deck {
	// generating the 52 card deck
	currDeck := Deck{}
	suits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	ordinates := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, ordinate := range ordinates {
			currDeck = append(currDeck, Card{
				suit:     suit,
				ordinate: ordinate,
				dealt:    false,
			})
		}
	}

	return currDeck
}

func (d Deck) deal() Deck {
	return Deck{}
}
