package deck

func getDeck() deck {
	cards := deck{"my", "name", "is", "HEY"}
	cards = newDeck()
	/*
		cards in the following code goes into the d variable of printDeck in deck type definition
	*/
	cards.printDeck()
	return cards
}
