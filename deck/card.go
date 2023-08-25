package deck

type Card struct {
	suit  string
	value string
}

func (c Card) newCard(suit string, value string) Card {
	card := Card{suit: suit, value: value}
	return card
}
