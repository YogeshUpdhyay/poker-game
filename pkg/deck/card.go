package deck

type Card struct {
	suit     string
	ordinate string
	dealt    bool
}

func (c Card) GetSuit() string {
	return c.suit
}

func (c Card) GetOrdinate() string {
	return c.ordinate
}

func (c Card) MarkDealt() {
	c.dealt = true
}
