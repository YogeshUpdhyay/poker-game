package game

type Player struct {
	name       string
	chipsValue int
}

func (p *Player) SetName(name string) {
	p.name = name
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetChipsValue(chipsValue int) {
	p.chipsValue = chipsValue
}

func (p *Player) GetChipsValue() int {
	return p.chipsValue
}

func (p *Player) DecrementChipsValue(chipsToTake int) int {
	p.chipsValue = p.chipsValue - chipsToTake
	return p.chipsValue
}

func (p *Player) IncrementChipsValue(chipsToAdd int) int {
	p.chipsValue = p.chipsValue + chipsToAdd
	return p.chipsValue
}
