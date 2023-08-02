package deck

import "testing"

func TestDeck(t *testing.T) {
	deck := getDeck()
	if len(deck) != 3 {
		t.Errorf("length of deck is not correct")
	}
}
