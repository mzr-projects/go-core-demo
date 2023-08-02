package deck

import "fmt"

/*
Here we defined a new type named deck of string
*/
type deck []string

/*
(d deck) is the receiver of the function, here we say any variable of type deck now gets access to the print method

	d is (the reference) actual copy of the deck we're working with is available in the function d is essentially reference
	to the cards variable we are passing to it, it's similar to the word "this" in java
*/
func (d deck) printDecks() {
	for i, deck := range d {
		fmt.Printf("index: %d, value : %s\n", i, deck)
	}
}
