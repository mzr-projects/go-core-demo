package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	d := getDeck()
	if len(d) != 16 {
		t.Errorf("length of d is not correct")
	}

	if d[0] != "ace of spades" {
		t.Errorf("We expexcted ace of spades as the first card but we got %s", d[0])
	}

	if d[len(d)-1] != "four of clubs" {
		t.Errorf("We expexcted four of clubs as the last card but we got %s", d[len(d)-1])
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
	} else {
		t.Errorf("Couldn't save to file")
	}
}

func TestReadFromFile(t *testing.T) {
	cards := newDeckFromFile("newDeck")
	cards.printDeck()
}

func TestShuffle(t *testing.T) {
	cards := newDeck()
	cards.shuffle()
	cards.printDeck()
}
