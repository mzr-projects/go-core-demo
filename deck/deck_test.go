package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := getDeck()
	if len(deck) != 16 {
		t.Errorf("length of deck is not correct")
	}
}

func TestDeal(t *testing.T) {
	cards := getDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println("========= Cards in hand")
	hand.printDeck()
	fmt.Println("========= Remaining cards")
	remainingCards.printDeck()
}

func TestToString(t *testing.T) {
	cards := getDeck()
	fmt.Printf("cards.toString() : %s", cards.toString())
}

func TestSaveToFiles(t *testing.T) {
	cards := getDeck()
	err := cards.saveToFile("newDeck")
	if err != nil {
		return
	}
}

func TestReadFromFile(t *testing.T) {
	cards := newDeckFromFile("newDeck")
	cards.printDeck()
}
