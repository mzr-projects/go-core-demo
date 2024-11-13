package deck

import "strings"

type Card struct {
	suit  string
	value string
}

// HandleCards processes cards concurrently using channels and pointers
func HandleCards(cards <-chan *Card, processedCards chan<- *Card, done chan<- bool) {
	for card := range cards {
		// Simulate some processing on the card
		card.value = "processed_" + card.value

		// Send the processed card to the output channel
		processedCards <- card
	}

	// Signal that all cards have been processed
	done <- true
}

// ProcessDeckConcurrently demonstrates concurrent card processing
func ProcessDeckConcurrently(d deck) {
	cardsChan := make(chan *Card, len(d))
	processedCardsChan := make(chan *Card, len(d))
	doneChan := make(chan bool)

	// Start the goroutine to handle cards
	go HandleCards(cardsChan, processedCardsChan, doneChan)

	// Send cards to the channel
	for i := range d {
		cardParts := strings.Split(d[i], " of ")
		card := &Card{value: cardParts[0], suit: cardParts[1]}
		cardsChan <- card
	}
	close(cardsChan)

	// Wait for processing to complete
	<-doneChan

	// Collect processed cards
	close(processedCardsChan)
	processedDeck := make(deck, 0, len(d))
	for card := range processedCardsChan {
		processedDeck = append(processedDeck, card.value+" of "+card.suit)
	}

	// Print the processed deck
	processedDeck.printDeck()
}

func (c Card) newCard(suit string, value string) Card {
	card := Card{suit: suit, value: value}
	return card
}
