package main

type Card struct {
	suit     string
	ordinate string
	dealt    bool
}

func (c Card) getSuit() string {
	return c.suit
}

func (c Card) getOrdinate() string {
	return c.ordinate
}

func (c Card) markDealt() {
	c.dealt = true
}
