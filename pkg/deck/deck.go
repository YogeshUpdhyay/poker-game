package deck

import (
	"fmt"
	"math/rand"
)

type Deck []Card

func NewDeck() Deck {
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

func (d Deck) Deal(noOfCards int) Deck {
	fmt.Println(noOfCards)
	deal := Deck{}
	for i := 0; i < noOfCards; i++ {
		randIndex := rand.Intn(len(d))
		d[randIndex].MarkDealt()
		cardDealt := d[randIndex]
		deal = append(deal, cardDealt)
	}
	return deal
}
