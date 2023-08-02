package data_structures

import "fmt"

/*
Arrays : Fixed length list of things
Slice  : An array that can grow or shrink but every element in the slice must be of same TYPE
*/

func workWithArrays() {
	cards := []string{"man", "you"}
	for i, card := range cards {
		fmt.Printf("index : %d\t, value : %s\n", i, card)
	}
}

func workWithSlices() {
	cards := []string{"me", "you", newCard("My man")}
	/*
		The append function doesn't modify the cards slice just create a new one
	*/
	cards = append(cards, "HOLY")

	for i, card := range cards {
		fmt.Printf("index : %d\t, value : %s\n", i, card)
	}
}

func newCard(value string) string {
	return value
}
