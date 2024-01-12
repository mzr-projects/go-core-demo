package deck

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Here we defined a new type named deck of string
--type deck []string
*/
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/*
(d deck) is the receiver of the function, here we say any variable of type deck now gets access to the print method

d is (the reference) actual copy of the deck we're working with is available in the function d is essentially
reference to the cards variable we are passing to it, it's similar to the word "this" in java
*/
func (d deck) printDeck() {
	for i, deck := range d {
		fmt.Printf("index: %d, value : %s\n", i, deck)
	}
}

/*
This function returns multiple values at the same time
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	/*
		deck is actually a slice of strings, so we can do like the following
	*/
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)

	if err != nil {
		/*
			Option 1 : log the error and create a new deck
			err = fmt.Errorf("error happened: %s", err)
			fmt.Println(err.Error())
			return newDeck()
		*/
		/*
			Option 2 : log the error and entirely quit the program
		*/
		err = fmt.Errorf("error happened: %s", err)
		fmt.Println(err.Error())
		/*
			If we pass 0 it means the program successfully did the job any other value means some errors happened
		*/
		os.Exit(1)
	}

	return strings.Split(string(byteSlice), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		/*
			Intn -> will generate a random number between 0 and len(d)-1
			newPosition := rand.Intn(len(d) - 1)
		*/
		/*
			Here we changed the mechanism of random number generation to a more robust one using seed and source
			to make sure get new numbers out of the random generator after each run
		*/
		newPosition := r.Intn(len(d) - 1)
		/*
			This is the weired syntax that swaps the values in a slice
			d[newPosition] on the right hand side will replace with the d[i] on the left hand side
			d[i] on the right hand side will replace with the d[newPosition] on the left hand side
		*/
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
