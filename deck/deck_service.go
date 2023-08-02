package deck

func getDeck() deck {
	cards := deck{"my", "name", "is", "HEY"}
	/*
		cards in the following code goes into the d variable of printDecks in deck type definition
	*/
	cards.printDecks()
	return cards
}
